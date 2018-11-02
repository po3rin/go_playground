package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shiroyagicorp/double_array"
)

func main() {
	data := []string{
		"abc",
		"abcd",
		"文字",
		"全角",
		"全角文字",
		"x",
		"y",
		"z",
		"xyzabc",
		"good",
		"漢字",
	}

	items := make([]double_array.Item, len(data))

	for i, item := range data {
		items[i] = double_array.Item(item)
	}

	// Trieの構築
	da, err := double_array.NewDoubleArray(items)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// 検索
	// Trieに引数の文字列が含まれていれば内部IDを返します。
	// 含まれていなければ double_array.ItemNotFound を返します。
	itemID := da.Lookup([]rune("漢字"))
	fmt.Printf("「漢字」のItemID=%d\n", itemID)

	itemIDNotFound := da.Lookup([]rune("存在しないエントリ"))
	fmt.Println(
		"存在しないエントリの場合ItemNotFoundになります: ",
		itemIDNotFound == double_array.ItemNotFound,
	)

	// 内部IDを文字列に戻す
	dict := double_array.ToInverseID(da)
	original := double_array.Deserialize(da, itemID, dict)
	fmt.Printf("ItemID=%d の復元結果=%s\n", itemID, original)

	// 文書中からの一括抽出
	// 後方からの最長一致で抽出します。
	// 一度抽出した場所は再度抽出されません。
	// 以下の例では「全角文字」「漢字」は抽出されますが「全角」「文字」は抽出されません。
	text := []rune("全角文字を使うと漢字を表現できます。")
	da.Scan(text, func(i, j int, id double_array.ItemID) {
		fmt.Printf("(%d, %d): %s\n", i, j, string(text[i:j]))
	})

	// でき上がった構造を保存しておきたいときはSerializeを呼び出し、返ってきたバイト配列を保存して下さい。
	serialized, _ := da.Serialize()
	if err != nil {
		log.Fatal(err)
	}

	// デシリアライズするときはNewDoubleArrayFromBytesを使って下さい。
	da, _ = double_array.NewDoubleArrayFromBytes(serialized)
	itemID = da.Lookup([]rune("xyzabc"))
	fmt.Printf("「xyzabc」のItemID=%d\n", itemID)
}