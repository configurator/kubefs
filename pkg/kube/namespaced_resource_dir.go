package kube

import (
	"fmt"
	"strings"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ f.Dir = (*NamespacedResource)(nil)

var namespaceGVR = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "namespaces",
}

func (n *NamespacedResource) List() ([]string, error) {
	kubectl := n.Context.kubectl

	list, err := kubectl.Resource(namespaceGVR).List(metav1.ListOptions{})
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

func (n *NamespacedResource) Get(name string) (f.Node, error) {
	if strings.Contains(name, ".") {
		// Many programs check for special files and directories, and these
		// are either dotfiles (which contain a dot), or files with an extension
		// - neither of which is a valid kubernetes namespace
		// This check prevents those programs from going haywire when cding into
		// a resource directory.
		return nil, &f.ErrorNotFound{}
	}

	return &Resource{
		Context:      n.Context,
		ResourceType: n.ResourceType,
		GVR:          n.GVR,
		Namespace:    name,
	}, nil
}
