package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/spf13/pflag"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
	kube "github.com/configurator/kubefs/pkg/kube"
)

func main() {
	pflag.Usage = func() {
		fmt.Printf("Usage:\n\t%s <mountpoint>\n", path.Base(os.Args[0]))
		pflag.PrintDefaults()
	}

	readonly := pflag.Bool("readonly", false, "readonly mode - never allow any write or change to the cluster")

	showJsonFiles := pflag.Bool("show-json-files", false, "show .json files in file listings")
	showYamlFiles := pflag.Bool("show-yaml-files", true, "show .yaml files in file listings (defaults to true, use =false to change)")
	prettyJson := pflag.Bool("pretty-json", false, "pretty-print json files")
	kubeconfig := pflag.StringP("kubeconfig", "c", "", "absolute path to the kubeconfig file")

	pflag.Parse()
	args := pflag.Args()
	if len(args) != 1 {
		pflag.Usage()
		os.Exit(1)
	}

	mountpoint := args[0]

	settings := &kube.Settings{
		ShowJsonFiles: *showJsonFiles,
		ShowYamlFiles: *showYamlFiles,
		PrettyJson:    *prettyJson,
		Readonly:      *readonly,
	}

	k := &kube.Kubernetes{
		Settings: settings,
	}
	k.LoadConfig(*kubeconfig)

	fs := &f.FS{
		Root:     k,
		Readonly: *readonly,
	}

	h := fuse.NewFileSystemHost(fs)
	h.Mount(mountpoint, nil)
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
