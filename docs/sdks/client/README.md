# Client SDK

## Overview

Messaging Gateway GSMService.pl: This package includes Messaging SDK for GO to send SMS and MMS messages directly from your app via [https://bramka.gsmservice.pl](https://bramka.gsmservice.pl) messaging platform.

To initialize SDK environment please use this syntax:

```
import (
   messagingsdkgo "github.com/gsmservice-pl/messaging-sdk-go"
   "os"
)

s := messagingsdkgo.New(
   messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
)
```

If you want to use a Sandbox test system please initialize it as follows:

```
s := messagingsdkgo.New(
   messagingsdkgo.WithSecurity(os.Getenv("GATEWAY_API_BEARER")),
   messagingsdkgo.WithServer(messagingsdkgo.ServerSandbox),
)
```

Bramka GSMService.pl
<https://bramka.gsmservice.pl>