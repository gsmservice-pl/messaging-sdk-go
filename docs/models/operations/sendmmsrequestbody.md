# SendMmsRequestBody

To send a single MMS or messages with the same content to multiple recipients, please use `operations.CreateSendMmsRequestBodyMmsMessage()` method with a single `MmsMessage` object with the properties of this message. To send multiple messages with different content at the same time, please use `operations.CreateSendMmsRequestBodyArrayOfMmsMessage()` method passing to it array of type `[]MmsMessage` with the properties of each message.


## Supported Types

### MmsMessage

```go
sendMmsRequestBody := operations.CreateSendMmsRequestBodyMmsMessage(components.MmsMessage{/* values here */})
```

### 

```go
sendMmsRequestBody := operations.CreateSendMmsRequestBodyArrayOfMmsMessage([]components.MmsMessage{/* values here */})
```

