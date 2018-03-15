package tour

import (
	"fmt"
	"io"
	"image"
	"image/color"
)

type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type SqrtError float64

func (e SqrtError) Error() string  {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return "hah"
}

func LoopSqrtError(s float64, n int) (float64, error)  {
	if s < 0 {
		return 0, SqrtError(s)
	}
	z := 1.0
	for i := 0; i < n; i ++ {
		z -= (s * s - z) / 2*z
	}
	return  z, nil
}


type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, e := rot.r.Read(b)

	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, e
}

type Image struct {
	w, h int
	c    uint8
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.c + uint8(x*y+x*y), img.c + uint8(x*y+x*y), 255, 255}
}
func ColorImage() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
