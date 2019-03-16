/*
Copyright 2019 The Kubernetes Authors.

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

	"github.com/kubernetes-client/go/kubernetes/client"
	"github.com/kubernetes-client/go/kubernetes/config"
)

func main() {
	c, err := config.LoadKubeConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset := client.NewAPIClient(c)
	pod := client.V1Pod{
		Metadata: &client.V1ObjectMeta{
			Name:      "my-pod",
			Namespace: "default",
		},
		Spec: &client.V1PodSpec{
			Containers: []client.V1Container{
				client.V1Container{
					Name:  "www",
					Image: "nginx",
				},
			},
		},
	}

	_, _, err = clientset.CoreV1Api.CreateNamespacedPod(context.Background(), "default", pod, nil)
	if err != nil {
		panic(err)
	}
}
