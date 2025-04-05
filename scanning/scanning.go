package scanning

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
)

func DecodeImage(filename string) (image.Image, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}

	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

func GetImageMatrix(pic image.Image) (grid [][]color.Color) {
	size := pic.Bounds()
	for i := range size.Max.X {
		var y []color.Color
		for j := range size.Max.Y {
			y = append(y, pic.At(i, j))
		}
		grid = append(grid, y)
	}
	return grid
}
