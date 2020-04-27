package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/pflag"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"github.com/configurator/kubefs/pkg/kubefs"
)

func main() {
	pflag.Usage = func() {
		fmt.Printf("Usage:\n\t%s <mountpoint>\n", path.Base(os.Args[0]))
		pflag.PrintDefaults()
	}

	unmount := pflag.BoolP("unmount", "u", false, "Unmount")

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
		err := mount(mountpoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Mounted")
	}
}

func mount(mountpoint string) error {
	c, err := fuse.Mount(mountpoint)
	if err != nil {
		return err
	}
	defer c.Close()

	kfs := &kubefs.KubeFS{}

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
