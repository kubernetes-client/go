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
	"fmt"

	"github.com/kubernetes-client/go/kubernetes/client"
	"github.com/kubernetes-client/go/kubernetes/config"
)

func main() {
	c, err := config.LoadKubeConfig()
	if err != nil {
		panic(err.Error())
	}
	cl := client.NewAPIClient(c)

	watch := client.WatchClient{
		Cfg:     c,
		Client:  cl,
		Path:    "/api/v1/namespaces",
		MakerFn: func() interface{} { return &client.V1Namespace{} },
	}
	if resultChan, errChan, err := watch.Connect(context.Background(), ""); err != nil {
		panic(err)
	} else {
		for obj := range resultChan {
			fmt.Printf("%s\n%#v\n", obj.Type, obj.Object)
		}
		for err := range errChan {
			fmt.Printf("ERROR: %#v", err)
		}
		fmt.Printf("Closed!\n")
	}
}
