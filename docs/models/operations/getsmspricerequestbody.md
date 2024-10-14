# GetSmsPriceRequestBody

To check the price of a single SMS or messages with the same content to multiple recipients, pass in the Request Body a single `Sms` object with the properties of this message. To check the price of multiple messages with different content at the same time, pass in the Request Body an `array` of `Sms` objects with the properties of each message.


## Supported Types

### Sms

```go
getSmsPriceRequestBody := operations.CreateGetSmsPriceRequestBodySms(components.Sms{/* values here */})
```

### 

```go
getSmsPriceRequestBody := operations.CreateGetSmsPriceRequestBodyArrayOfSms([]components.Sms{/* values here */})
```

