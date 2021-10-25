package gofish

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Url    string
	Method string
	Header *http.Header
	Body   io.Reader
	Handle Handle
	Client http.Client
}

func (r *Request) Do() error {
	request, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return err
	}
	request.Header = *r.Header

	resp, err := r.Client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error status code: %d", resp.StatusCode)
	}
	r.Handle.Worker(resp.Body, r.Url)
	defer resp.Body.Close()
	return nil
}

func NewRequest(method, Url, userAgent string, handle Handle, Body io.Reader) (*Request, error) {
	_, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}
	hdr := http.Header{}
	if userAgent != "" {
		hdr.Add("User-Agent", userAgent)
	} else {
		hdr.Add("User-Agent", UserAgent)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
	return &Request{
		Url:    Url,
		Method: method,
		Header: &hdr,
		Body:   Body,
		Handle: handle,
		Client: client,
	}, nil
}
