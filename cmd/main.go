package main

import (
	"fmt"
)

func main() {
	var p string
	fmt.Println("Para limpar a tela")
	fmt.Println("depois de escrever isso")
	fmt.Scan(&p)
	fmt.Println(p)
	// cmd := exec.Command("clear")
	// log.Printf("Running command and waiting for it to finish...")
	// _, err := cmd.Output()
	// fmt.Println("depois de escrever limpar a tela")
	// log.Printf("Command finished with error: %v", err)
}
