package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type GlobalResource struct {
	config     *clientcmdapi.Config
	restConfig *rest.Config
	context    *clientcmdapi.Context
	Resource   metav1.APIResource
}
