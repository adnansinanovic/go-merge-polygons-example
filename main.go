package main

import (
	"fmt"
	"image/color"

	"github.com/akavel/polyclip-go"
	// only for drawing
	"github.com/fogleman/gg"
)

func newDrawContext() *gg.Context {
	dc := gg.NewContext(400, 400)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	return dc
}

type example struct {
	Title  string
	P1     polyclip.Polygon
	P2     polyclip.Polygon
	Result polyclip.Polygon
	dc1    *gg.Context
	dc2    *gg.Context
}

func newExample(title string, p1, p2 polyclip.Polygon) *example {
	ex := new(example)
	ex.Title = title
	ex.dc1 = newDrawContext()
	ex.dc2 = newDrawContext()
	ex.P1 = p1
	ex.P2 = p2
	// JOIN/MERGE polygons
	ex.Result = ex.P1.Construct(polyclip.UNION, ex.P2)

	return ex
}

func main() {
	colorRed := color.RGBA{R: 255, G: 50, B: 50}
	colorBlue := color.RGBA{R: 50, G: 50, B: 255}
	colorViolet := color.RGBA{R: 255, G: 50, B: 255}
	examples := []*example{
		newExample(
			"example 1",
			polyclip.Polygon{{{20, 40}, {140, 40}, {190, 200}, {150, 240}, {10, 90}}},
			polyclip.Polygon{{{100, 100}, {100, 300}, {300, 100}}}),
		newExample(
			"example 2",
			polyclip.Polygon{{{20, 40}, {140, 40}, {190, 200}, {150, 360}, {10, 90}}},
			polyclip.Polygon{{{100, 100}, {100, 300}, {300, 100}}}),
		newExample(
			"example 3",
			polyclip.Polygon{{{20, 40}, {140, 40}, {190, 200}, {150, 240}, {10, 90}}},
			polyclip.Polygon{{{150, 280}, {330, 390}, {190, 380}}}),
		newExample(
			"example 4",
			polyclip.Polygon{{{20, 40}, {140, 40}, {190, 200}, {150, 240}, {10, 90}}},
			polyclip.Polygon{{{30, 50}, {110, 60}, {180, 180}, {120, 200}}}),
	}

	for _, v := range examples {
		drawpPolygon(v.dc1, v.P1, colorRed)
		drawpPolygon(v.dc1, v.P2, colorBlue)

		drawpPolygon(v.dc2, v.Result, colorViolet)
		v.dc1.SavePNG(fmt.Sprintf("%v-step-1.png", v.Title))
		v.dc2.SavePNG(fmt.Sprintf("%v-step-2.png", v.Title))

	}
}

func drawpPolygon(c *gg.Context, p polyclip.Polygon, color color.Color) {
	c.SetColor(color)
	width := float64(4)

	for _, v := range p {
		index := 1
		for index < len(v) {
			previousPoint := v[index-1]
			point := v[index]
			c.DrawLine(previousPoint.X, previousPoint.Y, point.X, point.Y)
			c.SetLineWidth(width)
			index++
		}

		// draw last line
		lastPoint := v[len(v)-1]
		firstPoint := v[0]
		c.DrawLine(lastPoint.X, lastPoint.Y, firstPoint.X, firstPoint.Y)
		c.SetLineWidth(width)
	}

	if len(p) > 1 {
		c.SetRGB(200, 200, 50)
		c.DrawString("no intersected polygons ", 20, float64(c.Height())-20)
	}
	c.Stroke()

}
