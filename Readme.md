# Go Buffer 
[![Build Status](https://travis-ci.org/B-Stefan/go-buffer.svg?branch=master)](https://travis-ci.org/B-Stefan/go-buffer)

A simple wrapper for the [buffer rest api](https://buffer.com/developers/api)

**WIP - Not production ready!**

## Getting started

`go get github.com/B-Stefan/go-buffer`

### Minimal example

````go
package main

import (
	"fmt"
	"github.com/b-stefan/go-buffer/api"
	"log"
)

func main() {

	service := api.NewClient(nil)
	profiles, err := service.Profile.ListProfiles()

	if err != nil{
		fmt.Println("Got error... EOF means authentication failed")
		log.Fatal(err)
	}
	fmt.Println(profiles)
}
````

### OAuth example 

See [example_cli.go](./examole_cli.go)


## Todos 

* Add POST routes for updates
* Enhance error handling 
* Enhance parsing of options (see @Todo)
