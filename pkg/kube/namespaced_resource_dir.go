package kube

import (
	"fmt"

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
	return &Resource{
		Context:      n.Context,
		ResourceType: n.ResourceType,
		GVR:          n.GVR,
		Namespace:    name,
	}, nil
}
