package kube

import (
	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

var _ f.Dir = (*Kubernetes)(nil)

func (k *Kubernetes) List() ([]string, error) {
	result := []string{}
	for name := range k.Contexts {
		result = append(result, name)
	}

	return result, nil
}

func (k *Kubernetes) Get(name string) (f.Node, error) {
	context, ok := k.Contexts[name]
	if !ok {
		return nil, &f.ErrorNotFound{}
	}
	return context, nil
}
