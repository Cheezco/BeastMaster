package debug

import (
	"fmt"
	"math"
	"strings"
)

func Log(a ...any) {
	if !IsDebugMode() {
		return
	}
	printPrefix()
	fmt.Println(a...)
}

func Logf(format string, a ...any) {
	if !IsDebugMode() {
		return
	}

	printPrefix()
	//fmt.Printf(format+"\n", append([]interface{}{}, a...))
	fmt.Printf(format+"\n", a...)
}

func printPrefix() {
	name := GetExecutableName()
	fmt.Printf("[DEBUG] | %s%s | ", name, strings.Repeat(" ", int(math.Max(float64(20-len(name)), 0))))
}
