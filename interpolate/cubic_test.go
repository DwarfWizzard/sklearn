package interpolate

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func ExampleCubicSpline() {
	xs := make([]float64, 20)
	ys := make([]float64, len(xs))
	for i := range xs {
		xs[i] = float64(i) - 9.5
		ys[i] = rand.NormFloat64()
	}
	f := CubicSpline(xs, ys)
	x2 := make([]float64, 220)
	y2 := make([]float64, 220)
	for i := range x2 {
		x2[i] = float64(i)/10 - 10.5
		y2[i] = f(x2[i])
	}

	p, err := plot.New()
	check(err)
	scatter, err := plotter.NewScatter(&xy{xs, ys})
	line, err := plotter.NewLine(&xy{x2, y2})
	p.Add(scatter, line)
	// if err := p.Save(16*vg.Inch, 9*vg.Inch, "/tmp/spline.png"); err != nil {
	// 	panic(err)
	// }
	// _, err = os.Stat("/usr/bin/display")
	// if err == nil {
	// 	cmd := exec.Command("/usr/bin/display", "/tmp/spline.png")
	// 	cmd.Start()
	// }

	// Output:
}