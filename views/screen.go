package views

import (
	"fmt"
	"g0/models"
)

// Screen ...
type Screen struct {
	device string
}

// NewScreen ...
func NewScreen() *Screen {
	return &Screen{device: "Screen"}
}

// Render ...
func (s *Screen) Render(t *models.Task) {
	lineBreak()
	fmt.Println("\t", t.ID(), ":", t.Task())
	lineBreak()
}

// RenderList ...
func (s *Screen) RenderList(list []*models.Task) {
	lineBreak()
	for _, t := range list {
		//fmt.Println(t)
		fmt.Println("\t", t.ID(), ":", t.Task())
	}
	lineBreak()
}

func lineBreak() {
	fmt.Println("")
}
