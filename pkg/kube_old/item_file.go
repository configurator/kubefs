package kube

import (
	"encoding/json"
	"fmt"

	"github.com/configurator/kubefs/pkg/kfuse"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/yaml"
)

func (i *Item) ToFile(kfs *kfuse.KubeFS) *kfuse.File {

	return &kfuse.File{
		KubeFS:       kfs,
		ReadContents: i.ReadContents,
	}
}

func (i *Item) ReadContents() ([]byte, error) {
	d, err := dynamic.NewForConfig(i.restConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	gvr := schema.GroupVersionResource{
		Group:    i.Resource.Group,
		Version:  i.Resource.Version,
		Resource: i.Resource.Name,
	}

	resourceApi := d.Resource(gvr)

	var resourceInterface dynamic.ResourceInterface
	if i.Namespace != "" {
		resourceInterface = resourceApi.Namespace(i.Namespace)
	} else {
		resourceInterface = resourceApi
	}
	object, err := resourceInterface.Get(i.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var data []byte
	switch i.Extension {
	case "json":
		data, err = json.Marshal(object)
	default:
		data, err = yaml.Marshal(object)
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
