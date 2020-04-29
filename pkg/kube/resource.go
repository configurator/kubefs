package kube

import (
	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Resource struct {
	f.BaseDir
	Context      *Context
	ResourceType metav1.APIResource
	GVR          schema.GroupVersionResource
	Namespace    string
}
