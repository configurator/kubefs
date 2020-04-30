package kube

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

var _ f.Dir = (*Resource)(nil)

func (r *Resource) List() ([]string, error) {
	kubectl := r.Context.kubectl

	// Note we use Namespace here; for global resources namespace "" is expected and does nothing
	list, err := kubectl.Resource(r.GVR).Namespace(r.Namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := []string{}
	for _, i := range list.Items {
		if r.Context.Settings.ShowYamlFiles {
			result = append(result, i.GetName()+".yaml")
		}
		if r.Context.Settings.ShowJsonFiles {
			result = append(result, i.GetName()+".json")
		}
	}
	return result, nil
}

func (r *Resource) Get(name string) (f.Node, error) {
	name, extension := SplitFileExtension(name)

	return &Item{
		Resource:  r,
		Name:      name,
		Extension: extension,
	}, nil
}
