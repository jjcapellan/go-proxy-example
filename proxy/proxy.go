package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	proxy *httputil.ReverseProxy
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.proxy.ServeHTTP(w, r)
}

var BackendProxy Proxy

func Init() error {
	proxy, err := newProxy("http://localhost:3002/")
	if err != nil {
		return err
	}
	BackendProxy = Proxy{proxy}
	return nil
}

func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}
