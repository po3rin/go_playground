package main

import (
	"log"

	"github.com/po3rin/bloomfilter"
)

func main() {
	s := "reg"
	strList := []string{
		"reg",
		"fw",
		"eb",
		"gefg",
		"ibgo",
		"vbf",
		"pk",
		"gr",
		"trtb",
		"pgkf",
		"rebd",
		"po",
	}

	strNum := len(strList)
	filter, err := bloomfilter.NewBloomFilter(1000, strNum)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range strList {
		filter.Add(s)
	}

	ffp := filter.EstimateFPP()
	log.Printf("ffp: %e", ffp)

	if filter.MightContain(s) {
		log.Print("maybe contain!!")
		return
	}
	log.Print("no")
}
