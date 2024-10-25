package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func mainGoRoutines() {
	for i := 0; i < 5; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}

/*
Sample output of this is
4 5
3 4 5
1 2 3 4 5
0 1 2 3 4 5
2 3 4 5

This is out of order because the go routine is managed by the Go scheduler which runs these commands at undeterministic times
*/
