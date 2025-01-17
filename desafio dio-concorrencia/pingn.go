package main

import (
	"fmt"
	"time"
)

func pingPong(c chan string) {
	for {
		c <- "ping"
		time.Sleep(1 * time.Second)
		c <- "pong"
		time.Sleep(1 * time.Second)
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string)

	go pingPong(c) // Alterna entre "ping" e "pong"
	go imprimir(c) // Exibe as mensagens recebidas do canal

	// Mantém o programa rodando
	var entrada string
	fmt.Scanln(&entrada)
}
