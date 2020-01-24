# fakeuuid

[![go-test](https://github.com/johejo/fakeuuid/workflows/go-test/badge.svg)](https://github.com/johejo/fakeuuid/actions?query=workflow%3Ago-test)
[![GitHub license](https://img.shields.io/github/license/johejo/fakeuuid)](https://github.com/johejo/fakeuuid/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/johejo/fakeuuid?status.svg)](https://godoc.org/github.com/johejo/fakeuuid)

Mockable uuid.New().String()

## Install

```
go get github.com/johejo/fakeuuid
```

## Example

https://play.golang.org/p/P6ysWK_RBqe

```go
package main

import (
    "fmt"

    uuid "github.com/johejo/fakeuuid"
)

func main() {
    fmt.Println(uuid.New().String()) // print uuid

    restore := uuid.Mock("hello")
    fmt.Println(uuid.New().String()) // print "hello"

    restore()
    fmt.Println(uuid.New().String()) // print uuid
}
```

## Licence

MIT

## Author

Mitsuo Heijo (johejo)
