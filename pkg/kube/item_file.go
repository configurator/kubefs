package kube

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

var _ f.File = (*Item)(nil)

func (i *Item) ReadEntireContents() ([]byte, error) {
	r := i.Resource
	kubectl := r.Context.kubectl

	object, err := kubectl.Resource(r.GVR).Namespace(r.Namespace).Get(i.Name, metav1.GetOptions{})
	if err != nil {
		if status, ok := err.(*errors.StatusError); ok {
			if status.ErrStatus.Code == 404 {
				return nil, &f.ErrorNotFound{}
			}
		}
		return nil, err
	}

	var data []byte
	switch i.Extension {
	case "json":
		if r.Context.PrettyJson {
			data, err = json.MarshalIndent(object, "", "    ")
		} else {
			data, err = json.Marshal(object)
		}
	default:
		data, err = yaml.Marshal(object)
	}
	if err != nil {
		return nil, err
	}

	if len(data) > 0 && data[len(data)-1] != '\n' {
		// Add a trailing newline if it's missing
		data = append(data, '\n')
	}

	return data, nil
}
