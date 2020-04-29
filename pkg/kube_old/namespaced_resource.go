package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type NamespacedResource struct {
	config     *clientcmdapi.Config
	restConfig *rest.Config
	context    *clientcmdapi.Context
	Resource   metav1.APIResource
}

type NamespacedResourceWithNamespace struct {
	*NamespacedResource
	Namespace string
}
