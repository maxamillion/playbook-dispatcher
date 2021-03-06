// Package rbac provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package rbac

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Access defines model for Access.
type Access struct {
	Permission          string               `json:"permission"`
	ResourceDefinitions []ResourceDefinition `json:"resourceDefinitions"`
}

// AccessPagination defines model for AccessPagination.
type AccessPagination struct {
	// Embedded struct due to allOf(#/components/schemas/ListPagination)
	ListPagination `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Data []Access `json:"data"`
}

// Error defines model for Error.
type Error struct {
	Errors []struct {
		Detail *string `json:"detail,omitempty"`
		Status *string `json:"status,omitempty"`
	} `json:"errors"`
}

// ListPagination defines model for ListPagination.
type ListPagination struct {
	Links *PaginationLinks `json:"links,omitempty"`
	Meta  *PaginationMeta  `json:"meta,omitempty"`
}

// PaginationLinks defines model for PaginationLinks.
type PaginationLinks struct {
	First    *string `json:"first,omitempty"`
	Last     *string `json:"last,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

// PaginationMeta defines model for PaginationMeta.
type PaginationMeta struct {
	Count *int64 `json:"count,omitempty"`
}

// ResourceDefinition defines model for ResourceDefinition.
type ResourceDefinition struct {
	AttributeFilter ResourceDefinitionFilter `json:"attributeFilter"`
}

// ResourceDefinitionFilter defines model for ResourceDefinitionFilter.
type ResourceDefinitionFilter struct {
	Key       string `json:"key"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}

// QueryLimit defines model for QueryLimit.
type QueryLimit int

// QueryOffset defines model for QueryOffset.
type QueryOffset int

// GetPrincipalAccessParams defines parameters for GetPrincipalAccess.
type GetPrincipalAccessParams struct {

	// The application name(s) to obtain access for the principal. This is an exact match. When no application is supplied, all permissions for the principal are returned. You may also use a comma-separated list to match on multiple applications.
	Application string `json:"application"`

	// Unique username of the principal to obtain access for (only available for admins, and if supplied, takes precedence over the identity header).
	Username *string `json:"username,omitempty"`

	// Parameter for selecting the amount of data returned.
	Limit *QueryLimit `json:"limit,omitempty"`

	// Parameter for selecting the offset of data.
	Offset *QueryOffset `json:"offset,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetPrincipalAccess request
	GetPrincipalAccess(ctx context.Context, params *GetPrincipalAccessParams) (*http.Response, error)
}

func (c *Client) GetPrincipalAccess(ctx context.Context, params *GetPrincipalAccessParams) (*http.Response, error) {
	req, err := NewGetPrincipalAccessRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewGetPrincipalAccessRequest generates requests for GetPrincipalAccess
func NewGetPrincipalAccessRequest(server string, params *GetPrincipalAccessParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/access/")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "application", params.Application); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Username != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "username", *params.Username); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "limit", *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Offset != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "offset", *params.Offset); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetPrincipalAccess request
	GetPrincipalAccessWithResponse(ctx context.Context, params *GetPrincipalAccessParams) (*GetPrincipalAccessResponse, error)
}

type GetPrincipalAccessResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AccessPagination
	JSON404      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetPrincipalAccessResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPrincipalAccessResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetPrincipalAccessWithResponse request returning *GetPrincipalAccessResponse
func (c *ClientWithResponses) GetPrincipalAccessWithResponse(ctx context.Context, params *GetPrincipalAccessParams) (*GetPrincipalAccessResponse, error) {
	rsp, err := c.GetPrincipalAccess(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetPrincipalAccessResponse(rsp)
}

// ParseGetPrincipalAccessResponse parses an HTTP response from a GetPrincipalAccessWithResponse call
func ParseGetPrincipalAccessResponse(rsp *http.Response) (*GetPrincipalAccessResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPrincipalAccessResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AccessPagination
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
