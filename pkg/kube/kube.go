package kube

import (
	"log"
	"os"
	"path/filepath"

	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
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
	*Settings
	f.BaseDir

	config *clientcmdapi.Config

	Contexts map[string]*Context
}

func (k *Kubernetes) LoadConfig(kubeconfig string) error {
	configLoader := clientcmd.NewDefaultClientConfigLoadingRules()
	configLoader.ExplicitPath = kubeconfig

	config, err := configLoader.Load()
	if err != nil {
		return err
	}

	k.config = config
	err = k.createContextsMap()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (k *Kubernetes) createContextsMap() error {
	config := k.config

	k.Contexts = map[string]*Context{}
	if config == nil {
		log.Println("Error in createContextsMap(): config == nil")
		return nil
	}

	for name, context := range k.config.Contexts {
		clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{
			Context: *context,
		})

		restConfig, err := clientConfig.ClientConfig()
		if err != nil {
			log.Println(err)
			return err
		}

		kubectl, err := dynamic.NewForConfig(restConfig)
		if err != nil {
			log.Println(err)
			return err
		}

		dc, err := discovery.NewDiscoveryClientForConfig(restConfig)
		if err != nil {
			log.Println(err)
			return err
		}

		k.Contexts[name] = &Context{
			ContextName: name,
			Settings:    k.Settings,
			config:      k.config,
			kubectl:     kubectl,
			discovery:   dc,
		}
	}

	return nil
}
