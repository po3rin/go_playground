package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	// select using where.
	users := sq.Select("*").From("posts")
	sql, args, _ := users.
		Where(sq.Eq{"name": "moe"}).
		Where(sq.Eq{"name": "ddd"}).
		ToSql()
	fmt.Println(sql)
	fmt.Println(args)
}
