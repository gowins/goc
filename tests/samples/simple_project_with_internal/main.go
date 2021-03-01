package main

import (
	"example.com/simple-project/foo"
	"example.com/simple-project/internal"
	"time"
)

func main() {
	foo.Bar1()
	foo.Bar2()
	internal.Hello()

	time.Sleep(time.Minute)
}
