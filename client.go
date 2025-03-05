// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package messagingsdkgo

import (
	"context"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/internal/hooks"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/internal/utils"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/v3/retry"
	"net/http"
	"time"
)

const (
	// Production system
	ServerProd string = "prod"
	// Test system (SANDBOX)
	ServerSandbox string = "sandbox"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd:    "https://api.szybkisms.pl/rest",
	ServerSandbox: "https://api.szybkisms.pl/rest-sandbox",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "prod"
	}

	return ServerList[c.Server], nil
}

// Client - Messaging Gateway SzybkiSMS.pl (former GSMService.pl): This package includes Messaging SDK for GO to send SMS and MMS messages directly from your app via [https://szybkisms.pl](https://szybkisms.pl) messaging platform.
//
// To initialize SDK environment please use this syntax:
//
// ```
// import (
//
//	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
//	"os"
//
// )
//
// s := messagingsdkgo.New(
//
//	messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
//
// )
// ```
//
// If you want to use a Sandbox test system please initialize it as follows:
//
// ```
// s := messagingsdkgo.New(
//
//	messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
//	messagingsdkgo.WithServer(messagingsdkgo.ServerSandbox),
//
// )
// ```
//
// https://szybkisms.pl - SzybkiSMS.pl
type Client struct {
	Accounts *Accounts
	Outgoing *Outgoing
	Incoming *Incoming
	Common   *Common
	Senders  *Senders

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Client)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Client) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Client) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *Client) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Client) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(bearer string) SDKOption {
	return func(sdk *Client) {
		security := components.Security{Bearer: &bearer}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Client) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Client) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Client) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Client {
	sdk := &Client{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.2.1",
			SDKVersion:        "3.0.1",
			GenVersion:        "2.539.1",
			UserAgent:         "speakeasy-sdk/go 3.0.1 2.539.1 1.2.1 github.com/gsmservice-pl/messaging-sdk-go",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Accounts = newAccounts(sdk.sdkConfiguration)

	sdk.Outgoing = newOutgoing(sdk.sdkConfiguration)

	sdk.Incoming = newIncoming(sdk.sdkConfiguration)

	sdk.Common = newCommon(sdk.sdkConfiguration)

	sdk.Senders = newSenders(sdk.sdkConfiguration)

	return sdk
}
