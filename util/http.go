package util

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpOption struct {
	Client *http.Client
	Ctx    context.Context
	Body   io.Reader
	Header http.Header
}

func DoHttpRequest(url, method string, resp interface{}, option HttpOption) error {
	var (
		respRaw jsoniter.RawMessage
		err     error
	)
	if respRaw, err = doHttpRequest(url, method, option); err != nil {
		return err
	}
	if err = jsoniter.Unmarshal(respRaw, resp); err != nil {
		return err
	}
	return nil
}

func doHttpRequest(url, method string, option HttpOption) (jsoniter.RawMessage, error) {
	var (
		err    error
		ctx    context.Context
		client *http.Client
		req    *http.Request
		resp   *http.Response
	)
	if option.Ctx != nil {
		ctx = option.Ctx
	} else {
		ctx = context.TODO()
	}
	if http.MethodGet == method {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	} else if http.MethodPost == method {
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, option.Body)
	} else {
		return nil, errors.New("unimplemented method")
	}
	if err != nil {
		return nil, errors.New("failed to build req")
	}
	if option.Header != nil {
		req.Header = option.Header
	}
	if option.Client != nil {
		client = option.Client
	} else {
		client = http.DefaultClient
	}
	if resp, err = client.Do(req); err != nil {
		return nil, errors.Wrap(err, "failed to do req")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("resp status code error. code: " + string(resp.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read resp body")
	}
	return responseBody, nil
}
