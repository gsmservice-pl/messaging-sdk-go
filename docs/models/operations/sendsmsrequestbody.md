# SendSmsRequestBody

To send a single SMS or messages with the same content to multiple recipients, please use `operations.CreateGetSmsPriceRequestBodySmsMessage()` method with a single `SmsMessage` object with the properties of this message. To send multiple messages with different content at the same time, please use `operations.CreateSendSmsRequestBodyArrayOfSmsMessage()` method passing to it array of type `[]SmsMessage` with the properties of each message.


## Supported Types

### SmsMessage

```go
sendSmsRequestBody := operations.CreateSendSmsRequestBodySmsMessage(components.SmsMessage{/* values here */})
```

### 

```go
sendSmsRequestBody := operations.CreateSendSmsRequestBodyArrayOfSmsMessage([]components.SmsMessage{/* values here */})
```

