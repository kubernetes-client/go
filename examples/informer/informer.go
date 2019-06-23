package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kubernetes-client/go/kubernetes/client"
	"github.com/kubernetes-client/go/kubernetes/config"
)

type handler struct{}

func (h handler) OnAdd(obj interface{}) {
	ns := obj.(*client.V1Namespace)
	fmt.Printf("Added %s\n", ns.Metadata.Name)
}

func (h handler) OnUpdate(oldObj, newObj interface{}) {
	ns := newObj.(*client.V1Namespace)
	fmt.Printf("Updated %s\n", ns.Metadata.Name)
}

func (h handler) OnDelete(obj interface{}) {
	ns := obj.(*client.V1Namespace)
	fmt.Printf("Deleted %s\n", ns.Metadata.Name)
}

func main() {
	c, err := config.LoadKubeConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset := client.NewAPIClient(c)

	lister := func() ([]interface{}, string, error) {
		namespaces, _, err := clientset.CoreV1Api.ListNamespace(context.Background(), nil)
		if err != nil {
			return nil, "", err
		}
		result := make([]interface{}, len(namespaces.Items))
		for ix := range namespaces.Items {
			result[ix] = &namespaces.Items[ix]
		}
		return result, namespaces.Metadata.ResourceVersion, nil
	}

	watcher := func(resourceVersion string) (<-chan *client.Result, <-chan error) {
		watch := client.WatchClient{
			Cfg:    c,
			Client: clientset,
			Path:   "/api/v1/namespaces",
			MakerFn: func() interface{} {
				return &client.V1Namespace{}
			},
		}
		results, errors, err := watch.Connect(context.Background(), resourceVersion)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
		}
		return results, errors
	}

	extractor := func(obj interface{}) *client.V1ObjectMeta {
		return obj.(*client.V1Namespace).Metadata
	}

	cache := client.Cache{
		Extractor: extractor,
		Lister:    lister,
		Watcher:   watcher,
	}
	cache.AddEventHandler(handler{})
	go cache.Run(make(chan bool))

	for {
		fmt.Printf("----------\n")
		list := cache.List()
		for ix := range list {
			ns := list[ix].(*client.V1Namespace)
			fmt.Printf("%s %#v\n", ns.Metadata.Name, ns.Metadata.Labels)
		}
		fmt.Printf("----------\n")
		time.Sleep(5 * time.Second)
	}
}
