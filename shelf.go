package main

func NewShelf() *Shelf {
	return &Shelf{
		products: make(map[string]float64),
	}
}

type Shelf struct {
	products map[string]float64
}
