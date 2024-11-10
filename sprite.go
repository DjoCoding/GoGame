package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

const (
	LEFT = iota
	RIGHT
	UP
	DOWN
)

type Sprite struct {
	X, Y int
	Char rune
}

type Sprites []Sprite

func NewSprite(x, y int, char rune) Sprite {
	return Sprite{
		X:    x,
		Y:    y,
		Char: char,
	}
}

func (s *Sprite) Render(screen tcell.Screen) {
	screen.SetContent(s.X, s.Y, s.Char, nil, tcell.StyleDefault)
}

func (s *Sprite) Move(direction int) {
	switch direction {
	case LEFT:
		s.X--
	case RIGHT:
		s.X++
	case UP:
		s.Y--
	case DOWN:
		s.Y++
	}
}

func (sprites *Sprites) Append(sprite Sprite) {
	*sprites = append(*sprites, sprite)
}

func (sprites *Sprites) Clear() {
	*sprites = nil
}

func (sprites *Sprites) Render(screen tcell.Screen) {
	for _, s := range *sprites {
		s.Render(screen)
	}
}

func (sprites *Sprites) InitRandom(spriteChar rune, length, width, height, offsetWidth, offsetHeight int) {
	sprites.Clear();
	for i := 0; i < length; i++ {
		sprites.Append(NewSprite(offsetWidth+rand.Intn(width-offsetWidth), offsetHeight+rand.Intn(height-offsetHeight), spriteChar))
	}
}

func (sprites *Sprites) Remove(idx int) {
	// the idx is always in range
	*sprites = append((*sprites)[:idx], (*sprites)[idx+1:]...)
}
