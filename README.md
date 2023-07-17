# go-client

Go libary to ease the integration with the Beem Africa (SMS, AIRTIME, OTP, 2WAY-SMS, BPAY, USSD)

## Installation

To start using this in your project

```bash
go get -u github.com/Jkarage/beemafrica
```

## Authentication

To authenticate and usage of the package add your credentials to your environment.

Or

```golang
 sms := beemafrica.NewSMS()
 sms.ApiKey = "xxxxxx"
 sms.SecretKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

```

### For Unix Based Os

``` bash

export BEEM_API_KEY=<your beem api key>
export BEEM_SECRET_KEY=<your beem secret key>

```

### For Windows

```shell

set BEEM_API_KEY=<your beem api key>
set BEEM_SECRET_KEY=<your beem secret key>

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

 "github.com/Jkarage/beemafrica"
)

func main() {
 sms := beemafrica.NewSMS()
 resp, err := sms.SendSMS("Sample text message", []string{"2557135070XX"}, "")
 if err != nil {
  log.Fatal(err)
 }
 io.Copy(os.Stdout, resp.Body)
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

 "github.com/Jkarage/beemafrica"
)

func main() {
 client := beemafrica.NewSMS()
 resp, err := client.GetBallance()
 if err != nil {
  log.Fatal(err)
 }
 io.Copy(os.Stdout, resp.Body)

 io.Copy(os.Stdout, resp.Body)
}
```

## AIRTIME

### Sending airtime to a friend

``` golang
    client := beemafrica.NewAirtime()
    resp, err := client.Transfer("2557135070XX", 2000, 1234)
    if err != nil {
    log.Fatal(err)
    }
    io.Copy(os.Stdout, resp.Body)
```

### Getting airtime ballance

``` golang
    client := beemafrica.NewAirtime()
    resp, err := client.GetBallance()
    if err != nil {
    log.Fatal(err)
    }
    io.Copy(os.Stdout, resp.Body)
```
