# OpenAPI based Generated Go client for Kubernetes

This repo hosts an experimental Golang client library generated using [swagger-codegen](https://github.com/swagger-api/swagger-codegen) and [Kubernetes OpenAPI spec](https://github.com/kubernetes/kubernetes/tree/master/api/openapi-spec). Currently the client capability meets the Bronze Requirement and is supported in Alpha state, as described in [Kubernetes: New Client Library Procedure](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/csi-new-client-library-procedure.md). For more advanced functionality (e.g. controller and shared informer), please refer to the more mature Kubernetes [Go client](https://github.com/kubernetes/client-go).

## Installation

Get the source:

```
cd $GOPATH/src/k8s.io
git clone --recursive https://github.com/kubernetes-client/go.git client
cd client
```

**NOTE on `go get`:**

Currently this repo is still under experimental state and the domains haven't been set up for `go get` yet. Please get the library from source.

## Example

Please see https://github.com/kubernetes-client/go/tree/master/examples for
basic examples of using this client library.

## Documentation

All APIs and Models' documentation can be found at the [Generated client's README file](kubernetes/README.md)

## Compatibility

This client library follows [semver](http://semver.org/), so until the major version of
client gets increased, your code will continue to work with explicitly
supported versions of Kubernetes clusters.

#### Compatibility matrix

|                  | Kubernetes 1.10 |
|------------------|-----------------|
| client 0.1.0a1   | ✓               |

Key:

* `✓` Exactly the same features / API objects in both Go client and the Kubernetes
  version.
* `+` Go client has features or api objects that may not be present in the
  Kubernetes cluster, but everything they have in common will work.
* `-` The Kubernetes cluster has features the Go client library can't use
  (additional API objects, etc).

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for instructions on how to contribute.
