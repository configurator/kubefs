package kube

import (
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Context struct {
	config  *clientcmdapi.Config
	context *clientcmdapi.Context
	Name    string
}
