xdg
===

An implementation of the XDG Base Directory Specification.

For more information, see: [XDG Base Directory Specification](http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

Why
---

Built as a simple and correct XDG implementation that maintains convenience,
but doesn't unnecessarily hide complexity. To simplify searching, `DATA` and
`CONFIG` are provided as combined `*_HOME` and `*_DIRS` accessors.

Usage
-----

```go
import (
    "github.com/miquella/xdg"
)

var (
    XDG = xdg.WithSuffix("example")
)

// FindExampleConfigs returns order-preference example.config files
func FindExampleConfigs() ([]string, error) {
    return XDG.CONFIG.Find("example.config")
}
```
