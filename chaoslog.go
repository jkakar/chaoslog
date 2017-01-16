package chaoslog

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

// Log writes key/value values into the log.
func Log(data map[string]interface{}) {
	values := []interface{}{}
	for _, v := range data {
		values = append(values, v)
	}

	i := 0
	indexes := rand.Perm(len(values))
	pairs := []string{}
	for k, _ := range data {
		pairs = append(pairs, fmt.Sprintf("%s=%v", k, values[indexes[i]]))
		i += 1
	}
	log.Print(strings.Join(pairs, " "))
}
