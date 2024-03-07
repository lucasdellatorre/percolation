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

var (
	screenWidth  float64
	screenHeight float64
	CELL_SIZE    float64
	centerX      float64
	centerY      float64
	matrixWidth  float64
	matrixHeight float64
	startX       float64
	startY       float64
)

func Run(n int) {
	win, err := initWindow()
	if err != nil {
		panic(err)
	}

	u := unionfind.NewUnionFind(n)
	randomNumbers := util.GenerateUniqueRandomNumbers(n * n)

	win.Clear(colornames.White)

	setSimulationMatrixSpecs(n)

	drawSubtitles(win, n)

	drawFullBlockedMatrix(win, n)

	for !win.Closed() {
		win.Update()

		for i := 0; !u.Percolates(); i++ {
			win.SetTitle(fmt.Sprintf("Percolation | Open sites %d", i))
			pos := randomNumbers[i]
			u.Open(pos)
			cellStyle := Cell{Border: true, Color: colornames.White}
			drawCell(win, pos, cellStyle, n)
			win.Update()
		}

		// Draw full open site
		for i := range u.BlockedGrid {
			if !u.BlockedGrid[i] && u.IsConnectedToTop(i) {
				cellStyle := Cell{Border: true, Color: colornames.Skyblue}
				drawCell(win, i, cellStyle, n)
				win.Update()
			}
		}
	}
}

func initWindow() (*pixelgl.Window, error) {
	screenWidth = 800
	screenHeight = 800
	cfg := pixelgl.WindowConfig{
		Title:  "Percolation",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	return win, err
}

func setSimulationMatrixSpecs(n int) {
	CELL_SIZE = screenWidth / float64(n+n)

	centerX = screenWidth / 2
	centerY = screenHeight / 2

	matrixWidth = float64(n) * CELL_SIZE
	matrixHeight = float64(n) * CELL_SIZE

	startX = centerX - matrixWidth/2
	startY = centerY - matrixHeight/2
}

func drawSubtitles(win *pixelgl.Window, n int) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	nText := text.New(pixel.V(startX-125, startY+centerY-50), basicAtlas)
	nText.Color = colornames.Black
	fmt.Fprintf(nText, "N = %d", n)
	nText.Draw(win, pixel.IM.Scaled(nText.Orig, 1.5))

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
}

func drawFullBlockedMatrix(win *pixelgl.Window, n int) {
	baseMatrix := imdraw.New(nil)

	baseMatrix.Color = colornames.Black
	baseMatrix.Push(pixel.V(startX, startY), pixel.V(startX+CELL_SIZE*float64(n), startY+CELL_SIZE*float64(n)))
	baseMatrix.Rectangle(0)
	baseMatrix.Draw(win)
}

func convert(pos int, n int) (int, int) {
	return pos / n, pos % n
}

func drawCell(win *pixelgl.Window, cellPos int, cell Cell, n int) {
	imd := imdraw.New(nil)
	row, col := convert(cellPos, n)
	cell.Rect = pixel.R(startX+float64(col)*CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE-CELL_SIZE, startX+float64(col)*CELL_SIZE+CELL_SIZE, startY+matrixHeight-float64(row)*CELL_SIZE)
	imd.Color = cell.Color
	imd.Push(cell.Rect.Min, cell.Rect.Max)
	imd.Rectangle(0)

	if cell.Border {
		imd.Color = colornames.Black
		imd.Push(cell.Rect.Min, cell.Rect.Max)
		imd.Rectangle(1)
	}
	imd.Draw(win)
}
