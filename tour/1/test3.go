package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("xizhao:" + value)
	return nil
}

// go run test3.go -name=测试
func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}
