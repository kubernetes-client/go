package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func makerFn() interface{} { return &V1Namespace{} }

type staticHandler struct {
	Code        int
	Body        string
	QueryParams url.Values
}

func (s *staticHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(s.Code)
	res.Write([]byte(s.Body))
	s.QueryParams = req.URL.Query()
}

func TestFullError(t *testing.T) {
	server := httptest.NewServer(&staticHandler{
		Code: 404,
		Body: `{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}\n`,
	})
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	cfg := &Configuration{}
	cfg.Host = u.Host
	cfg.Scheme = u.Scheme

	watch := WatchClient{
		Cfg:     cfg,
		Client:  NewAPIClient(cfg),
		MakerFn: makerFn,
	}

	if _, _, err := watch.Connect(context.Background(), ""); err == nil {
		t.Error("unexpected nil error")
	}
}

func TestFull(t *testing.T) {
	server := httptest.NewServer(&staticHandler{
		Code: 200,
		Body: `{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}\n`,
	})
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	cfg := &Configuration{}
	cfg.Host = u.Host
	cfg.Scheme = u.Scheme

	watch := WatchClient{
		Cfg:     cfg,
		Client:  NewAPIClient(cfg),
		MakerFn: makerFn,
	}

	resultChan, errChan, err := watch.Connect(context.Background(), "")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	out := []*Result{}
	outErrs := []error{}

	for r := range resultChan {
		out = append(out, r)
	}

	for e := range errChan {
		outErrs = append(outErrs, e)
	}

	if len(out) != 1 {
		t.Errorf("unexpected results: %v", out)
		t.FailNow()
	}
	for ix, val := range []string{"ADDED"} {
		if out[ix].Type != val {
			t.Errorf("unexpected value (%d): %v", ix, out[ix])
		}
	}
	if len(outErrs) != 0 {
		t.Errorf("unexpected errors: %v", outErrs)
	}
}

func TestResourceVersion(t *testing.T) {
	handler := &staticHandler{
		Code: 200,
		Body: `{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}\n`,
	}
	server := httptest.NewServer(handler)
	defer server.Close()

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	cfg := &Configuration{}
	cfg.Host = u.Host
	cfg.Scheme = u.Scheme

	watch := WatchClient{
		Cfg:     cfg,
		Client:  NewAPIClient(cfg),
		MakerFn: makerFn,
	}

	version := "12345"
	_, _, err = watch.Connect(context.Background(), version)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	versionOut := handler.QueryParams["resourceVersion"][0]
	if versionOut != version {
		t.Errorf("unexpected resource version %s vs %s", version, versionOut)
	}
}

func TestDecode(t *testing.T) {
	line := `{"type":"ADDED","object":{"kind":"Namespace","apiVersion":"v1","metadata":{"name":"kube-system","selfLink":"/api/v1/namespaces/kube-system","uid":"164931a7-3d75-11e9-a0a0-2683b9459238","resourceVersion":"227","creationTimestamp":"2019-03-03T05:27:50Z","annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"}},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}}}`
	result, err := decode(line, makerFn)
	if err != nil {
		t.Errorf("Unexpected non-nil: %v", err)
	}
	if result.Type != "ADDED" {
		t.Errorf("Unexpected event type: %s expected 'ADDED'", result.Type)
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
	if len(outErrs) != 0 {
		t.Errorf("unexpected errors: %v", outErrs)
	}
}
