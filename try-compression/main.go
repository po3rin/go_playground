package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	// open test.jpg
	file, err := os.Open("gophers.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// ファイル情報取得
	fileinfo, staterr := file.Stat()
	if staterr != nil {
		fmt.Println(staterr)
		return
	}

	fmt.Println(fileinfo.Size())
	ext := filepath.Ext(fileinfo.Name())

	var img image.Image
	switch ext {
	case ".jpeg", ".jpg":
		img, err = jpeg.Decode(file)
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	case ".png":
		img, err = png.Decode(file)
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Lanczos3 で圧縮
	m := resize.Resize(0, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resize" + ext)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	// ファイル情報取得
	fileinfo, staterr = out.Stat()

	if staterr != nil {
		fmt.Println(staterr)
		return
	}
	// ファイルサイズを表示
	fmt.Println(fileinfo.Size())
}
