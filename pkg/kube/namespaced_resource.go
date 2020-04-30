package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

type NamespacedResource struct {
	f.BaseDir
	Context      *Context
	ResourceType metav1.APIResource
	GVR          schema.GroupVersionResource
	GVK          schema.GroupVersionKind
}
