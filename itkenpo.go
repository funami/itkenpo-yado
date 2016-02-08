package main

import (
	"fmt"
	"itkenpo/yado"
)

func main() {
	m := yado.GetYadoInfo()
	for k, v := range m {
		fmt.Printf("%s : %s\n", k, v)
	}
}
