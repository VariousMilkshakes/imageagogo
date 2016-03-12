package process

import "image/color"

// Darken the current pixel
func Darken(pixel color.Color, modifier uint8) color.Color {
	r, g, b, a := pixel.RGBA()
	newr := uint8((r >> 8)) / modifier
	newg := uint8((g >> 8)) / modifier
	newb := uint8((b >> 8)) / modifier
	newa := uint8((a >> 8)) / modifier
	pixel = color.RGBA{newr, newg, newb, newa}
	return pixel
}

// GreyScale averages the pixel colors into a grey
func GreyScale(pixel color.Color) color.Color {
	r, g, b, a := pixel.RGBA()
	avg := uint8((r>>8 + g>>8 + b>>8 + a>>8) / 4)
	pixel = color.RGBA{avg, avg, avg, 255}
	return pixel
}

// Threshold seperates bright and dark areas of an image
func Threshold(pixel color.Color, min, max uint8) (color.Color, bool) {
	black := color.Black
	white := color.White
	bit := true

	pixel = GreyScale(pixel)
	r, _, _, _ := pixel.RGBA()
	sample := uint8(r >> 8)

	if sample > min && sample < max {
		pixel = white
	} else {
		pixel = black
		bit = false
	}

	return pixel, bit
}

// BitwiseMask removes color from image using a Threshold
func BitwiseMask(pixel color.Color, min, max uint8) color.Color {
	_, bit := Threshold(pixel, min, max)

	if !bit {
		pixel = color.White
	}

	return pixel
}
