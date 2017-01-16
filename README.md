Structured logging package.

## Example

```go
package main

import log "github.com/jkakar/chaoslog"

func main() {
	data := map[string]interface{}{
		"string": "hello",
		"int":    1,
		"float":  3.14,
		"bool":   true,
	}
	log.Log(data)
}
```

Run it:

```bash
$ for ((i = 1; i <= 10; i++)); do go run cmd/chaoslog-example/main.go; done
2017/01/15 20:39:52 string=hello int=1 float=3.14 bool=true
2017/01/15 20:39:52 string=hello int=1 float=3.14 bool=true
2017/01/15 20:39:52 string=1 int=3.14 float=true bool=hello
2017/01/15 20:39:52 string=hello int=1 float=3.14 bool=true
2017/01/15 20:39:53 float=1 bool=3.14 string=true int=hello
2017/01/15 20:39:53 string=hello int=1 float=3.14 bool=true
2017/01/15 20:39:53 int=true float=hello bool=1 string=3.14
2017/01/15 20:39:53 string=3.14 int=true float=hello bool=1
2017/01/15 20:39:53 string=hello int=1 float=3.14 bool=true
2017/01/15 20:39:54 string=hello int=1 float=3.14 bool=true
```
