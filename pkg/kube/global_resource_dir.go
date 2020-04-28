package kube

import (
	"fmt"

	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func (g *GlobalResource) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return g.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return g.lookupNode(kfs, name) },
	}
}

func (g *GlobalResource) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	d, err := dynamic.NewForConfig(g.restConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%#v\n", g.Resource)
	gvr := schema.GroupVersionResource{
		Group:    g.Resource.Group,
		Version:  g.Resource.Version,
		Resource: g.Resource.Name,
	}

	fmt.Println(gvr)

	list, err := d.Resource(gvr).List(metav1.ListOptions{})
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

func (g *GlobalResource) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	return nil, nil
}
