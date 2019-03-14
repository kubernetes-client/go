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

package config

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client/kubernetes/config/api"
)

func TestSetUserWithName(t *testing.T) {
	tcs := []struct {
		Origin   []api.NamedAuthInfo
		Name     string
		User     *api.AuthInfo
		Expected []api.NamedAuthInfo
	}{
		{
			Origin: []api.NamedAuthInfo{
				{"A", api.AuthInfo{}},
				{"B", api.AuthInfo{}},
				{"C", api.AuthInfo{}},
			},
			Name: "B",
			User: &api.AuthInfo{Token: "test-token"},
			Expected: []api.NamedAuthInfo{
				{"A", api.AuthInfo{}},
				{"B", api.AuthInfo{Token: "test-token"}},
				{"C", api.AuthInfo{}},
			},
		},
	}

	for _, tc := range tcs {
		if err := setUserWithName(tc.Origin, tc.Name, tc.User); err != nil {
			t.Errorf("unexpected error setting user with name %v: %v", tc.Name, err)
		}

		if !reflect.DeepEqual(tc.Origin, tc.Expected) {
			t.Errorf("setUserWithName mismatch: want %v, got %v", tc.Expected, tc.Origin)
		}
	}
}

func TestGetUserWithName(t *testing.T) {
	users := []api.NamedAuthInfo{
		{"A", api.AuthInfo{}},
		{"B", api.AuthInfo{Token: "test-token"}},
		{"D", api.AuthInfo{}},
		{"D", api.AuthInfo{}},
	}

	tcs := []struct {
		Name          string
		ExpectedUser  *api.AuthInfo
		ExpectedError error
	}{
		{
			Name:          "A",
			ExpectedUser:  &api.AuthInfo{},
			ExpectedError: nil,
		},
		{
			Name:          "B",
			ExpectedUser:  &api.AuthInfo{Token: "test-token"},
			ExpectedError: nil,
		},
		{
			Name:         "C",
			ExpectedUser: nil,
			// A context may have no user, or using non-existing user name.
			// We simply return nil *api.AuthInfo in this case.
			ExpectedError: nil,
		},
		{
			Name:          "D",
			ExpectedUser:  nil,
			ExpectedError: fmt.Errorf("error parsing kube config: duplicate users with name D"),
		},
	}

	for _, tc := range tcs {
		user, err := getUserWithName(users, tc.Name)

		if !reflect.DeepEqual(tc.ExpectedUser, user) {
			t.Errorf("getUserWithName mismatch: want %v, got %v", tc.ExpectedUser, user)
		}

		if !reflect.DeepEqual(tc.ExpectedError, err) {
			t.Errorf("getUserWithName error mismatch: want %v, got %v", tc.ExpectedError, err)
		}
	}
}

func TestGetContextWithName(t *testing.T) {
	contexts := []api.NamedContext{
		{"A", api.Context{}},
		{"B", api.Context{
			Cluster:   "test-cluster",
			AuthInfo:  "test-user",
			Namespace: "test-namespace",
		}},
		{"D", api.Context{}},
		{"D", api.Context{}},
	}

	tcs := []struct {
		Name            string
		ExpectedContext *api.Context
		ExpectedError   error
	}{
		{
			Name:            "A",
			ExpectedContext: &api.Context{},
			ExpectedError:   nil,
		},
		{
			Name: "B",
			ExpectedContext: &api.Context{
				Cluster:   "test-cluster",
				AuthInfo:  "test-user",
				Namespace: "test-namespace",
			},
			ExpectedError: nil,
		},
		{
			Name:            "C",
			ExpectedContext: nil,
			ExpectedError:   fmt.Errorf("error parsing kube config: couldn't find context with name C"),
		},
		{
			Name:            "D",
			ExpectedContext: nil,
			ExpectedError:   fmt.Errorf("error parsing kube config: duplicate contexts with name D"),
		},
	}

	for _, tc := range tcs {
		context, err := getContextWithName(contexts, tc.Name)

		if !reflect.DeepEqual(tc.ExpectedContext, context) {
			t.Errorf("getContextWithName mismatch: want %v, got %v", tc.ExpectedContext, context)
		}

		if !reflect.DeepEqual(tc.ExpectedError, err) {
			t.Errorf("getContextWithName error mismatch: want %v, got %v", tc.ExpectedError, err)
		}
	}
}

func TestGetClusterWithName(t *testing.T) {
	clusters := []api.NamedCluster{
		{"A", api.Cluster{}},
		{"B", api.Cluster{Server: "test-server"}},
		{"D", api.Cluster{}},
		{"D", api.Cluster{}},
	}

	tcs := []struct {
		Name            string
		ExpectedCluster *api.Cluster
		ExpectedError   error
	}{
		{
			Name:            "A",
			ExpectedCluster: &api.Cluster{},
			ExpectedError:   nil,
		},
		{
			Name:            "B",
			ExpectedCluster: &api.Cluster{Server: "test-server"},
			ExpectedError:   nil,
		},
		{
			Name:            "C",
			ExpectedCluster: nil,
			ExpectedError:   fmt.Errorf("error parsing kube config: couldn't find cluster with name C"),
		},
		{
			Name:            "D",
			ExpectedCluster: nil,
			ExpectedError:   fmt.Errorf("error parsing kube config: duplicate clusters with name D"),
		},
	}

	for _, tc := range tcs {
		cluster, err := getClusterWithName(clusters, tc.Name)

		if !reflect.DeepEqual(tc.ExpectedCluster, cluster) {
			t.Errorf("getClusterWithName mismatch: want %v, got %v", tc.ExpectedCluster, cluster)
		}

		if !reflect.DeepEqual(tc.ExpectedError, err) {
			t.Errorf("getClusterWithName error mismatch: want %v, got %v", tc.ExpectedError, err)
		}
	}
}
