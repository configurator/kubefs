package kube

import (
	"strings"
)

func SplitGroupVersion(groupVersion string) (string, string) {
	index := strings.LastIndex(groupVersion, "/")
	if index == -1 {
		return "", groupVersion
	} else {
		return groupVersion[0:index], groupVersion[index+1:]
	}
}

func SplitFileExtension(name string) (string, string) {
	if dot := strings.LastIndex(name, "."); dot != -1 {
		return name[0:dot], name[dot+1:]
	}
	return name, ""
}
