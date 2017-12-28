package views

import (
	"fmt"
	"g0/models"
)

// Screen object
type Screen struct {
	device string
}

// NewScreen : screen onbject constructor
func NewScreen() *Screen {
	return &Screen{device: "Screen"}
}

// RenderList : render list of tasks
func (s *Screen) RenderList(list []*models.Task) {
	lineBreak()
	for _, t := range list {
		fmt.Printf("\t%03d : %s\n", t.ID(), t.Task())
	}
	lineBreak()
}

// RenderHelp : render documentation
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
