package main

import (
	"github.com/gdamore/tcell/v2"
)

type cwtcell struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
	currentLine                   int
	currentChar                   int
	commandHistory                []string
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
	// c.screen.EnableMouse()
	c.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	c.screen.SetStyle(c.style)
	c.CONSOLE_WIDTH, c.CONSOLE_HEIGHT = c.screen.Size()
}

func (c *cwtcell) print(s string) {
	for i := 0; i < len(s); i++ {
		i += c.considerColorInStringAtPosition(s, i)
		if i >= len(s) {
			return
		}
		c.screen.SetCell(c.currentChar, c.currentLine, c.style, rune(s[i]))
		c.currentChar++
		if c.currentChar == c.CONSOLE_WIDTH {
			c.currentChar = 0
			c.currentLine++
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
	c.screen.Clear()
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
			for i:=len(c.commandHistory)-1; currLine == c.commandHistory[i] && i > 0; i-- {
				currLine = c.commandHistory[i]
			}
		}
		if key == "ENTER" {
			c.commandHistory = append(c.commandHistory, currLine)
			return currLine
		}
		if key == "BACKSPACE" {
			if len(currLine) > 0 {
				c.screen.SetCell(c.currentChar+len(currLine)-1, c.currentLine, c.style, ' ')
				currLine = currLine[:len(currLine)-1]
			}
		}
		if len(key) == 1 {
			currLine += key
		}
		c.putString(currLine, c.currentChar, c.currentLine)
		c.flush()
	}
}

// non-public

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

func (c *cwtcell) considerColorInStringAtPosition(s string, pos int) int {
	if s[pos] == "\033"[0] {
		if s[pos:pos+4] == Reset {
			c.style = c.style.Foreground(tcell.ColorWhite)
			return 4
		}
		switch s[pos : pos+5] {
		case White:
			c.style.Foreground(tcell.ColorWhite + tcell.ColorValid)
		case Red:
			c.style = c.style.Foreground(tcell.ColorRed)
		case Blue:
			c.style = c.style.Foreground(tcell.ColorDarkBlue)
		case Green:
			c.style = c.style.Foreground(tcell.ColorGreen)
		case Gray:
			c.style = c.style.Foreground(tcell.ColorGray)
		case Cyan:
			c.style = c.style.Foreground(tcell.ColorDarkCyan)
		default:
			panic("no color")
		}
		return 5
	}
	return 0
}