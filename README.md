# zenduty-go-sdk
 Zenduty API client in Go, primarily used by [Zenduty](https://github.com/Zenduty/zenduty-go-sdk) provider in Terraform.
 
 ## Installation
```bash
go get github.com/Zenduty/zenduty-go-sdk
```


## Getting started
Before you begin making use of the SDK, make sure you have your Zenduty Access Token.

```
import "github.com/Zenduty/zenduty-go-sdk/client"
```
Configure the Token and Url

```
config := &client.Config{
	Token: "", // enter token for authentication

	}
```
Based on the service you want to communicate with,Create object for required class,For example, to create Team



## Example usage
```go
package main

import (
	"fmt"

	"github.com/Zenduty/zenduty-go-sdk/client"
)


func main() {
	config := &client.Config{
		Token: "",
	}
	c, err := client.NewClient(config)
	if err != nil {
		panic(err)
	}

	newteam := &client.CreateTeams{}
	newteam.Name = "demoteam"

	resp, err := c.Teams.CreateTeam(newteam)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)

}


```

