package main

import (
	"fmt"

	"github.com/hgkcho/aja/internal/tool"
	"github.com/hgkcho/aja/pkg/fooooo"
)

func main() {
	t := tool.NewTool("aaaa")
	fmt.Println(t.Name())
	s := fooooo.NewFoo("bbbb")
	fmt.Println(s.Name())
}
