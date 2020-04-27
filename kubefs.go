package main

import (
	"fmt"
	"os"
	"path"

	_ "bazil.org/fuse"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s <mountpoint>\n", path.Base(os.Args[0]))
		pflag.PrintDefaults()
	}

	pflag.Parse()
	args := pflag.Args()
	if len(args) != 1 {
		pflag.Usage()
		os.Exit(1)
	}

	mountpoint := args[0]

	mount(mountpoint)
}

func mount(mountpoint string) {}
