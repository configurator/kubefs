package kube

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getDefaultKubeConfig() string {
	envValue := os.Getenv("KUBECONFIG")
	if envValue != "" {
		return envValue
	}

	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home != "" {
		return filepath.Join(home, ".kube", "config")
	}

	// Cannot find default file path
	return ""
}

type Kubernetes struct {
	config *clientcmdapi.Config

	Contexts map[string]Context
}

func (k *Kubernetes) LoadConfig(kubeconfig string) error {
	configLoader := clientcmd.NewDefaultClientConfigLoadingRules()
	configLoader.ExplicitPath = kubeconfig

	config, err := configLoader.Load()
	if err != nil {
		return err
	}

	k.config = config
	k.createContextsMap()

	return nil
}

func (k *Kubernetes) createContextsMap() {
	k.Contexts = map[string]Context{}
	if k.config == nil {
		return
	}

	for name, context := range k.config.Contexts {
		k.Contexts[name] = Context{
			k.config,
			context,
			name,
		}
	}
}
