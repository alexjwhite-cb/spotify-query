package connect

import (
	"errors"
	"fmt"
	"io"
)

type RequestMethod int

func (r RequestMethod) String() string {
	if _, ok := requestMethods[r]; ok {
		return requestMethods[r]
	}
	return ""
}

const (
	GET RequestMethod = iota
	POST
	PUT
)

var requestMethods = map[RequestMethod]string{
	GET:  "GET",
	POST: "POST",
	PUT:  "PUT",
}

type ResponseType int

func (r ResponseType) String() string {
	if _, ok := responseTypes[r]; ok {
		return responseTypes[r]
	}
	return ""
}

const (
	JSON ResponseType = iota
	PLAIN
)

var responseTypes = map[ResponseType]string{
	JSON:  "application/json",
	PLAIN: "text/plain",
}

type HttpConnector interface {
	SetRequestMethod(method RequestMethod) error
	SetResponseType(responseType ResponseType) error
}

type httpConnection struct {
	requestMethod RequestMethod
	responseType  ResponseType
	request       io.Reader
	response      io.Reader
}

func NewHttpConnector() HttpConnector {
	return &httpConnection{}
}

func (h *httpConnection) SetRequestMethod(method RequestMethod) error {
	if _, ok := requestMethods[method]; !ok {
		return errors.New(fmt.Sprintf("unknown request method: %v", method))
	}
	h.requestMethod = method
	return nil
}

func (h *httpConnection) SetResponseType(responseType ResponseType) error {
	if _, ok := responseTypes[responseType]; !ok {
		return errors.New(fmt.Sprintf("unknown request method: %v", responseType))
	}
	h.responseType = responseType
	return nil
}
