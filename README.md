echo-static
====

# SYNOPSIS
```go
package main

import (
    static "github.com/Code-Hex/echo-static"

    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    r.Use(static.ServeRoot("/static", NewAssets("assets")))
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

# DESCRIPTION

echo-static is File server middleware for [go-bindata](https://github.com/jteeuwen/go-bindata) and [echo](https://github.com/labstack/echo).

# INSTALLATION

    go get github.com/Code-Hex/echo-static

# AUTHOR

[Code-Hex](https://github.com/Code-Hex)  
<x00.x7f@gmail.com>
