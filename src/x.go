package src

import "fmt"

/*
go source code to asm
1. go tool compile -S x.go

// 看链接之后的汇编
go build -o x.exe x.go
go tool objdump -s print3 x.exe
*/

// print3 print num 3 to stdout
func print3() {
	fmt.Println(3)
}
