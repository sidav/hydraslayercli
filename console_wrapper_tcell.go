package main

import (
	"github.com/gdamore/tcell/v2"
)

type cwtcell struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
	bg_color, fg_color            tcell.Color
	currentLine                   int
	currentChar                   int
}

func (c *cwtcell) init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	var e error
	c.screen, e = tcell.NewScreen()
	if e != nil {
		panic(e)
	}
	if e = c.screen.Init(); e != nil {
		panic(e)
	}
	c.screen.EnableMouse()
	c.style = tcell.StyleDefault.Foreground(c.fg_color).Background(c.bg_color)
	c.screen.SetStyle(c.style)
	c.CONSOLE_WIDTH, c.CONSOLE_HEIGHT = c.screen.Size()
}

func (c *cwtcell) print(s string) {
	for i := 0; i < len(s); i++ {
		c.screen.SetCell(c.currentChar, c.currentLine, c.style, rune(s[i]))
		c.currentChar++
		if c.currentChar == c.CONSOLE_WIDTH {
			c.currentChar = 0
			c.CONSOLE_HEIGHT++
		}
	}
}

func (c *cwtcell) println(s string) {
	c.print(s)
	c.currentLine++
	c.currentChar = 0
}

func (c *cwtcell) closeConsole() {
	c.screen.Fini()
}

func (c *cwtcell) clear() { // is suddenly less buggy than screen.Clear()
	for x := 0; x < c.CONSOLE_WIDTH; x++ {
		for y := 0; y < c.CONSOLE_HEIGHT; y++ {
			c.screen.SetCell(c.currentChar, c.currentLine, c.style, ' ')
		}
	}
	c.currentChar = 0
	c.currentLine = 0
}

func (c *cwtcell) flush() {
	c.screen.Show()
}

func (c *cwtcell) read() string {
	currLine := ""
	key := ""
	for {
		ev := c.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			key = eventToKeyString(ev)
		case *tcell.EventResize:
			c.screen.Sync()
			c.CONSOLE_WIDTH, c.CONSOLE_HEIGHT = c.screen.Size()
		}
		if key == "UP" {

		}
		if key == "ENTER" {
			return currLine
		}
		if len(key) == 1 {
			currLine += key
		}
		c.putString(currLine, c.currentChar, c.currentLine)
		c.flush()
	}
}

func eventToKeyString(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyUp:
		return "UP"
	case tcell.KeyRight:
		return "RIGHT"
	case tcell.KeyDown:
		return "DOWN"
	case tcell.KeyLeft:
		return "LEFT"
	case tcell.KeyEscape:
		return "ESCAPE"
	case tcell.KeyEnter:
		return "ENTER"
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		return "BACKSPACE"
	case tcell.KeyTab:
		return "TAB"
	case tcell.KeyDelete:
		return "DELETE"
	case tcell.KeyInsert:
		return "INSERT"
	case tcell.KeyEnd:
		return "END"
	case tcell.KeyHome:
		return "HOME"
	default:
		return string(ev.Rune())
	}
}

func (c *cwtcell) putString(s string, x, y int) {
	length := len([]rune(s))
	for i := 0; i < length; i++ {
		c.screen.SetCell(x+i, y, c.style, rune(s[i]))
	}
}
