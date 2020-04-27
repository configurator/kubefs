module github.com/configurator/kubefs

go 1.13

require (
	bazil.org/fuse v0.0.0-20180421153158-65cc252bf669
	github.com/spf13/pflag v1.0.5
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
)

replace bazil.org/fuse => bazil.org/fuse v0.0.0-20180421153158-65cc252bf669 // pin to latest version that supports macOS. see https://github.com/bazil/fuse/issues/224
