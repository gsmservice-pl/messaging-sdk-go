# Attachments

Attachments for the message. You can pass here images, audio and video files bodies. To set one attachment please use `components.CreateAttachmentsStr()` method simply passing to it a `string` with attachment body encoded with `base64`. To set multiple attachments - please use `components.CreateAttachmentsArrayOfStr()` method passing to it `[]string` with attachment bodies encoded by `base64`. Max 3 attachments per message.


## Supported Types

### 

```go
attachments := components.CreateAttachmentsStr(string{/* values here */})
```

### 

```go
attachments := components.CreateAttachmentsArrayOfStr([]string{/* values here */})
```

