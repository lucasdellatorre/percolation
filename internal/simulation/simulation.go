package simulation

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"
	"github.com/lucasdellatorre/percolation/internal/unionfind"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Cell struct {
	Rect  pixel.Rect
	Color color.Color
}

type Subtitle struct {
	Text       string
	TextColor  color.Color
	TextVector pixel.Vec
	RectColor  color.Color
	Rect       pixel.Rect
}

func Run() {
	var width float64 = 800
	var height float64 = 800
	cfg := pixelgl.WindowConfig{
		Title:  "Percolation",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.White)

	n := 4
	u := unionfind.NewUnionFind(n)

	randomNumbers := generateUniqueRandomNumbers(n * n)

	animationCells := make([]Cell, n*n)

	CELL_SIZE := width / float64(n+n)

	centerX := width / 2
	centerY := height / 2

	matrixWidth := float64(n) * CELL_SIZE
	matrixHeight := float64(n) * CELL_SIZE

	startX := centerX - matrixWidth/2
	startY := centerY - matrixHeight/2

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	nText := text.New(pixel.V(startX-125, startY+centerY-50), basicAtlas)
	nText.Color = colornames.Black
	fmt.Fprintf(nText, "N = %d", n)

	//////////////////////////////

	subtitles := []Subtitle{
		{
			Text:       "open site",
			TextColor:  color.Black,
			TextVector: pixel.V(startX+10+40, 130),
			RectColor:  color.White,
			Rect:       pixel.R(startX, 160, startY+40, 120),
		},
		{
			Text:       "blocked site",
			TextColor:  color.Black,
			TextVector: pixel.V(startX+10+40, 80),
			RectColor:  color.Black,
			Rect:       pixel.R(startX, 120-10, startY+40, 80-10),
		},
	}

	for !win.Closed() {
		win.Update()
		win.Clear(colornames.White)
		nText.Draw(win, pixel.IM.Scaled(nText.Orig, 1.5))
		for _, subtitle := range subtitles {
			imd := imdraw.New(nil)
			imd.Color = subtitle.RectColor
			imd.Push(subtitle.Rect.Min, subtitle.Rect.Max)
			imd.Rectangle(0)
			imd.Draw(win)

			// BorderLine
			imd = imdraw.New(nil)
			imd.Color = color.Black
			imd.Push(subtitle.Rect.Min, subtitle.Rect.Max)
			imd.Rectangle(1)
			imd.Draw(win)

			subtitleText := text.New(subtitle.TextVector, basicAtlas)

			subtitleText.Color = subtitle.TextColor

			fmt.Fprintln(subtitleText, subtitle.Text)
			subtitleText.Draw(win, pixel.IM.Scaled(subtitleText.Orig, 1.5))
		}

		// Update open cells and draw
		for i := 0; !u.Percolates(); i++ {
			win.SetTitle(fmt.Sprintf("%s | Open sites: %d", cfg.Title, i))
			u.Open(randomNumbers[i])

			var minX, maxX = startX - CELL_SIZE, startY
			var minY, maxY = startY + matrixHeight, startY + matrixHeight - CELL_SIZE
			for j := range u.BlockedGrid {
				if u.BlockedGrid[j] {
					animationCells[j].Color = colornames.Black
				} else {
					animationCells[j].Color = colornames.White
				}

				if j > 0 && j%n == 0 {
					minX = startX
					maxX = startX + CELL_SIZE
					minY = minY - CELL_SIZE
					maxY = maxY - CELL_SIZE
				} else {
					minX = minX + CELL_SIZE
					maxX = maxX + CELL_SIZE
				}
				animationCells[j].Rect = pixel.R(minX, minY, maxX, maxY)

				// Draw cell
				imd := imdraw.New(nil)
				imd.Color = animationCells[j].Color
				imd.Push(animationCells[j].Rect.Min, animationCells[j].Rect.Max)
				imd.Rectangle(0)
				imd.Draw(win)

				// Draw borderline
				imd = imdraw.New(nil)
				imd.Color = colornames.Black
				imd.Push(animationCells[j].Rect.Min, animationCells[j].Rect.Max)
				imd.Rectangle(1)
				imd.Draw(win)
			}

			// Update window
			win.Update()
		}
		fmt.Println("Percolates", u.Percolates())
		time.Sleep(time.Second * 5)
	}

	fmt.Println(u.Percolates())
}

func generateUniqueRandomNumbers(n int) []int {
	randomNumbers := make([]int, n)
	usedNumbers := make(map[int]bool)

	for i := 0; i < n; i++ {
		var num int
		for {
			num = rand.Intn(n)
			if !usedNumbers[num] {
				break
			}
		}
		randomNumbers[i] = num
		usedNumbers[num] = true
	}

	return randomNumbers
}
