/*
Copyright 2018 The Kubernetes Authors.

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
	"fmt"
	"time"

	"k8s.io/client/kubernetes/client"
	"k8s.io/client/kubernetes/config"
)

func main() {
	// creates the in-cluster config
	config, err := config.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	c := client.NewAPIClient(config)
	for {
		pods, _, err := c.CoreV1Api.ListPodForAllNamespaces(context.Background(), nil)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		time.Sleep(10 * time.Second)
	}
}
