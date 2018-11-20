package main

import (
	"fmt"
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
	Word      Words  `json:"word"`
}

type Words []Word

type Word string

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
	contactForm := new(ContactForm)
	if errs := binding.Bind(req, contactForm); errs != nil {
		http.Error(resp, errs.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "from:    %d\n", contactForm.User.ID)
	fmt.Fprintf(resp, "field_words: %s\n", contactForm.FieldWords)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
