package main

import (
	"flag"
	"log"
)

// go run test1.go -name=测试 -n=煎鱼
func main() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅2", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}

