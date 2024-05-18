<div align=center>

# 🚀 GoSplash 🚀

[![Apache2.o License](https://img.shields.io/badge/license-Apache2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

<img src="./logo/logo.png" style="width: 200px" alt="">

[Unsplash](https://unsplash.com/) provides freely licensed high-resolution photos that can be used for anything.

This repository is a Golang SDK developed based on the [Unsplash API](https://unsplash.com/documentation#getting-started) content and covers all current APIs.

</div>

## Installation

```bash
  go get github.com/Heyra-Joker/gosplash
```

## Documentation

[Sign up](https://unsplash.com/login) on Unsplash and register as a [developer](https://unsplash.com/developers).
You can then create a new application and use the AppID and Secret for authentication.


## Usage

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/Heyra-Joker/gosplash"
)

func main() {
    clientId := "<YOUR CLIENT ID>"
    g := gosplash.Photo{Client: gosplash.NewClient(clientId)}
    reply, response, err := g.Random(gosplash.PhotoRandomReq{})
    if err != nil {
        log.Fatal(err)
    }
    log.Println(response.OriginResponseBody)
    log.Println(response.RateLimitLimit)
    log.Println(response.RateLimitRemaining)
    data, _ := json.Marshal(reply)
    log.Println(string(data))
}

```


## Examples

| API           | Example                                          |
|---------------|--------------------------------------------------|
| Authorization | [authorization_test.go](./authorization_test.go) |
| Current User  | [current_user_test.go](./current_user_test.go)   |
| Users         | [users_test.go](./users_test.go)                 |
| Photos        | [photo_test.go](./photo_test.go)                 |
| Search        | [search_test.go](./search_test.go)               |
| Collections   | [collection_test.go](./collection_test.go)       |
| Topics        | [topics_test.go](./topics_test.go)               |
| Stats         | [stats_test.go](./stats_test.go)                 |

## FAQ

#### 1. About `UnimplementedUpdate` API

Sorry about this function, I have no permission to upload image, if you have the necessary permissions, you can submit
an issue to fix this function.

#### 2. Why is the API response `Reply` different from the official SDKs in other languages?

This is because the project creates the `Reply` based on actual request results, and the official SDK versions in other languages are really outdated. 
However, rest assured that in most cases, the parsing content you need is already included in `Reply`. 
If you find that some request responses are not included in `Reply`, please submit an `issue` and I will update it as soon as possible.

## License

[Apache2.0 License](https://www.apache.org/licenses/LICENSE-2.0)


