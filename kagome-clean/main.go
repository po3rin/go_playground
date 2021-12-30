package main

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func main() {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}

	// tokenize
	fmt.Println("---tokenize---")
	tokens := t.Tokenize("すももももももももものうち")
	for _, token := range tokens {
		fmt.Println(len(token.Features()))
		fmt.Printf("%+v\n", token.Features())
		fmt.Printf("%+v\n", token.Surface)
	}
}
