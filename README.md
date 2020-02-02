# UUID generator for Go

This package generates version 4 UUIDs ([RFC 4122][rfc]). It's 35 times
faster than other UUID libraries ([1][gofrs], [2][google]), and, unlike
those libraries, cannot fail to generate a UUID.

    BenchmarkSelf-8         100000000               13.4 ns/op
    BenchmarkGofrs-8         2000000               623 ns/op
    BenchmarkGoogle-8        2000000               627 ns/op

It's faster and more reliable because it generates UUIDs from an
internal CSPRNG. Operating system entropy is only used to seed the
generator once. Other libraries read 16 bytes of operating system
entropy for each UUID, which is why they're slower and can fail.

API documentation: <https://godoc.org/nullprogram.com/x/uuid>

## Example usage

```go
package main

import (
	"flag"
	"fmt"

	"nullprogram.com/x/uuid"
)

func main() {
	n := flag.Int("n", 1, "number of UUIDs to generate")
	flag.Parse()

	g := uuid.NewGen()
	for i := 0; i < *n; i++ {
		fmt.Println(g.NewV4())
	}
}
```

Result:

    $ ./example -n 4
    66ccddf3-be66-455d-8bab-13a20cf57d05
    17e7e09c-fafd-4482-999c-7fbd87a1e7a8
    b7146884-9678-435e-b76e-9fc8407544e7
    cc6f0bf7-9762-456c-890a-3213c6965e95

[gofrs]: https://github.com/gofrs/uuid
[google]: https://github.com/google/uuid
[rfc]: https://tools.ietf.org/html/rfc4122
