package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mholt/binding"
)

type ContactForm struct {
	User struct {
		ID int
	}
	FieldWords FieldWords `json:"field_words"`
}

type FieldWords []FieldWord

type FieldWord struct {
	FieldName string `json:"field_name"`
	Word      Words  `json:"word`
}

type Words []Word

type Word interface {
	GetQuery(string) string
}

func (cf *ContactForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.User.ID: "user_id",
		&cf.FieldWords: binding.Field{
			Form:     "field_words",
			Required: false,
		},
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	// bind word_field only
	var buf bytes.Buffer
	req.Body = NewTeeReadCloser(req.Body, &buf)
	decoder := json.NewDecoder(req.Body)

	var Binder struct {
		FieldWords []struct {
			FieldName string      `json:"field_name"`
			Word      interface{} `json:"word"`
		} `json:"field_words"`
	}
	err := decoder.Decode(&Binder)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(Binder)

	// secondly bind.
	rawData := buf.Bytes()
	contactForm := new(ContactForm)
	// if errs := binding.Bind(rawData, contactForm); errs != nil {
	// 	http.Error(resp, errs.Error(), http.StatusBadRequest)
	// 	return
	// }

	if err := json.Unmarshal(rawData, contactForm); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println(contactForm)
}

type readCloser struct {
	io.Reader
	io.Closer
}

func NewTeeReadCloser(rc io.ReadCloser, w io.Writer) io.ReadCloser {
	return &readCloser{
		Reader: io.TeeReader(rc, w),
		Closer: rc,
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
