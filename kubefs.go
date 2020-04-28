package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/pflag"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"github.com/configurator/kubefs/pkg/kfuse"
	"github.com/configurator/kubefs/pkg/kube"
)

func main() {
	pflag.Usage = func() {
		fmt.Printf("Usage:\n\t%s <mountpoint>\n", path.Base(os.Args[0]))
		pflag.PrintDefaults()
	}

	unmount := pflag.BoolP("unmount", "u", false, "Unmount")
	kubeconfig := pflag.StringP("kubeconfig", "c", "", "absolute path to the kubeconfig file")

	pflag.Parse()
	args := pflag.Args()
	if len(args) != 1 {
		pflag.Usage()
		os.Exit(1)
	}

	mountpoint := args[0]

	if *unmount {
		err := fuse.Unmount(mountpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		k := &kube.Kubernetes{}
		k.LoadConfig(*kubeconfig)

		err := mount(mountpoint, k)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func cleanPathAndValidateEmptyDir(mountpoint string) (string, error) {
	mountpoint, err := filepath.Abs(mountpoint)
	if err != nil {
		return "", err
	}

	stat, err := os.Stat(mountpoint)
	if err != nil {
		// Includes both "path does not exist" and other errors
		return "", err
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("Mount point %s is not a directory", mountpoint)
	}

	file, err := os.OpenFile(mountpoint, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Reading one file is enough to verify if directory isn't empty
	// We expect EOF!
	_, err = file.Readdirnames(1)
	if err != io.EOF {
		if err != nil {
			return "", err
		} else {
			return "", fmt.Errorf("Mount point %s is not an empty directory", mountpoint)
		}
	}

	return mountpoint, nil
}

func mount(mountpoint string, k *kube.Kubernetes) error {
	mountpoint, err := cleanPathAndValidateEmptyDir(mountpoint)
	if err != nil {
		return err
	}

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		return err
	}
	defer c.Close()

	kfs := &kfuse.KubeFS{}
	err = kfs.ReadCurrentUser()
	if err != nil {
		fmt.Println("Could not read current uid and gid; defaulting to root")
	}
	kfs.RootDir = k.ToDir(kfs)

	fmt.Println("Mounting kubefs on " + mountpoint)
	err = fs.Serve(c, kfs)
	if err != nil {
		return err
	}

	<-c.Ready
	if c.MountError != nil {
		return c.MountError
	}

	return nil
}
