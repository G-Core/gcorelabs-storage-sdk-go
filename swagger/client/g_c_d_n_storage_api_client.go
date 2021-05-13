// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/key"
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/location"
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/statistics"
	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/client/storage"
)

// Default g c d n storage API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "storage.gcorelabs.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new g c d n storage API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GCDNStorageAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new g c d n storage API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GCDNStorageAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new g c d n storage API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GCDNStorageAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GCDNStorageAPI)
	cli.Transport = transport
	cli.Key = key.New(transport, formats)
	cli.Location = location.New(transport, formats)
	cli.Statistics = statistics.New(transport, formats)
	cli.Storage = storage.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GCDNStorageAPI is a client for g c d n storage API
type GCDNStorageAPI struct {
	Key key.ClientService

	Location location.ClientService

	Statistics statistics.ClientService

	Storage storage.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GCDNStorageAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Key.SetTransport(transport)
	c.Location.SetTransport(transport)
	c.Statistics.SetTransport(transport)
	c.Storage.SetTransport(transport)
}