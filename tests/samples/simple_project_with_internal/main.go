package main

import (
	"example.com/simple-project/foo"
	"example.com/simple-project/internal"
)

func main() {
	foo.Bar1()
	foo.Bar2()
	internal.Hello()
	select {}
}
