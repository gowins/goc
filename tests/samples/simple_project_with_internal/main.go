package main

import (
	"example.com/simple-project/foo"
	"example.com/simple-project/internal"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer func() {
		fmt.Println("close")
	}()

	foo.Bar1()
	foo.Bar2()
	internal.Hello()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Printf("get a signal %s", si.String())
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
