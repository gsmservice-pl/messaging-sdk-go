// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package messagingsdkgo

import (
	"context"
	"fmt"
	"github.com/gsmservice-pl/messaging-sdk-go/internal/hooks"
	"github.com/gsmservice-pl/messaging-sdk-go/internal/utils"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
	"github.com/gsmservice-pl/messaging-sdk-go/retry"
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
	ServerProd:    "https://api.gsmservice.pl/rest",
	ServerSandbox: "https://api.gsmservice.pl/rest-sandbox",
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

// Client - Messaging Gateway GSMService.pl:
// # Introduction
//
// This document was created to explain the process of integration any application or system with the **GSMService.pl** SMS Gateway via the *REST API*. Currently, there are several ways to send messages with the GSMService.pl platform:
//
// * Directly from the [https://bramka.gsmservice.pl](https://bramka.gsmservice.pl) website [User Panel](https://panel.gsmservice.pl)
// * Via this *REST API* and provided *SDKs*
// * Via the legacy (deprecated) versions API: *Webservices (SOAP)* and *HTTP*
// * Via the *MAIL2SMS* service
//
// This document describes the possibilities offered by **REST API**.
//
// > **We kindly ask you to read this documentation carefully before starting the integration. This will make the whole process easier and will help you avoid many problems.**
//
// ## Documentation and Try Outs
//
// This documentation is available in two formats: [**REDOC**](https://api.gsmservice.pl/rest/) and [**SWAGGER**](https://api.gsmservice.pl/rest/swagger). You can test any endpoint directly from documentation using **Try Out** feature of Swagger. Also you can [download a **YAML** file](https://api.gsmservice.pl/rest/swagger/messaging.yaml) with doc in OpenApi 3.0 format.
//
// ## Account signup and setup
//
// Firstly, it is necessary to create your personal account at the GSMService.pl SMS Gateway platform if you haven't one and activate access to the API. To register a new account please [signup the form](https://panel.gsmservice.pl/rejestracja). After signing up and fully activation of an account you have to activate an access to the API.
//
// To do it please use [this site](https://panel.gsmservice.pl/api) - fill the *New API Access* form with your preferred API login, set your API password, select which API standard you want to activate for this account (select **REST API** there). Optionally you can add IP adresses (or IP pool with CIDR notation) from which access to your API account will be possible. You can also set a callback address there to collect any messages status updates automatically. When a status of a messaga changes, the callback address will be called with passing parameters with new message status.
//
// After setup an API access you will get an unique **API Access Token** - please write it down as there won't be possible to display it again (you will have the possibility to regenerate it only). This token will be required to authenticate all the requests made with API on your account.
//
// ## Authentication of API requests
//
// All the endpoints of this REST API have to be authenticated using your API Access Token with one exception: */rest/ping* endpoint which doesn't need an authentication.
//
// To make an authenticated request you should add to all requests an ***Authorization* header** which you have generated in previous step:
//
// ```
// Authorization: Bearer &lt;YOUR_API_ACCESS_TOKEN&gt;
// ```
//
// ## URLs to connect to API
//
// Please use this SSL secured adresses to connect to REST API:
//
// * ```https://api.gsmservice.pl/rest``` - for production system
//
// * ```https://api.gsmservice.pl/rest-sandbox``` - for test system (Sandbox)
//
// > [!NOTE]
// > **When calling our API, make sure you are using TLS version 1.2 or higher. Older versions are no longer supported.**
//
// # SDK Client Libraries
//
// For developers integrating SMS functionality into their app, we provide a convenient SDK Libraries.
//
// Our SDKs allow you to quickly start interacting with the Gateway using your favorite programming language. Currently we support the following languages:
//
// ## PHP 8
//
// To install PHP SDK issue the following command:
//
// ```shell
// composer require gsmservice-pl/messaging-sdk-php
// ```
// More information and documentation you can find at our [GitHub](https://github.com/gsmservice-pl/messaging-sdk-php)
//
// ## Typescript
//
// To install Typescript SDK issue the following command:
//
// ### NPM
//
// ```shell
// npm add @gsmservice-pl/messaging-sdk-typescript
// ```
//
// More information and documentation you can find at our [GitHub](https://github.com/gsmservice-pl/messaging-sdk-typescript)
type Client struct {
	Accounts *Accounts
	Outgoing *Outgoing
	Incoming *Incoming
	// This section describes other usefull operations and tools
	Common  *Common
	Senders *Senders

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
			OpenAPIDocVersion: "0.9.2",
			SDKVersion:        "0.0.4",
			GenVersion:        "2.438.3",
			UserAgent:         "speakeasy-sdk/go 0.0.4 2.438.3 0.9.2 github.com/gsmservice-pl/messaging-sdk-go",
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
