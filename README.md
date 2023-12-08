# beemafrica-go-client

Go libary to ease the integration with the Beem Africa (SMS, AIRTIME, OTP, 2WAY-SMS, BPAY, USSD)

## Installation

To start using this in your project

```bash
# Download this project
go get  github.com/Jkarage/beemafrica@v1.0.1
```

## Authentication

To authenticate and usage of the package add your beemafrica API credentials to your environment.

### For Unix Based

```bash
# Add API and Secret Key to your Environment
export BEEM_API_KEY=<your beem api key>
export BEEM_SECRET_KEY=<your beem secret key>
```

### For Windows

``` code
# Add API and Secret Key to your Environment
set BEEM_API_KEY=<your beem api key>
set BEEM_SECRET_KEY=<your beem secret key>
```

## Run Tests with one command

``` bash
    # Run tests
    go test ./... -v
```

## Send sms sample

The package wraps the send sms functions within it, now you can use to send the sms.

``` golang
package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/beemafrica"
)

func main() {
// start sms client and send sms with it.
 client := beemafrica.NewSMS()
 resp, err := sms.SendSMS("Sample text message", []string{"2557135070XX"}, "")
 if err != nil {
  log.Fatal(err)
 }
 io.Copy(os.Stdout, resp.Body)
}
```

## Request ballance

The package provides an api for requesting the sms ballance in a specific account.

### SMS Usage Example

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

### Send airtime sample

``` golang
package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/beemafrica"
)
    client := beemafrica.NewAirtime()
    resp, err := client.Transfer("2557135070XX", 2000, 1234)
    if err != nil {
    log.Fatal(err)
    }
    io.Copy(os.Stdout, resp.Body)
```

### Getting airtime ballance

``` golang
package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/beemafrica"
)
    client := beemafrica.NewAirtime()
    resp, err := client.GetBallance()
    if err != nil {
    log.Fatal(err)
    }
    io.Copy(os.Stdout, resp.Body)
```

## OTP

### Requesting an OTP PIN

``` golang
    package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/beemafrica"
)
    client := beemafrica.NewOTP()
    resp, err := client.Request("2557135070XX", 12XX)
    if err != nil {
    log.Fatal(err)
    }

    io.Copy(os.Stdout, resp.Body)
```

### Verifying an OTP PIN

``` golang
package main

import (
 "io"
 "log"
 "os"

 "github.com/Jkarage/beemafrica"
)
client := beemafrica.NewOTP()
 resp, err := client.Verify("44bcae75-15ff-4885-915c-6eeba6xxxxx", "57XXXX")
 if err != nil {
  log.Fatal(err)
 }

 io.Copy(os.Stdout, resp.Body)
```
