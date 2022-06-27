# Go/Golang package to list all timezones defined by OS (Linux or macOS)

## Status

Tested on macOS and RHEL7, should work on other Linux distos as well.

Windows is currently not supported.

## Installation <a id="installation"></a>
```
go get github.com/mileusna/timezones
```

## Example<a id="example"></a>

```go
package main

import (
    "fmt"
    "strings"

    "github.com/mileusna/timezones"
)

func main() {

    // checks common zone info directories on your OS
    tzones := timezones.List()
    for _, tz := range tzones {
        fmt.Println(tz)
    }

    // if yor linux disto use some uncommon location for zoneinfo
    tzones = timezones.ListDir("/usr/share/zoneinfo/")
}
```