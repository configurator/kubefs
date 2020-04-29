package kube

import (
	"bazil.org/fuse/fs"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Context struct {
	config     *clientcmdapi.Config
	restConfig *rest.Config
	context    *clientcmdapi.Context
	Name       string
	resources  map[string]fs.Node
}
