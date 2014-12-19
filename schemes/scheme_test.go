package schemes

import (
	"image/color"
	"reflect"
	"testing"
)

func TestFromImage(t *testing.T) {
	colors, err := FromImage("nonexistent.png")
	if err == nil {
		t.Errorf("Expected failure on nonexistent flie, got %v", colors)
	}
	colors, err = FromImage("classic.go")
	if err == nil {
		t.Errorf("Expected failure on classic.go, got %v", colors)
	}

	colors, err = FromImage("classic.png")
	if err != nil {
		t.Fatalf("Failed loading classic.png: %v", err)
	}
	if !reflect.DeepEqual(colors, Classic) {
		t.Errorf("Got %v, wanted %v", colors, Classic)
	}
}

func TestBuild(t *testing.T) {
	colors := Build(SchemeSpec{
		SchemeRange{
			From:  color.NRGBA{255, 0, 0, 255},
			To:    color.NRGBA{0, 255, 0, 255},
			Steps: 10,
		},
		SchemeRange{
			From:  color.NRGBA{0, 255, 0, 255},
			To:    color.NRGBA{0, 0, 255, 255},
			Steps: 10,
		}})

	exp := []color.Color{
		mycolor{R: 65535, G: 0, B: 0, A: 65535},
		mycolor{R: 58981.5, G: 6553.5, B: 0, A: 65535},
		mycolor{R: 52428, G: 13107, B: 0, A: 65535},
		mycolor{R: 45874.5, G: 19660.5, B: 0, A: 65535},
		mycolor{R: 39321, G: 26214, B: 0, A: 65535},
		mycolor{R: 32767.5, G: 32767.5, B: 0, A: 65535},
		mycolor{R: 26214, G: 39321, B: 0, A: 65535},
		mycolor{R: 19660.5, G: 45874.5, B: 0, A: 65535},
		mycolor{R: 13107, G: 52428, B: 0, A: 65535},
		mycolor{R: 6553.5, G: 58981.5, B: 0, A: 65535},

		mycolor{R: 0, G: 65535, B: 0, A: 65535},
		mycolor{R: 0, G: 58981.5, B: 6553.5, A: 65535},
		mycolor{R: 0, G: 52428, B: 13107, A: 65535},
		mycolor{R: 0, G: 45874.5, B: 19660.5, A: 65535},
		mycolor{R: 0, G: 39321, B: 26214, A: 65535},
		mycolor{R: 0, G: 32767.5, B: 32767.5, A: 65535},
		mycolor{R: 0, G: 26214, B: 39321, A: 65535},
		mycolor{R: 0, G: 19660.5, B: 45874.5, A: 65535},
		mycolor{R: 0, G: 13107, B: 52428, A: 65535},
		mycolor{R: 0, G: 6553.5, B: 58981.5, A: 65535}}

	if !reflect.DeepEqual(colors, exp) {
		t.Errorf("Got %v, wanted %v", colors, exp)
	}
}

func TestColorRGBA(t *testing.T) {
	c := mycolor{0x1fff, 0x1fff, 0x1fff, 0xffff}
	r, g, b, a := c.RGBA()
	for g, w := range map[uint32]uint32{r: 0x1f1f, g: 0x1f1f, b: 0x1f1f, a: 0xffff} {
		if g != w {
			t.Errorf("Expected %x, got %x")
		}
	}
}

func TestSchemeSpec(t *testing.T) {
	ss := SchemeSpec{
		SchemeRange{
			From:  color.NRGBA{255, 0, 0, 255},
			To:    color.NRGBA{0, 255, 0, 255},
			Steps: 256,
		}}

	if ss.ColorModel() != color.NRGBAModel {
		t.Errorf("Expected NRGBAModel, got %v", ss.ColorModel())
	}

	bounds := ss.Bounds()
	if bounds.Size().X != 32 {
		t.Errorf("Expected width of 32, got %v", bounds.Size().X)
	}
	if bounds.Size().Y != 256 {
		t.Errorf("Expected width of 256, got %v", bounds.Size().Y)
	}

	r1, g1, b1, a1 := ss.At(0, 0).RGBA()
	if r1 != 65535 || g1 != 0 || b1 != 0 || a1 != 65535 {
		t.Errorf("Expected rgb1 == 65535, 0, 0, 65535, got %v %v %v %v", r1, g1, b1, a1)
	}

	r2, g2, b2, a2 := ss.At(31, 255).RGBA()
	if r2 != 0 || g2 != 65278 || b2 != 0 || a2 != 65535 {
		t.Errorf("Expected rgb2 == 65278, 0, 0, 65535, got %v %v %v %v", r2, g2, b2, a2)
	}
}
