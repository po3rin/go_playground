package image_test

import (
	"bytes"
	"flag"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"reflect"
	"testing"
)

var genGoldenFiles = flag.Bool("gen_golden_files", false, "whether to generate the TestXxx golden files.")

// func TestResizeJPEG(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		imgFilename    string
// 		goldenFilename string
// 	}{
// 		{
// 			name:           "x1.0",
// 			imgFilename:    "testdata/gopher.jpeg",
// 			goldenFilename: "testdata/resize_golden_1.jpeg",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			f, err := os.Open(tt.imgFilename)
// 			if err != nil {
// 				t.Fatalf("failed to open file\nerr: %v", err)
// 			}
// 			defer f.Close()
// 			img, _, err := image.Decode(f)
// 			if err != nil {
// 				t.Fatalf("failed to decode file\nerr: %v", err)
// 			}

// 			// any image processing ...
// 			got := img

// 			if *genGoldenFiles {
// 				goldenFile, err := os.Create(tt.goldenFilename)
// 				if err != nil {
// 					t.Errorf("failed to create file\nerr: %v", err)
// 				}
// 				defer goldenFile.Close()
// 				err = jpeg.Encode(goldenFile, got, &jpeg.Options{Quality: 100})
// 				if err != nil {
// 					t.Errorf("failed to encode file\nerr: %v", err)
// 				}
// 				return
// 			}

// 			// want
// 			f, err = os.Open(tt.goldenFilename)
// 			if err != nil {
// 				t.Fatalf("failed to open file\nerr: %v", err)
// 			}
// 			defer f.Close()
// 			want, _, err := image.Decode(f)
// 			if err != nil {
// 				t.Fatalf("failed to decode file\nerr: %v", err)
// 			}

// 			// compare RGBA.
// 			if !reflect.DeepEqual(convertRGBA(got), convertRGBA(want)) {
// 				t.Errorf("actual image differs from golden image")
// 				return
// 			}

// 			// b := want.Bounds()
// 			// for y := b.Min.Y; y < b.Max.Y; y++ {
// 			// 	for x := b.Min.X; x < b.Max.X; x++ {
// 			// 		if !eq(convertRGBA(got).At(x, y), convertRGBA(want).At(x, y)) {
// 			// 			t.Errorf("unexpected color in [%v,%v]\ngot: %v\nwant:%v\n", x, y, convertRGBA(got).At(x, y), convertRGBA(want).At(x, y))
// 			// 			return
// 			// 		}
// 			// 	}
// 			// }
// 		})
// 	}
// }
func TestResizePNG(t *testing.T) {
	tests := []struct {
		name           string
		imgFilename    string
		goldenFilename string
	}{
		{
			name:           "x1.0",
			imgFilename:    "testdata/gopher.jpeg",
			goldenFilename: "testdata/resize_golden_1.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.imgFilename)
			if err != nil {
				t.Fatalf("failed to open file\nerr: %v", err)
			}
			defer f.Close()
			img, _, err := image.Decode(f)
			if err != nil {
				t.Fatalf("failed to decode file\nerr: %v", err)
			}

			// any image processing ...
			got := img

			if *genGoldenFiles {
				goldenFile, err := os.Create(tt.goldenFilename)
				if err != nil {
					t.Errorf("failed to create file\nerr: %v", err)
				}
				defer goldenFile.Close()
				err = png.Encode(goldenFile, got)
				if err != nil {
					t.Errorf("failed to encode file\nerr: %v", err)
				}
				return
			}

			// want
			f, err = os.Open(tt.goldenFilename)
			if err != nil {
				t.Fatalf("failed to open file\nerr: %v", err)
			}
			defer f.Close()
			want, _, err := image.Decode(f)
			if err != nil {
				t.Fatalf("failed to decode file\nerr: %v", err)
			}

			// compare RGBA.
			if !reflect.DeepEqual(convertRGBA(got), convertRGBA(want)) {
				t.Errorf("actual image differs from golden image")
				return
			}

			// b := want.Bounds()
			// for y := b.Min.Y; y < b.Max.Y; y++ {
			// 	for x := b.Min.X; x < b.Max.X; x++ {
			// 		if !eq(convertRGBA(got).At(x, y), convertRGBA(want).At(x, y)) {
			// 			t.Errorf("unexpected color in [%v,%v]\ngot: %v\nwant:%v\n", x, y, convertRGBA(got).At(x, y), convertRGBA(want).At(x, y))
			// 			return
			// 		}
			// 	}
			// }
		})
	}
}

func convertRGBA(raw image.Image) *image.RGBA {
	want, ok := raw.(*image.RGBA)
	if !ok {
		b := raw.Bounds()
		want = image.NewRGBA(b)
		draw.Draw(want, b, raw, b.Min, draw.Src)
	}
	return want
}

func eq(c0, c1 color.Color) bool {
	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()
	return r0 == r1 && g0 == g1 && b0 == b1 && a0 == a1
}

// RGBAImageEqual returns true if the parameter images a and b match
// or false if otherwise.
func RGBAImageEqual(a, b *image.RGBA) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	for y := 0; y < a.Bounds().Dy(); y++ {
		for x := 0; x < a.Bounds().Dx(); x++ {
			pos := y*a.Stride + x*4
			if a.Pix[pos+0] != b.Pix[pos+0] {
				return false
			}
			if a.Pix[pos+1] != b.Pix[pos+1] {
				return false
			}
			if a.Pix[pos+2] != b.Pix[pos+2] {
				return false
			}
			if a.Pix[pos+3] != b.Pix[pos+3] {
				return false
			}
		}
	}
	return true
}

func checkBoundsAndPix(tb testing.TB, b1, b2 image.Rectangle, pix1, pix2 []uint8) bool {
	tb.Helper()
	if !b1.Eq(b2) {
		tb.Errorf("got unexpected result")
	}
	if !bytes.Equal(pix1, pix2) {
		tb.Errorf("got unexpected result")
	}
	return true
}

func BenchmarkEqNormal(b *testing.B) {
	f, err := os.Open("./testdata/gopher.jpeg")
	if err != nil {
		b.Fatalf("failed to open file\nerr: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		b.Fatalf("failed to decode file\nerr: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			for x := 0; x < img.Bounds().Dx(); x++ {
				if !eq(img.At(x, y), img.At(x, y)) {
					b.Error("failed to eq")
				}
			}
		}
	}
}

func BenchmarkEqWithStride(b *testing.B) {
	f, err := os.Open("./testdata/gopher.jpeg")
	if err != nil {
		b.Fatalf("failed to open file\nerr: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		b.Fatalf("failed to decode file\nerr: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rgba := convertRGBA(img)
		if !RGBAImageEqual(rgba, rgba) {
			b.Fatalf("failed to decode file\nerr: %v", err)
		}
	}
}

func BenchmarkEqWithBytes(b *testing.B) {
	f, err := os.Open("./testdata/gopher.jpeg")
	if err != nil {
		b.Fatalf("failed to open file\nerr: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		b.Fatalf("failed to decode file\nerr: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rgba := convertRGBA(img)
		if !checkBoundsAndPix(b, rgba.Bounds(), rgba.Bounds(), rgba.Pix, rgba.Pix) {
			b.Fatalf("failed to decode file\nerr: %v", err)
		}
	}
}

func BenchmarkEqWithReflect(b *testing.B) {
	f, err := os.Open("./testdata/gopher.jpeg")
	if err != nil {
		b.Fatalf("failed to open file\nerr: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		b.Fatalf("failed to decode file\nerr: %v", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rgba := convertRGBA(img)
		if !reflect.DeepEqual(rgba, rgba) {
			b.Fatalf("failed to decode file\nerr: %v", err)
		}
	}
}
