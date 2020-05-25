package kube

import (
	"github.com/configurator/kubefs/pkg/cgofusewrapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

func (i *Item) Write(data []byte) error {
	if i.Resource.Context.Readonly {
		return &cgofusewrapper.ErrorNotImplemented{}
	}

	r := i.Resource
	kubectl := r.Context.kubectl

	// We need to clean the json a bit:
	// 1. remove some fields which are not allowed to be set in a server-side apply, and cause conflicts
	// 2. set the namespace, name, apiVersion, and kind, overriding what's in the file
	//    this allows copying one object into another, and also creating objects by touching files
	object := map[string]interface{}{}
	err := yaml.Unmarshal(data, object)
	if err != nil {
		return err
	}

	// Get or create the metadata subkey
	// If it's not a json object, overrides it with a json object
	metadataUntyped, _ := object["metadata"]
	metadata, _ := metadataUntyped.(map[string]interface{})
	if metadata == nil {
		metadata = make(map[string]interface{})
		object["metadata"] = metadata
	}

	// Clear conflict-inducing fields
	delete(metadata, "managedFields")
	delete(metadata, "generation")
	delete(metadata, "resourceVersion")
	delete(metadata, "selfLink")
	delete(metadata, "uid")

	// Set fields known by file path
	object["apiVersion"] = r.GVK.GroupVersion().String()
	object["kind"] = r.GVK.Kind
	metadata["namespace"] = r.Namespace
	metadata["name"] = i.Name

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
