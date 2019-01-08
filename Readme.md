# Go Buffer 
[![Build Status](https://travis-ci.org/B-Stefan/go-buffer.svg?branch=master)](https://travis-ci.org/B-Stefan/go-buffer)
[![Maintainability](https://api.codeclimate.com/v1/badges/8391f026b21e3a252567/maintainability)](https://codeclimate.com/github/B-Stefan/go-buffer/maintainability)

A simple wrapper to simplify and type the awesome [buffer rest api](https://buffer.com/developers/api)

*You encountered a bug or want to improve this module? 
Wow, great let's improve together! I'm happy about every issue or merge request! 👍*

**Known issue with oAuth see [#4](https://github.com/B-Stefan/go-buffer/issues/4)**

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


## Todos / Learnings

* Fist go project - go is fun to lean! 👍
* This this API with real return values. **Blocked by [#4](https://github.com/B-Stefan/go-buffer/issues/4)**
* Enhance error handling according to docs 
* Use composition for response / options types? (Learn how to use composition in go 😉)
* Test more unmarshal / marshal or create a integration test with real api (favored)