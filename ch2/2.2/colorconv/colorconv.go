// Exercise 2.2: Create a command line unit conversion program

package colorconv

import "fmt"

type Vec4 struct {
	X, Y, Z, W float64
}

type Luminance Vec4
type Srgb Vec4
type Lrgb Vec4
type Xyz Vec4

func White() Srgb { return Srgb{1.0, 1.0, 1.0, 1.0} }
func Black() Srgb { return Srgb{0.0, 0.0, 0.0, 1.0} }

// There so many color spaces. I did more in Python, but I'm not doing anymore
// here.

func (v Vec4) String() string {
	return fmt.Sprintf("(%g, %g, %g, %g)", v.X, v.Y, v.Z, v.W)
}

func (y Luminance) String() string {
	return fmt.Sprintf("luminance: %s", Vec4(y))
}
func (c Srgb) String() string { return fmt.Sprintf("srgb: %s", Vec4(c)) }
func (c Lrgb) String() string { return fmt.Sprintf("lrgb: %s", Vec4(c)) }
func (c Xyz) String() string  { return fmt.Sprintf("xyx: %s", Vec4(c)) }
