package kube

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Item) Delete() error {
	fmt.Printf("Deleting file %s\n", i.Name)

	r := i.Resource
	kubectl := r.Context.kubectl

	err := kubectl.Resource(r.GVR).Namespace(r.Namespace).Delete(i.Name, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
