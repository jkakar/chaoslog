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
