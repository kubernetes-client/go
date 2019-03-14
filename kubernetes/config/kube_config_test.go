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
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"k8s.io/client/kubernetes/config/api"
)

const (
	testData        = "test-data"
	testAnotherData = "test-another-data"

	testServer   = "http://test-server"
	testUsername = "me"
	testPassword = "pass"

	// token for me:pass
	testBasicToken = "Basic bWU6cGFzcw=="

	testSSLServer  = "https://test-server"
	testCertAuth   = "cert-auth"
	testClientKey  = "client-key"
	testClientCert = "client-cert"

	bearerTokenFormat = "Bearer %s"
	testTokenExpiry   = "2000-01-01 12:00:00" // always in past
)

var (
	// base64 encoded string, used as a test token
	testDataBase64 = b64.StdEncoding.EncodeToString([]byte(testData))

	// base64 encoded string, used as another test token
	testAnotherDataBase64 = b64.StdEncoding.EncodeToString([]byte(testAnotherData))

	testCertAuthBase64 = stringToBase64(testCertAuth)

	testClientKeyBase64 = stringToBase64(testClientKey)

	testClientCertBase64 = stringToBase64(testClientCert)

	// test time set to time.Now() + 2 * expirySkewPreventionDelay, which doesn't expire
	testTokenNoExpiry = time.Now().Add(2 * expirySkewPreventionDelay).UTC().Format(gcpRFC3339Format)
)

var testKubeConfig = api.Config{
	CurrentContext: "no_user",
	Contexts: []api.NamedContext{
		{
			Name: "no_user",
			Context: api.Context{
				Cluster: "default",
			},
		},
		{
			Name: "non_existing_user",
			Context: api.Context{
				Cluster:  "default",
				AuthInfo: "non_existing_user",
			},
		},
		{
			Name: "simple_token",
			Context: api.Context{
				Cluster:  "default",
				AuthInfo: "simple_token",
			},
		},
		{
			Name: "gcp",
			Context: api.Context{
				Cluster:  "default",
				AuthInfo: "gcp",
			},
		},
		{
			Name: "expired_gcp",
			Context: api.Context{
				Cluster:  "default",
				AuthInfo: "expired_gcp",
			},
		},
		{
			Name: "user_pass",
			Context: api.Context{
				Cluster:  "default",
				AuthInfo: "user_pass",
			},
		},
		{
			Name: "ssl",
			Context: api.Context{
				Cluster:  "ssl",
				AuthInfo: "ssl",
			},
		},
		{
			Name: "ssl_no_verification",
			Context: api.Context{
				Cluster:  "ssl_no_verification",
				AuthInfo: "ssl",
			},
		},
		{
			Name: "ssl_no_file",
			Context: api.Context{
				Cluster:  "ssl_no_file",
				AuthInfo: "ssl_no_file",
			},
		},
	},
	Clusters: []api.NamedCluster{
		{
			Name: "default",
			Cluster: api.Cluster{
				Server: testServer,
			},
		},
		{
			Name: "ssl",
			Cluster: api.Cluster{
				Server: testSSLServer,
				CertificateAuthorityData: testCertAuthBase64,
			},
		},
		{
			Name: "ssl_no_verification",
			Cluster: api.Cluster{
				Server:                testSSLServer,
				InsecureSkipTLSVerify: true,
			},
		},
		{
			Name: "ssl_no_file",
			Cluster: api.Cluster{
				Server:               testSSLServer,
				CertificateAuthority: "test-cert-no-file",
			},
		},
	},
	AuthInfos: []api.NamedAuthInfo{
		{
			Name: "simple_token",
			AuthInfo: api.AuthInfo{
				Token:    testDataBase64,
				Username: testUsername,
				Password: testPassword,
			},
		},
		{
			Name: "gcp",
			AuthInfo: api.AuthInfo{
				AuthProvider: &api.AuthProviderConfig{
					Name: "gcp",
					Config: map[string]string{
						"access-token": testDataBase64,
						"expiry":       testTokenNoExpiry,
					},
				},
				Token:    testDataBase64,
				Username: testUsername,
				Password: testPassword,
			},
		},
		{
			Name: "expired_gcp",
			AuthInfo: api.AuthInfo{
				AuthProvider: &api.AuthProviderConfig{
					Name: "gcp",
					Config: map[string]string{
						"access-token": testDataBase64,
						"expiry":       testTokenExpiry,
					},
				},
				Token:    testDataBase64,
				Username: testUsername,
				Password: testPassword,
			},
		},
		{
			Name: "user_pass",
			AuthInfo: api.AuthInfo{
				Username: testUsername,
				Password: testPassword,
			},
		},
		{
			Name: "ssl",
			AuthInfo: api.AuthInfo{
				Token: testDataBase64,
				ClientCertificateData: testClientCertBase64,
				ClientKeyData:         testClientKeyBase64,
			},
		},
		{
			Name: "ssl_no_file",
			AuthInfo: api.AuthInfo{
				Token:             testDataBase64,
				ClientCertificate: "test-client-cert-no-file",
				ClientKey:         "test-client-key-no-file",
			},
		},
	},
}

func TestLoadKubeConfig(t *testing.T) {
	tcs := []struct {
		ActiveContext string

		Server        string
		Token         string
		CACert        []byte
		Cert          []byte
		Key           []byte
		SkipTLSVerify bool
		GCLoader      GoogleCredentialLoader
	}{
		{
			ActiveContext: "no_user",
			Server:        testServer,
		},
		{
			ActiveContext: "non_existing_user",
			Server:        testServer,
		},
		{
			ActiveContext: "simple_token",
			Server:        testServer,
			Token:         fmt.Sprintf(bearerTokenFormat, testDataBase64),
		},
		{
			ActiveContext: "user_pass",
			Server:        testServer,
			Token:         testBasicToken,
		},
		{
			ActiveContext: "gcp",
			Server:        testServer,
			Token:         fmt.Sprintf(bearerTokenFormat, testDataBase64),
			GCLoader:      FakeGoogleCredentialLoaderNoRefresh{},
		},
		{
			ActiveContext: "expired_gcp",
			Server:        testServer,
			Token:         fmt.Sprintf(bearerTokenFormat, testAnotherDataBase64),
			GCLoader:      FakeGoogleCredentialLoader{},
		},
		{
			ActiveContext: "ssl",
			Server:        testSSLServer,
			Token:         fmt.Sprintf(bearerTokenFormat, testDataBase64),
			CACert:        testCertAuthBase64,
			Cert:          testClientCertBase64,
			Key:           testClientKeyBase64,
		},
		{
			ActiveContext: "ssl_no_verification",
			Server:        testSSLServer,
			Token:         fmt.Sprintf(bearerTokenFormat, testDataBase64),
			Cert:          testClientCertBase64,
			Key:           testClientKeyBase64,
			SkipTLSVerify: true,
		},
	}

	for _, tc := range tcs {
		expected, err := FakeConfig(tc.Server, tc.Token, tc.CACert, tc.Cert, tc.Key, tc.SkipTLSVerify)
		if err != nil {
			t.Errorf("context %v, unexpected error setting up fake config: %v", tc.ActiveContext, err)
		}

		actual := KubeConfigLoader{
			rawConfig:         testKubeConfig,
			skipConfigPersist: true,
			gcLoader:          tc.GCLoader,
		}
		if err := actual.SetActiveContext(tc.ActiveContext); err != nil {
			t.Errorf("context %v, unexpected error setting config active context: %v", tc.ActiveContext, err)
		}

		// We are only testing loading auth and TLS info in LoadAndSet; we are not testing setting
		// the generate client's Configuration based on the restConfig, because we are using fake
		// data as TLS cert, which would fail PEM validation
		actual.loadAuthentication()
		if err := actual.loadClusterInfo(); err != nil {
			t.Errorf("context %v, unexpected error loading kube config: %v", tc.ActiveContext, err)
		}
		if !reflect.DeepEqual(expected, actual.RestConfig()) {
			t.Errorf("context %v, config loaded mismatch: want %v, got %v", tc.ActiveContext, expected, actual.RestConfig())
		}
	}
}

func TestLoadKubeConfigSSLNoFile(t *testing.T) {
	actual := KubeConfigLoader{
		rawConfig:         testKubeConfig,
		skipConfigPersist: true,
	}
	if err := actual.SetActiveContext("ssl_no_file"); err != nil {
		t.Errorf("context %v, unexpected error setting config active context: %v", "ssl_no_file", err)
	}

	// We are only testing loading auth and TLS info in LoadAndSet; we are not testing setting
	// the generate client's Configuration based on the restConfig, because we are using fake
	// data as TLS cert, which would fail PEM validation
	actual.loadAuthentication()
	if err := actual.loadClusterInfo(); err == nil || !strings.Contains(err.Error(), "failed to get data or file") {
		t.Errorf("context %v, expecting failure to get file, got: %v", "ssl_no_file", err)
	}
}

func TestLoadKubeConfigSSLLocalFile(t *testing.T) {
	tc := struct {
		ActiveContext string

		Server        string
		Token         string
		CACert        []byte
		Cert          []byte
		Key           []byte
		SkipTLSVerify bool
	}{

		ActiveContext: "ssl_local_file",
		Server:        testSSLServer,
		Token:         fmt.Sprintf(bearerTokenFormat, testDataBase64),
		CACert:        testCertAuthBase64,
		Cert:          testClientCertBase64,
		Key:           testClientKeyBase64,
	}

	expected, err := FakeConfig(tc.Server, tc.Token, tc.CACert, tc.Cert, tc.Key, tc.SkipTLSVerify)
	if err != nil {
		t.Errorf("context %v, unexpected error setting up fake config: %v", tc.ActiveContext, err)
	}

	// Set up CA cert file
	testCACertFile, err := ioutil.TempFile(os.TempDir(), "ca-cert")
	if err != nil {
		t.Errorf("error: failed to create temp ca-cert file")
	}
	defer os.Remove(testCACertFile.Name())
	if err := ioutil.WriteFile(testCACertFile.Name(), testCertAuthBase64, 0644); err != nil {
		t.Errorf("context %v, unexpected error writing temp file %v: %v", tc.ActiveContext, testCACertFile.Name(), err)
	}

	// Set up token file
	testTokenFile, err := ioutil.TempFile(os.TempDir(), "token")
	if err != nil {
		t.Errorf("error: failed to create temp token file")
	}
	defer os.Remove(testTokenFile.Name())
	if err := ioutil.WriteFile(testTokenFile.Name(), []byte(testDataBase64), 0644); err != nil {
		t.Errorf("context %v, unexpected error writing temp file %v: %v", tc.ActiveContext, testTokenFile.Name(), err)
	}

	// Set up client cert file
	testClientCertFile, err := ioutil.TempFile(os.TempDir(), "client-cert")
	if err != nil {
		t.Errorf("error: failed to create temp client-cert file")
	}
	defer os.Remove(testClientCertFile.Name())
	if err := ioutil.WriteFile(testClientCertFile.Name(), testClientCertBase64, 0644); err != nil {
		t.Errorf("context %v, unexpected error writing temp file %v: %v", tc.ActiveContext, testClientCertFile.Name(), err)
	}

	// Set up client key file
	testClientKeyFile, err := ioutil.TempFile(os.TempDir(), "client-key")
	if err != nil {
		t.Errorf("error: failed to create temp client-key file")
	}
	defer os.Remove(testClientKeyFile.Name())
	if err := ioutil.WriteFile(testClientKeyFile.Name(), testClientKeyBase64, 0644); err != nil {
		t.Errorf("context %v, unexpected error writing temp file %v: %v", tc.ActiveContext, testClientKeyFile.Name(), err)
	}

	actual := KubeConfigLoader{
		rawConfig: api.Config{
			CurrentContext: "ssl_local_file",
			Contexts: []api.NamedContext{
				{
					Name: "ssl_local_file",
					Context: api.Context{
						Cluster:  "ssl_local_file",
						AuthInfo: "ssl_local_file",
					},
				},
			},
			Clusters: []api.NamedCluster{
				{
					Name: "ssl_local_file",
					Cluster: api.Cluster{
						Server:               testSSLServer,
						CertificateAuthority: testCACertFile.Name(),
					},
				},
			},
			AuthInfos: []api.NamedAuthInfo{
				{
					Name: "ssl_local_file",
					AuthInfo: api.AuthInfo{
						TokenFile:         testTokenFile.Name(),
						ClientCertificate: testClientCertFile.Name(),
						ClientKey:         testClientKeyFile.Name(),
					},
				},
			},
		},
		skipConfigPersist: true,
	}
	if err := actual.SetActiveContext(tc.ActiveContext); err != nil {
		t.Errorf("context %v, unexpected error setting config active context: %v", tc.ActiveContext, err)
	}

	// We are only testing loading auth and TLS info in LoadAndSet; we are not testing setting
	// the generate client's Configuration based on the restConfig, because we are using fake
	// data as TLS cert, which would fail PEM validation
	actual.loadAuthentication()
	if err := actual.loadClusterInfo(); err != nil {
		t.Errorf("context %v, unexpected error loading kube config: %v", tc.ActiveContext, err)
	}
	if !reflect.DeepEqual(expected, actual.RestConfig()) {
		t.Errorf("context %v, config loaded mismatch: want %v, got %v", tc.ActiveContext, expected, actual.RestConfig())
	}
}

func FakeConfig(server, token string, caCert, clientCert, clientKey []byte, skipTLSVerify bool) (RestConfig, error) {
	u, err := url.Parse(server)
	if err != nil {
		return RestConfig{}, err
	}

	return RestConfig{
		basePath:      strings.TrimRight(server, "/"),
		host:          u.Host,
		scheme:        u.Scheme,
		token:         token,
		caCert:        caCert,
		clientCert:    clientCert,
		clientKey:     clientKey,
		skipTLSVerify: skipTLSVerify,
	}, nil
}

func stringToBase64(str string) []byte {
	return []byte(b64.StdEncoding.EncodeToString([]byte(str)))
}

type FakeGoogleCredentialLoader struct{}

func (l FakeGoogleCredentialLoader) GetGoogleCredentials() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: testAnotherDataBase64, Expiry: time.Now().UTC()}, nil
}

type FakeGoogleCredentialLoaderNoRefresh struct{}

func (l FakeGoogleCredentialLoaderNoRefresh) GetGoogleCredentials() (*oauth2.Token, error) {
	return nil, fmt.Errorf("should not be called")
}
