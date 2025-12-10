# GetSmsPriceRequestBody

To check the price of a single message or messages with the same content to multiple recipients, pass a single `SmsMessage` object with the properties of this message using `operations.CreateGetSmsPriceRequestBodySmsMessage()` method. To check the price of multiple messages with different content at the same time, pass a `[]SmsMessage` array with the properties of each message using `operations.CreateGetSmsPriceRequestBodyArrayOfSmsMessage()` method.


## Supported Types

### SmsMessage

```go
getSmsPriceRequestBody := operations.CreateGetSmsPriceRequestBodySmsMessage(components.SmsMessage{/* values here */})
```

### 

```go
getSmsPriceRequestBody := operations.CreateGetSmsPriceRequestBodyArrayOfSmsMessage([]components.SmsMessage{/* values here */})
```

