package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/nfnt/resize"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Open("assets/")
	defer dir.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	allImageNames, err := dir.Readdirnames(-1) // それぞれの画像ファイルの名前を配列に格納します
	if err != nil {
		log.Fatalln("No files")
	}
	var decodeAllImages []image.Image
	for _, imageName := range allImageNames { // 全ての画像をデコード、リサイズしてdecodeAllImageseに格納します
		file, _ := os.Open("assets/" + imageName)
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodeImage, _, err := image.Decode(file)
		resizedDecodeImage := resize.Resize(300, 0, decodeImage, resize.Lanczos3) // サイズを揃えるために横幅を300に固定します
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodeAllImages = append(decodeAllImages, resizedDecodeImage)
	}
	writeImageWithTemplate(w, decodeAllImages)
}

// writeImageWithTemplateで画像をエンコードします。
func writeImageWithTemplate(w http.ResponseWriter, decodeAllImages []image.Image) {
	var encordImages []string
	for _, decodeImage := range decodeAllImages {
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, decodeImage, nil); err != nil {
			log.Fatalln("Unable to encode image.")
		}
		str := base64.StdEncoding.EncodeToString(buffer.Bytes())
		encordImages = append(encordImages, str)
	}
	data := map[string]interface{}{"Images": encordImages}
	renderTemplate(w, data)
}

// renderTemplateで渡された画像をテンプレートエンジンに渡します。
func renderTemplate(w http.ResponseWriter, data interface{}) {
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
