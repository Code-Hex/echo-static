echo-static
====
[![Build Status](https://travis-ci.org/Code-Hex/echo-static.svg?branch=master)](https://travis-ci.org/Code-Hex/echo-static)  
[![Coverage Status](https://coveralls.io/repos/github/Code-Hex/echo-static/badge.svg?branch=master)](https://coveralls.io/github/Code-Hex/echo-static?branch=master)  
[![Go Report Card](https://goreportcard.com/badge/github.com/Code-Hex/echo-static)](https://goreportcard.com/report/github.com/Code-Hex/echo-static)  
[![GoDoc](https://godoc.org/github.com/Code-Hex/echo-static?status.svg)](https://godoc.org/github.com/Code-Hex/echo-static)  
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)  
# SYNOPSIS
```go
package main

import (
    static "github.com/Code-Hex/echo-static"

    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.Use(static.ServeRoot("/static", NewAssets("assets")))
    e.GET("/ping", func(c echo.Context) error {
        return c.String(200, "test")
    })
    // Listen and Server in 0.0.0.0:8080
    e.Start(":8080")
}

func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}
```
and put your asset file in the `assets` directory and execute the following code before compile it.

    go-bindata -o bindata.go assets/...
	
# DESCRIPTION

echo-static is File server middleware for [go-bindata](https://github.com/jteeuwen/go-bindata) and [echo](https://github.com/labstack/echo).

# INSTALLATION

    go get github.com/Code-Hex/echo-static

# AUTHOR

[codehex](https://twitter.com/CodeHex)