package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["bell labs"] = Vertex{
		40.6564, -74.2823,
	}

	fmt.Println(m["bell labs"])
	
}