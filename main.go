package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

// сетевой адрес
const addr = "0.0.0.0:12345"

// протокол сетевой службы
const protoс = "tcp4"

func main() {
	// запуск сетевой службы по протоколу TCP на порту 12345
	listener, err := net.Listen(protoс, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// бесконечный цикл обработки подключений
	for {
		// принимаем подключение
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// вызов обработчика подключения
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for range time.Tick(time.Second * 3) {
		conn.Write([]byte(readProverb("proverbs.txt") + "\n"))
	}

	// Закрытие соединения.
	conn.Close()
}

func readProverb(fileName string) (line string) {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	x := len(lines)
	line = lines[rand.Intn(x)]
	readFile.Close()

	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
