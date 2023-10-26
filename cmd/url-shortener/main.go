package main

import (
	"fmt"
	//"log/slog"
	"url-shortener/internal/config"
)

func main() {
	// init config: cleanenv (read json, yaml etc + tags)
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// init logger: slog 

	// init storage: sqlite

	// init router: chi (compatible with "net/http"), chi render

	// run server
}
