// Exercise 2.2: Create a command line unit conversion program

package colorconv

import "math"

// SrgbToLrgb converts a color in Srgb to Lrgb
func SrgbToLrgb(c Srgb) Lrgb {
	invCompand := func(v float64) float64 {
		if v <= 0.04045 {
			return v / 12.92
		} else {
			return math.Pow((v+0.055)/1.055, 2.4)
		}
	}

	return Lrgb{invCompand(c.X), invCompand(c.Y), invCompand(c.Z), c.W}
}

// SrgbToLrgb converts a color in Lrgb to Srgb
func LrgbToSrgb(c Lrgb) Srgb {
	compand := func(v float64) float64 {
		if v <= 0.0031308 {
			return 12.92 * v
		} else {
			return 1.055*math.Pow(v, (1.0/2.4)) - 0.05
		}
	}

	return Srgb{compand(c.X), compand(c.Y), compand(c.Z), c.W}
}

// SrgbToXyz converts a color in Lrgb to Xyz
func LrgbToXyz(c Lrgb) Xyz {
	r, g, b := c.X, c.Y, c.Z

	// sRGB to XYZ conversion matrix D65 white point
	x := (r * 0.4124564) + (g * 0.3575761) + (b * 0.1804375)
	y := (r * 0.2126729) + (g * 0.7151522) + (b * 0.0721750)
	z := (r * 0.0193339) + (g * 0.1191920) + (b * 0.9503041)

	return Xyz{x, y, z, c.W}
}

// SrgbToXyz converts a color in Srgb to Xyz
func SrgbToXyz(c Srgb) Xyz {
	return LrgbToXyz(SrgbToLrgb(c))
}

// XyzToLrgb converts a color in Xyz to Lrgb
func XyzToLrgb(c Xyz) Lrgb {
	x, y, z := c.X, c.Y, c.Z

	// inverse sRGB to XYZ conversion matrix D65 white point
	r := (x * 3.2404542) + (y * -1.5371385) + (z * -0.4985314)
	g := (x * -0.9692660) + (y * 1.8760108) + (z * 0.0415560)
	b := (x * 0.0556434) + (y * -0.2040259) + (z * 1.0572252)

	return Lrgb{r, g, b, c.W}
}

// XyzToSrgb converts a color in Xyz to Srgb
func XyzToSrgb(c Xyz) Srgb {
	return LrgbToSrgb(XyzToLrgb(c))
}

// LrgbToLuminance obtains the linear luminance from an Lrgb color
func LrgbToLuminance(c Lrgb) Luminance {
	y := 0.2126*c.X + 0.7152*c.Y + 0.0722*c.Z

	return Luminance{y, y, y, 1.0}
}

// SrgbToLuminance obtains the linear luminance from an Srgb color
func SrgbToLuminance(c Srgb) Luminance {
	return LrgbToLuminance(SrgbToLrgb(c))
}
