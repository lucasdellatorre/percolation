package guisimulation

import (
	"fmt"
	"image/color"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/gopxl/pixel/text"
	"github.com/lucasdellatorre/percolation/internal/unionfind"
	"github.com/lucasdellatorre/percolation/internal/util"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Cell struct {
	Rect   pixel.Rect
	Color  color.Color
	Border bool
}

type Subtitle struct {
	Text       string
	TextColor  color.Color
	TextVector pixel.Vec
	RectColor  color.Color
	Rect       pixel.Rect
}

func Run(n int) {
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

	u := unionfind.NewUnionFind(n)

	randomNumbers := util.GenerateUniqueRandomNumbers(n * n)

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

	subtitles := []Subtitle{
		{
			Text:       "full open site",
			TextColor:  colornames.Black,
			TextVector: pixel.V(startX+10+40, 130),
			RectColor:  colornames.Skyblue,
			Rect:       pixel.R(startX, 160, startY+40, 120),
		},
		{
			Text:       "open site",
			TextColor:  colornames.Black,
			TextVector: pixel.V(startX+10+40, 80),
			RectColor:  colornames.White,
			Rect:       pixel.R(startX, 120-10, startY+40, 80-10),
		},
		{
			Text:       "blocked site",
			TextColor:  colornames.Black,
			TextVector: pixel.V(startX+10+40, 30),
			RectColor:  colornames.Black,
			Rect:       pixel.R(startX, 70-10, startY+40, 30-10),
		},
	}

	fmt.Println(subtitles)

	win.Clear(colornames.White)

	// var minX, maxX = startX - CELL_SIZE, startY
	// var minY, maxY = startY + matrixHeight, startY + matrixHeight - CELL_SIZE
	win.Clear(colornames.White)

	baseMatrix := imdraw.New(nil)

	baseMatrix.Color = colornames.Black
	baseMatrix.Push(pixel.V(startX, startY), pixel.V(startX+CELL_SIZE*float64(n), startY+CELL_SIZE*float64(n)))
	baseMatrix.Rectangle(0)
	baseMatrix.Draw(win)

	for !win.Closed() {
		win.Update()
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

		for i := 0; !u.Percolates(); i++ {
			win.SetTitle(fmt.Sprintf("Percolation | Open sites %d", i))
			pos := randomNumbers[i]
			u.Open(pos)
			imd := imdraw.New(nil)
			cell := Cell{Border: true, Color: colornames.White}
			row, col := convert(pos, n)
			cell.Rect = pixel.R(startX+float64(col)*CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE-CELL_SIZE, startX+float64(col)*CELL_SIZE+CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE)
			drawCell(imd, cell)
			imd.Draw(win)
			win.Update()
		}

		for i := range u.BlockedGrid {
			if !u.BlockedGrid[i] && u.IsConnectedToTop(i) {
				imd := imdraw.New(nil)
				cell := Cell{Border: true, Color: colornames.Skyblue}
				row, col := convert(i, n)
				cell.Rect = pixel.R(startX+float64(col)*CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE-CELL_SIZE, startX+float64(col)*CELL_SIZE+CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE)
				drawCell(imd, cell)
				imd.Draw(win)
				win.Update()
			}
		}
		win.Update() //do not remove
	}
}

func convert(pos int, n int) (int, int) {
	return pos / n, pos % n
}

func drawCell(imd *imdraw.IMDraw, cell Cell) {
	imd.Color = cell.Color
	imd.Push(cell.Rect.Min, cell.Rect.Max)
	imd.Rectangle(0)

	if cell.Border {
		imd.Color = colornames.Black
		imd.Push(cell.Rect.Min, cell.Rect.Max)
		imd.Rectangle(1)
	}
}
