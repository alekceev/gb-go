package main

import (
	"fmt"
)

/*
1. Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.
*/

type Attr map[string]float64

// StructA
type StructA struct {
	Name  string
	Attrs Attr
}

func (s *StructA) SetAttr(attr string, value float64) {
	if s.Attrs == nil {
		s.Attrs = Attr{}
	}
	s.Attrs[attr] = value
}

func (s *StructA) GetAttr(attr string) float64 {
	return s.Attrs[attr]
}

// StructB
type StructB struct {
	Name  string
	Attrs Attr
}

func (s *StructB) SetAttr(attr string, value float64) {
	if s.Attrs == nil {
		s.Attrs = Attr{}
	}
	s.Attrs[attr] = value
}

func (s *StructB) GetAttr(attr string) float64 {
	return s.Attrs[attr]
}

// Shape interface
type Shape interface {
	GetAttr(attr string) float64
}

func totalAttrVal(attr string, shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		if shape == nil {
			continue
		}
		sum += shape.GetAttr(attr)
	}
	return sum
}

func main() {
	itemA := &StructA{Name: "itemA", Attrs: Attr{"length": 10}}
	itemB := &StructA{Name: "itemB"}
	itemB.SetAttr("length", 5)
	itemB.SetAttr("weigth", 100)

	fmt.Println("Total length: ", totalAttrVal("length", itemA, itemB))
	fmt.Println("Total weigth: ", totalAttrVal("weigth", itemA, itemB))
}
