// this is an introduction.
package main

import (
	"fmt"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Goodbye World !")
}

// Print the running command of this test
func TestPrintRunningCommand(t *testing.T) {
	fmt.Println("run command:")
	// /private/var/folders/_m/lyk8b0b92zb4wgw67qlbrx_c0000gn/T/GoLand/___TestEchoProgrammingAchieving_in_demo_test_go.test
	fmt.Println(os.Args[0])
	fmt.Println("the parameters of command:")
	// [-test.v -test.paniconexit0 -test.run ^\QTestEchoProgrammingAchieving\E$]
	fmt.Println(os.Args[1:])
}
