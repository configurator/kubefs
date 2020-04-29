package kube

import (
	"fmt"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

var _ f.Dir = (*Context)(nil)

func (c *Context) loadResources() error {
	if c.resourceTypes != nil {
		// Already loaded
		return nil
	}
	resources, err := discovery.ServerPreferredResources(c.discovery)
	if err != nil {
		fmt.Println(err)
		return err
	}

	result := make(map[string]f.Node)
	for _, list := range resources {
		group, version := SplitGroupVersion(list.GroupVersion)
		for _, r := range list.APIResources {
			// Add missing metadata
			if r.Group == "" {
				r.Group = group
			}
			if r.Version == "" {
				r.Version = version
			}

			gvr := schema.GroupVersionResource{
				Group:    group,
				Version:  version,
				Resource: r.Name,
			}

			if r.Namespaced {
				result[r.Name] = &NamespacedResource{
					Context:      c,
					ResourceType: r,
					GVR:          gvr,
				}
			} else {
				result[r.Name] = &Resource{
					Context:      c,
					ResourceType: r,
					GVR:          gvr,
				}
			}
		}
	}

	c.resourceTypes = result
	return nil
}

func (c *Context) List() ([]string, error) {
	err := c.loadResources()
	if err != nil {
		return nil, err
	}

	result := []string{}
	for name := range c.resourceTypes {
		result = append(result, name)
	}
	return result, nil
}
func (c *Context) Get(name string) (f.Node, error) {
	err := c.loadResources()
	if err != nil {
		return nil, err
	}

	r, _ := c.resourceTypes[name]
	// r will be nil if not found, which is what we want
	return r, nil
}
