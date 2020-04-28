package kube

import (
	"fmt"
	"strings"

	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

var namespaceGVR = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "namespaces",
}

func (n *NamespacedResource) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return n.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return n.lookupNode(kfs, name) },
	}
}

func (n *NamespacedResource) getNamespaces() ([]string, error) {
	d, err := dynamic.NewForConfig(n.restConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list, err := d.Resource(namespaceGVR).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := []string{}
	for _, i := range list.Items {
		result = append(result, i.GetName())
	}
	return result, nil
}

func (n *NamespacedResource) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	return n.getNamespaces()
}

func (n *NamespacedResource) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	return (&NamespacedResourceWithNamespace{
		NamespacedResource: n,
		Namespace:          name,
	}).ToDir(kfs), nil
}

func (n *NamespacedResourceWithNamespace) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return n.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return n.lookupNode(kfs, name) },
	}
}

func (n *NamespacedResourceWithNamespace) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	d, err := dynamic.NewForConfig(n.restConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%#v\n", n.Resource)
	gvr := schema.GroupVersionResource{
		Group:    n.Resource.Group,
		Version:  n.Resource.Version,
		Resource: n.Resource.Name,
	}

	fmt.Println(gvr)

	list, err := d.Resource(gvr).Namespace(n.Namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := []string{}
	for _, i := range list.Items {
		result = append(result, i.GetName()+".yaml")
	}
	return result, nil
}

func (n *NamespacedResourceWithNamespace) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	extension := "yaml"
	if dot := strings.LastIndex(name, "."); dot != -1 {
		extension = name[dot+1:]
		name = name[0:dot]
	}

	return (&Item{
		config:     n.config,
		restConfig: n.restConfig,
		context:    n.context,
		Resource:   n.Resource,
		Namespace:  n.Namespace,
		Name:       name,
		Extension:  extension,
	}).ToFile(kfs), nil
}
