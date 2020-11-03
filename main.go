package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int

	fmt.Println("Ingrese la cantidad de cadenas: ")
	fmt.Scanln(&n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Println("Ingrese cadena: ")
		fmt.Scanln(&s[i])
	}
	// Sort
	sort.Strings(s)

	// OS
	aFile, err := os.OpenFile("asecendente.txt", os.O_CREATE|os.SEEK_SET, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	dFile, err := os.OpenFile("descendente.txt", os.O_CREATE|os.SEEK_SET, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aFile.Close()
	defer dFile.Close()

	for _, v := range s {
		aFile.WriteString(v + "\n")
	}
	for i := len(s) - 1; i >= 0; i-- {
		dFile.WriteString(s[i] + "\n")
	}
	fmt.Println("Respaldo completo")
}
