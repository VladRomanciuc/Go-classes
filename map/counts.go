package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

//func that return counts of the words in the map
func GetCounts(words []string) map[string]int {
	m := map[string]int{}
	
	for _, id := range words {
		if _, ok := m[id]; !ok {
			m[id] = 0
		}
		m[id]++
	}
	return m
}


func main() {
	words := []string{}

	//map generator
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		words = append(words, key[:2])
	}

	//call the func
	counts := GetCounts(words)

	//testing
	fmt.Println(counts["00"])
	fmt.Println(counts["ff"])
	fmt.Println(counts["dd"])
}
