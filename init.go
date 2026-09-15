package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer s.Fini()

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()

	s.Put(0, 0, "H", defStyle)
	s.Put(1, 0, "i", defStyle)
	s.Put(2, 0, "!", defStyle)
	s.PutStrStyled(0, 1, "Hello with PutStrStyled", defStyle)
	s.Show()

	// Wait for a key press, otherwise the program exits
	// immediately and the terminal is restored before you see anything.
	for {
		ev := s.PollEvent()
		if _, ok := ev.(*tcell.EventKey); ok {
			return
		}
	}
}
