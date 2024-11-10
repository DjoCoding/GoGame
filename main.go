package main

import (
	"log"
	"math/rand"
	"github.com/gdamore/tcell/v2"
)

const (
	CoinChar          = '#'
	PlayerChar        = '@'
	InitialCoinsCount = 1
)

type Coins struct {
	items Sprites
}

func (coins *Coins) RemoveCoin(player Sprite) {
	for idx, coin := range coins.items {
		if coin.X == player.X && coin.Y == player.Y {
			coins.items.Remove(idx)
			break
		}
	}
}

func (coins *Coins) SetCoins(screen tcell.Screen, coinsCount int) {
	width, height := screen.Size()
	coins.items.InitRandom(CoinChar, coinsCount, width, height, 0, 0)
}

func (coins *Coins) Render(screen tcell.Screen) {
	coins.items.Render(screen)
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("ERROR: failed to create the game screen, ", err)
	}

	err = screen.Init()
	if err != nil {
		log.Fatal("ERROR: failed to initialize the game screen, ", err)
	}

	defer screen.Fini()

	screenStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlue)
	screen.SetStyle(screenStyle)

	player := NewSprite(rand.Intn(20), rand.Intn(20), '@')

	coinsCount := InitialCoinsCount
	coins := Coins{
		make(Sprites, coinsCount),
	}
	coins.SetCoins(screen, coinsCount)

	for {
		ev := screen.PollEvent()
		playerHasMoved := false

		if ev != nil {
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					return
				}

				switch ev.Rune() {
				case 'z':
					player.Move(UP)
					playerHasMoved = true
				case 's':
					player.Move(DOWN)
					playerHasMoved = true
				case 'd':
					player.Move(RIGHT)
					playerHasMoved = true
				case 'q':
					player.Move(LEFT)
					playerHasMoved = true
				}
			}
		}

		if playerHasMoved {
			coins.RemoveCoin(player)
			if len(coins.items) == 0 {
				coinsCount++
				coins.SetCoins(screen, coinsCount)
			}
		}

		screen.Clear()
		player.Render(screen)
		coins.Render(screen)
		screen.Show()
	}
}

func DrawString(screen tcell.Screen, x, y int, s string) {
	for i, c := range s {
		screen.SetContent(x+i, y, c, nil, tcell.StyleDefault)
	}
}

