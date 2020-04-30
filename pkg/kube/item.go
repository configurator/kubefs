package kube

import (
	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

type Item struct {
	f.BaseFile
	Resource         *Resource
	Name             string
	Extension        string
	OriginalContents []byte
}

var _ f.File = (*Item)(nil)
