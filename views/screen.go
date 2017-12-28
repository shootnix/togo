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
		fmt.Printf("\t%03d : %s\n", t.ID(), t.Task())
	}
	lineBreak()
}

func (s *Screen) RenderHelp() {
	lineBreak()
	fmt.Println("\t", "-add  -- Add task")
	fmt.Println("\t", "-d    -- Delete task")
	fmt.Println("\t", "-r    -- Mark task as resolved")
	fmt.Println("\t", "-up   -- Raise priority of the task")
	fmt.Println("\t", "-down -- Redure priority of the task")
	lineBreak()
}

func lineBreak() {
	fmt.Println("")
}
