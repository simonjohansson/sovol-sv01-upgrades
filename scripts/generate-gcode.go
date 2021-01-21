//usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"os"
)

func startGcode() string {
	return `
G28
G1 X150 Y129 Z10 F12000
G4 P10000


`
}

func cross() (output string) {
	output = startGcode()

	start := 5
	end := 240
	for start <= end {
		output += fmt.Sprintf("G1 X%d Y%d \n", start, start)
		output += fmt.Sprintf("G1 X%d Y%d \n", end, end)
		if end-start <= 20 {
			start += 1
			end -= 1
		} else {
			start += 10
			end -= 10
		}
	}

	start = 5
	end = 240
	for start <= end {
		output += fmt.Sprintf("G1 X%d Y%d \n", start, end)
		output += fmt.Sprintf("G1 X%d Y%d \n", end, start)
		if end-start <= 20 {
			start += 1
			end -= 1
		} else {
			start += 10
			end -= 10
		}
	}

	return

}

func xyTest(axis string, start int, end int) (output string) {
	output = startGcode()

	for start <= end {
		output += fmt.Sprintf("G1 %s%d \n", axis, start)
		output += fmt.Sprintf("G1 %s%d \n", axis, end)
		if end-start <= 20 {
			start += 1
			end -= 1
		} else {
			start += 10
			end -= 10
		}
	}
	return
}

func xTest() (output string) {
	return xyTest("X", 0, 300)
}

func yTest() (output string) {
	return xyTest("Y", 5, 240)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please call with './generate-gcode.go x|y|crazy' ")
		os.Exit(-1)
	}
	switch os.Args[1] {
	case "x":
		fmt.Println(xTest())
	case "y":
		fmt.Println(yTest())
	case "cross":
		fmt.Println(cross())
	default:
		fmt.Println("Must either be 'x', 'y' or 'crazy'")
		os.Exit(-1)
	}
}
