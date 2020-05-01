package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

func (i *Item) Write(data []byte) error {
	r := i.Resource
	kubectl := r.Context.kubectl

	// We need to clean the json a bit:
	// 1. remove some fields which are not allowed to be set in a server-side apply, and cause conflicts
	// 2. set the namespace, name, apiVersion, and kind, overriding what's in the file
	//    this allows copying one object into another, and also creating objects by touching files
	object := &unstructured.Unstructured{}
	err := yaml.Unmarshal(data, object)
	if err != nil {
		return err
	}

	object.SetManagedFields(nil)
	object.SetGeneration(0)
	object.SetResourceVersion("")
	object.SetSelfLink("")
	object.SetUID("")

	object.SetGroupVersionKind(i.Resource.GVK)
	object.SetNamespace(i.Resource.Namespace)
	object.SetName(i.Name)

	newYaml, err := yaml.Marshal(object)
	if err != nil {
		return err
	}

	force := true
	_, err = kubectl.Resource(r.GVR).Namespace(r.Namespace).Patch(
		i.Name,
		types.ApplyPatchType,
		newYaml,
		metav1.PatchOptions{
			FieldManager: "kubefs",
			Force:        &force,
		})
	return err
}
