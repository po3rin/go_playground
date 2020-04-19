package main

import (
	"encoding/csv"
	"log"
	"os"

	iconv "github.com/djimenez/iconv-go"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	// O_WRONLY:書き込みモード開く, O_CREATE:無かったらファイルを作成
	file, err := os.OpenFile("people.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()

	err = file.Truncate(0) // ファイルを空っぽにする(2回目以降用)
	failOnError(err)

	converter, err := iconv.NewWriter(file, "utf-8", "sjis")
	failOnError(err)

	writer := csv.NewWriter(converter)
	writer.Write([]string{"山田", "20"})
	writer.Write([]string{"田中", "21"})
	writer.Write([]string{"佐藤", "22"})
	writer.Flush()
}
