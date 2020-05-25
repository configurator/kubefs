package kube

import (
	"log"

	"github.com/configurator/kubefs/pkg/cgofusewrapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Item) Delete() (err error) {
	if i.Resource.Context.Readonly {
		return &cgofusewrapper.ErrorNotImplemented{}
	}

	defer LogDeleting("file %s", i.Name)(err)

	r := i.Resource
	kubectl := r.Context.kubectl

	err = kubectl.Resource(r.GVR).Namespace(r.Namespace).Delete(i.Name, &metav1.DeleteOptions{})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
