# go-client

Go libary to ease the integration with the Beem Africa (SMS, AIRTIME, OTP, 2WAY-SMS, BPAY, USSD)

## Installation

To start using this in your project

```bash
go get -u github.com/Jkarage/go-client
```

## Authentication

To authenticate and usage of the package add your credentials to your environment.

Or

```golang

    client := sms.New()
    client.Apikey = <Your beem api key>
    client.SecretKey = <Your beem secret Key>
    client.SendSMS("Beeming beem with go", []string{"2557135070XX"}, "")
```

### For Unix Based Os

``` bash

export BEEM_SMS_API_KEY=<your beem api key>
export BEEM_SMS_SECRET_KEY=<your beem secret key>

```

### For Windows

```shell

set BEEM_SMS_API_KEY=<your beem api key>
set BEEM_SMS_SECRET_KEY=<your beem secret key>

```

## Sending SMS with beemAfrica

You can now send sms to multiple recipients with this package by just adding your credentials to the environment.

### Usage Example

``` golang
package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/go-client/sms"
)

func main() {
 client := sms.New()

 resp, err := client.SendSMS("Hello from Beem and Golang", []string{"2557135070XX"}, "")
 if err != nil {
  log.Fatal(err)
 }

 _, err = io.Copy(os.Stdout, resp.Body)
}

```

## Requesting the sms Ballance from beemAfrica

The package provides an api for requesting the sms ballance in a specific account.

### Usage Example

``` golang
    package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/go-client/sms"
)

func main() {
 client := sms.New()
 resp, err := client.GetBallance()
 if err != nil {
  log.Fatal(err)
 }

 io.Copy(os.Stdout, resp.Body)
}
```
