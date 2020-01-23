# fakeuuid

Mockable uuid.New().String()

## Install

```
go get github.com/johejo/fakeuuid
```

## Example

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
