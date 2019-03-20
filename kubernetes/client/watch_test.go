package client

import (
	"strings"
	"testing"
)

func makerFn() interface{} { return &V1Namespace{} }

func TestDecode(t *testing.T) {
	line := `{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}`
	result, err := decode(line, makerFn)
	if err != nil {
		t.Errorf("Unexpeceted non-nil: %v", err)
	}
	ns, ok := result.Object.(*V1Namespace)
	if !ok {
		t.Errorf("Cast failed: %v", result.Object)
	}
	if ns.Kind != "Namespace" || ns.Metadata.Name != "kube-system" {
		t.Errorf("Unexpected value %#v", *ns)
	}
}

func TestMultiDecode(t *testing.T) {
	lines := `
	{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}
	{"type":"MODIFIED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system-2","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}
	{"type":"DELETED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system-3","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}
	`

	results := make(chan *Result)
	errs := make(chan error)
	processWatch(strings.NewReader(lines), makerFn, results, errs)

	out := []*Result{}
	outErrs := []error{}

	for r := range results {
		out = append(out, r)
	}

	for e := range errs {
		outErrs = append(outErrs, e)
	}

	if len(out) != 3 {
		t.Errorf("unexpected results: %v", out)
	}
	for ix, val := range []string{"ADDED", "MODIFIED", "DELETED"} {
		if out[ix].Type != val {
			t.Errorf("unexpected value (%d): %v", ix, out[ix])
		}
	}
}
