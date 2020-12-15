package main

import "fmt"

func main() {
	var balance int32 = 1500000000       // 15 миллионов в копейках
	var transfer int32 = 1000000000      // 10 миллионов в копейках
	total := balance + transfer // int32 + int32 будет int32
	fmt.Println(total)
}
