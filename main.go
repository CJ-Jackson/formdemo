package main

import (
	"fmt"
	_ "github.com/CJ-Jackson/formdemo/src"
	_ "github.com/CJ-Jackson/formdemo/web"
	"github.com/cjtoolkit/cli"
	"runtime/debug"
)

func main() {
	cli.SetCmdStartEvent(func(cmdName string, args []string) {
		fmt.Printf("Executing '%s'", cmdName)
		fmt.Println()
	})

	cli.SetCmdFinishEvent(func(cmdName string, args []string, recv interface{}) {
		if nil != recv {
			d := debug.Stack()

			fmt.Println("Oh no something did not go right")

			fmt.Printf("Info: %v", recv)

			fmt.Println("Args:")
			fmt.Println()
			fmt.Println(args)

			fmt.Println("Stack:")
			fmt.Println()
			fmt.Printf("%s", d)
		}
	})

	fmt.Println("-- Formdemo 2.0 --")
	fmt.Println()

	cli.Run()
}
