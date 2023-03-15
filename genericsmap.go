package main

import "fmt"

func keys[k comparable, v any](data map[k]v) []k {

	keys := make([]k, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}
	return keys
}
func main() {
	intmap := map[int]int{ //using int type
		1: 11,
		2: 22,
		3: 33,
	}

	fmt.Println(keys(intmap))
	stringmap := map[string]int{
		"sangeetha": 1,
		"mathi":     2,
		"gomathi":   3,
		"ganesh":    4,
	}
	fmt.Println(keys(stringmap))
}
