/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/glog"
	"k8s.io/go/kubernetes/client"
	restclient "k8s.io/go/rest"
	"k8s.io/go/tools/clientcmd"
)

// BuildConfigFromRESTConfig converts a rest config into generated client config
// TODO(roycaihw): move it to the right place
func BuildConfigFromRESTConfig(config *restclient.Config) (*client.Configuration, error) {
	transport, err := restclient.TransportFor(config)
	if err != nil {
		return nil, err
	}
	return &client.Configuration{
		BasePath: config.Host + config.APIPath,
		Host:     config.Host,
		// TODO(roycaihw): parse scheme
		Scheme:        "https",
		DefaultHeader: map[string]string{},
		UserAgent:     config.UserAgent,
		HTTPClient: &http.Client{
			Transport: transport,
		},
	}, nil
}

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	cfg, err := BuildConfigFromRESTConfig(config)
	if err != nil {
		panic(err.Error())
	}
	glog.Errorf("converted client config: %v", cfg)

	// creates the clientset
	c := client.NewAPIClient(cfg)
	for {
		pods, _, err := c.CoreV1Api.ListPodForAllNamespaces(context.Background(), nil)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		time.Sleep(10 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
