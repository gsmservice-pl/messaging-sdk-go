# Accounts
(*Accounts*)

## Overview

### Available Operations

* [Get](#get) - Get account details
* [GetSubaccount](#getsubaccount) - Get subaccount details

## Get

Get current account balance and other details of your account. You can check also account limit and if account is main one. Main accounts have unlimited privileges and using [User Panel](https://panel.szybkisms.pl) you can create as many subaccounts as you need.
 
As a successful result a details of current account you are logged in using an API Access Token will be returned.

### Example Usage

```go
package main

import(
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Accounts.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAccountDetailsResponse](../../models/operations/getaccountdetailsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 401, 403, 4XX            | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |

## GetSubaccount

Check account balance and other details such subcredit balance of a subaccount. Subaccounts are additional users who can access your account services and the details. You can restrict access level and setup privileges to subaccounts using [User Panel](https://panel.szybkisms.pl).

This method accepts a `string` type parameter with user login. You should pass there the full subaccount login to access its data.

As a successful result the details of subaccount with provided login will be returned.

### Example Usage

```go
package main

import(
	"context"
	messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go/v3"
	"log"
)

func main() {
    ctx := context.Background()

    s := messagingsdkgo.New(
        messagingsdkgo.WithSecurity("<YOUR API ACCESS TOKEN>"),
    )

    res, err := s.Accounts.GetSubaccount(ctx, "some-login")
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userLogin`                                              | *string*                                                 | :heavy_check_mark:                                       | Login of the subaccount (user) to get the data for       | some-login                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSubaccountDetailsResponse](../../models/operations/getsubaccountdetailsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ErrorResponse  | 401, 403, 404, 4XX       | application/problem+json |
| sdkerrors.ErrorResponse  | 5XX                      | application/problem+json |