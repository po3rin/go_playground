package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	once = new(sync.Once)
	c    *cache.Cache
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8880", nil))
}

type user struct {
	name string
	age  int
}

func handler(w http.ResponseWriter, r *http.Request) {

	once.Do(func() {
		c = cache.New(1*time.Minute, 2*time.Minute)
	})

	data, found := c.Get("gm")
	if found {
		fmt.Println("1 gm: ", data.(user))
	}

	data, found = c.Get("abc")
	if found {
		fmt.Println("1 abc: ", data.(user))
	}

	u := user{
		name: "po3rin",
		age:  27,
	}
	c.Set("gm", u, cache.DefaultExpiration)

	data, found = c.Get("gm")
	if found {
		fmt.Println("2 gm: ", data.(user))
	}

	data, found = c.Get("abc")
	if found {
		fmt.Println("2 abc: ", data.(user))
	}
}
