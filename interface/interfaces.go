// interface 接口

package main

import (
	"fmt"
	"log"
	"math"
)

type geometry interface {
	// method
	area() float64
	perim() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// method
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// function
func measure(g geometry) {
	log.Println(g)
	log.Println(g.area())
	log.Println(g.perim())
}

func measure2(e interface{}) {
	// fmt.Println(e.(type))
	switch v := e.(type) {
	case rect:
		fmt.Println("zhengfangxxing", v)
	}
}

// main
func main() {
	r := rect{3, 4}
	c := circle{5}
	measure(r)
	measure(c)

	measure2(c)
}
