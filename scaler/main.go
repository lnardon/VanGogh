package scaler

import (
	"image"

	"golang.org/x/image/draw"
)

func ScaleImage(src image.Image, scalingTechnique string, scalingFactor int) image.Image {
	blankImage := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/scalingFactor, src.Bounds().Max.Y/scalingFactor))

	switch scalingTechnique {
	case "ApproxBiLinear":
		draw.ApproxBiLinear.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
		return blankImage
	case "BiLinear":
		draw.BiLinear.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
		return blankImage
	case "CatmullRom":
		draw.CatmullRom.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
		return blankImage
	default:
		draw.NearestNeighbor.Scale(blankImage, blankImage.Rect, src, src.Bounds(), draw.Over, nil)
		return blankImage
	}
}
