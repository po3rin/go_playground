package main // import "try-expvar"

import (
	"net/http"
	_ "try-expvar/metrics"
)

func main() {
	http.ListenAndServe(":9999", nil)
}
