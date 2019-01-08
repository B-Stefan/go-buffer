# Go Buffer 
[![Build Status](https://travis-ci.org/B-Stefan/go-buffer.svg?branch=master)](https://travis-ci.org/B-Stefan/go-buffer)
[![Maintainability](https://api.codeclimate.com/v1/badges/8391f026b21e3a252567/maintainability)](https://codeclimate.com/github/B-Stefan/go-buffer/maintainability)

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
		fmt.Println("Got error...")
		log.Fatal(err)
	}
	fmt.Println(profiles)
}
````

### OAuth example 

See [example.go](./example.go)


## Todos 

* [#2 Enhance parsing of options (see @Todo)](https://github.com/B-Stefan/go-buffer/issues/2)
* [#1 - Add POST routes for updates](https://github.com/B-Stefan/go-buffer/issues/1)
* Enhance error handling 
