[![Go Reference](https://pkg.go.dev/badge/github.com/gsmservice-pl/messaging-sdk-go.svg)](https://pkg.go.dev/github.com/gsmservice-pl/messaging-sdk-go)
[![GitHub Release](https://img.shields.io/github/v/release/gsmservice-pl/messaging-sdk-go)](https://github.com/gsmservice-pl/messaging-sdk-go)
[![GitHub License](https://img.shields.io/github/license/gsmservice-pl/messaging-sdk-go)](https://github.com/gsmservice-pl/messaging-sdk-go/blob/main/LICENSE)
[![Static Badge](https://img.shields.io/badge/built_by-Speakeasy-yellow)](https://www.speakeasy.com/?utm_source=github-com/gsmservice-pl/messaging-sdk-go&utm_campaign=go)
# GSMService.pl Messaging REST API SDK for Go

This package includes Messaging SDK for Go to send SMS & MMS messages directly from your app via https://bramka.gsmservice.pl messaging platform.

## Additional documentation:

A documentation of all methods and types is available below in section [Available Resources and Operations
](#available-resources-and-operations).

Also you can refer to the [REST API documentation](https://api.gsmservice.pl/rest/) for additional details about the basics of this SDK.
<!-- No Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/gsmservice-pl/messaging-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Sending single SMS Message

This example demonstrates simple sending SMS message to a single recipient:

```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"github.com/gsmservice-pl/messaging-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Outgoing.Sms.Send(ctx, operations.CreateSendSmsRequestBodyArrayOfSms(
		[]components.Sms{
			components.Sms{
				Recipients: components.CreateRecipientsArrayOfStr(
					[]string{
						"+48999999999",
					},
				),
				Message: "To jest treść wiadomości",
				Sender:  messagingsdkgo.String("Bramka SMS"),
				Type:    components.SmsTypeSmsPro.ToPointer(),
				Unicode: messagingsdkgo.Bool(true),
				Flash:   messagingsdkgo.Bool(false),
				Date:    nil,
			},
		},
	))
	if err != nil {
		log.Fatal(err)
	}
	if res.Messages != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Accounts](docs/sdks/accounts/README.md)

* [Get](docs/sdks/accounts/README.md#get) - Get account details
* [GetSubaccount](docs/sdks/accounts/README.md#getsubaccount) - Get subaccount details


### [Common](docs/sdks/common/README.md)

* [Ping](docs/sdks/common/README.md#ping) - Checks API availability and version

### [Incoming](docs/sdks/incoming/README.md)

* [List](docs/sdks/incoming/README.md#list) - List the received SMS messages
* [GetByIds](docs/sdks/incoming/README.md#getbyids) - Get the incoming messages by IDs

### [Outgoing](docs/sdks/outgoing/README.md)

* [GetByIds](docs/sdks/outgoing/README.md#getbyids) - Get the messages details and status by IDs
* [CancelScheduled](docs/sdks/outgoing/README.md#cancelscheduled) - Cancel a scheduled messages
* [List](docs/sdks/outgoing/README.md#list) - Lists the history of sent messages

#### [Outgoing.Sms](docs/sdks/sms/README.md)

* [GetPrice](docs/sdks/sms/README.md#getprice) - Check the price of SMS Messages
* [Send](docs/sdks/sms/README.md#send) - Send SMS Messages

### [Senders](docs/sdks/senders/README.md)

* [List](docs/sdks/senders/README.md#list) - List allowed senders names
* [Add](docs/sdks/senders/README.md#add) - Add a new sender name
* [Delete](docs/sdks/senders/README.md#delete) - Delete a sender name
* [SetDefault](docs/sdks/senders/README.md#setdefault) - Set default sender name

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"github.com/gsmservice-pl/messaging-sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"github.com/gsmservice-pl/messaging-sdk-go/retry"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 401, 403, 4XX, 5XX       | application/problem+json |

### Example

```go
package main

import (
	"context"
	"errors"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"github.com/gsmservice-pl/messaging-sdk-go/models/sdkerrors"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx)
	if err != nil {

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `prod` | `https://api.gsmservice.pl/rest` | None |
| `sandbox` | `https://api.gsmservice.pl/rest-sandbox` | None |

#### Example

```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithServer("sandbox"),
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithServerURL("https://api.gsmservice.pl/rest"),
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                 | Type                 | Scheme               | Environment Variable |
| -------------------- | -------------------- | -------------------- | -------------------- |
| `Bearer`             | http                 | HTTP Bearer          | `GATEWAY_API_BEARER` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
	"log"
	"os"
)

func main() {
	s := messagingsdkgo.New(
		messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
	)

	ctx := context.Background()
	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in continuous development and there may be breaking changes between a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.
