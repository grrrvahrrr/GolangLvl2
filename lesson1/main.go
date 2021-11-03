//Lesson 1 Homework

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	letsPanicAndRecover()
}

func letsPanicAndRecover() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(checkError("we paniced, but survived"))
		}
	}()
	//Let's panic
	var a int
	fmt.Println(1 / a)
}

func createOneMillionFiles() {
	func() {
		mln := 1000
		for i := 0; i < mln; i++ {
			file, err := os.Create("onemlnfiles/testfile" + strconv.Itoa(i))
			if err != nil {
				fmt.Println(checkError("Couldn't create file"))
			}
			defer func() {
				err := file.Close()
				if err != nil {
					fmt.Println(checkError("Couldn't close file"))
				}
			}()
		}
	}()
}
