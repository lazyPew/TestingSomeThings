package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

/*
Поскольку сервер будет запущен на локальном компьтере на порте 4545,
то клиент подлючается к этому адресу: net.Dial("tcp", "127.0.0.1:4545")

После этого к серверу будет отправляться запрос, и с помощью вызова
io.Copy(os.Stdout, conn) выводим полученный ответ на консоль.
*/

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
