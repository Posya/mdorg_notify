package main

import "fmt"

var opts struct {
	tgKey string `short:"k" long:"tgkey" env:"TG_KEY" description:"API key for Telegram Bot API" required:"true"`
}

func main() {
	fmt.Printf("tgKey\t'%s'\n", opts.tgKey)
}
