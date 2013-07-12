GoREST
===========

[![Build Status](https://travis-ci.org/HourlyHost/gorest.png)](https://travis-ci.org/HourlyHost/gorest)

GoREST is a RESTful style web-services framework. [Originally](https://code.google.com/p/gorest/) created by Siyabonga Dlamini.

Creating services in Go is straight forward, GoREST takes this a step further by adding a layer that makes tedious tasks much more automated and avoids regular pitfalls. This gives you the opportunity to focus more on the task at hand: minor low-level HTTP handling. 

* Type safe conversion of URL parameters into method arguments.
* Automated marshal/unmarshall of HTTP entities.
* Highly configurable endpoint system that makes grouping of endpoints into services easy.
* Tag (annotations) based configuration of end-points and services.
* Easily pluggable URL system to allow various multiplexing configurations.
* Ability to override the default RESTful behavior.

## Setup

**Step 1: Get**

```
go get github.com/HourlyHost/gorest
```

**Step 2: Import**

```Go
import (
  "github.com/HourlyHost/gorest"
)
```

## Example

```Go
func main() {
    gorest.RegisterService(new(HelloService))
    http.Handle("/", gorest.Handle())    
    http.ListenAndServe(":8787", nil)
}

type HelloService struct {
    gorest.RestService `root:"/tutorial/"`
    helloWorld  gorest.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
    sayHello    gorest.EndPoint `method:"GET" path:"/hello/{name:string}" output:"string"`
}

func (serv HelloService) HelloWorld() string {
    return "Hello World"
}

func (serv HelloService) SayHello(name string) string {
    return "Hello " + name
}
```

## Usage

[Getting Started](https://code.google.com/p/gorest/wiki/GettingStarted)

Documentation is available at [GoDoc](http://godoc.org/github.com/HourlyHost/gorest).
