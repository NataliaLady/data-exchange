/*Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом: */
/*1. Сделайте обработчик создания пользователя.
У пользователя должны быть следующие поля: имя, возраст и массив друзей.
Пользователя необходимо сохранять в мапу. Пример запроса: */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server is running")
	con, err := lis.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		line, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("line reserved:", string(line))
		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatal(err)
		}

	}
}
