gin-statsd
========
[![GitHub release](https://img.shields.io/github/release/amalfra/gin-statsd.svg)](https://github.com/amalfra/gin-statsd/releases)
[![Build Status](https://travis-ci.org/amalfra/gin-statsd.svg?branch=master)](https://travis-ci.org/amalfra/gin-statsd)
[![GoDoc](https://godoc.org/github.com/amalfra/gin-statsd?status.svg)](https://godoc.org/github.com/amalfra/gin-statsd/middleware)
[![Go Report Card](https://goreportcard.com/badge/github.com/amalfra/gin-statsd)](https://goreportcard.com/report/github.com/amalfra/gin-statsd)

A [Gin](https://gin-gonic.github.io/gin/) middleware for reporting to statsd daemon.

## Installation
You can download the middleware using
```sh
go get github.com/amalfra/gin-statsd
```
## Usage
Next, import the package
``` go
import (
  statsdMiddleware "github.com/amalfra/gin-statsd/middleware"
)
```

1. Attach middleware to all routes
``` go
  r := gin.New()
  r.Use(statsdMiddleware.New(statsdMiddleware.Options{}))
```

or 

2. Attach middleware to specific route
``` go
  r := gin.New()
  r.GET("/", statsdMiddleware.New(statsdMiddleware.Options{}), func(c *gin.Context) {})
```

By default, the middleware will send status_code and response_time stats.
All stats are namespaced under `statsdKey` key by default. This can be configured using options.

### Configuring using Options
The middleware allows configuring using Options struct. The struct with config should be passed when initializing with New method. Supported configurations are:

* Host - statsd daemon's host; defaults to 127.0.0.1
* Port - statsd daemon's port; defaults to 8125
* RequestKey - the namespace for stats; defaults to statsdKey

Configuring namespace:
``` go
  r := gin.New()
  r.Use(statsdMiddleware.New(statsdMiddleware.Options{RequestKey: "myNamespace"}))
```
Configuring statsd host and port:
``` go
  r := gin.New()
  r.Use(statsdMiddleware.New(statsdMiddleware.Options{Host: "myhost.statsd", Port: 8089}))
```

## Development
Questions, problems or suggestions? Please post them on the [issue tracker](https://github.com/amalfra/gin-statsd/issues).

You can contribute changes by forking the project and submitting a pull request. Feel free to contribute :heart_eyes:

## UNDER MIT LICENSE

The MIT License (MIT)

Copyright (c) 2018 Amal Francis

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

