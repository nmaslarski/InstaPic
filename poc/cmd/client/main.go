package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Username: ")
	scanner.Scan()
	user := scanner.Text()
	fmt.Printf("Wellcome %x\n To upload text please just type it\n", user)

	go uploadPictures(&wg, user)
	go subscribeToFeed(&wg, user)
}

func uploadPictures(wg *sync.WaitGroup, user string) {
	defer wg.Done()

	scanner := bufio.NewScanner(os.Stdin)
	client := http.Client{}
	for {
		scanner.Scan()
		text := scanner.Bytes()

		req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/uploadPic", bytes.NewReader(text))
		req.SetBasicAuth(user, "")
		resp, _ := client.Do(req)
		_ = resp

		// buf := new(bytes.Buffer)
		// buf.ReadFrom(resp.Body)
		// fmt.Println(buf.String())
	}
}

func subscribeToFeed(wg *sync.WaitGroup, user string) {
	defer wg.Done()

	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	conn.Write([]byte(user))

	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		fmt.Println("Reading from conn")
		if err != nil {
			log.Printf("Connection failed for user %s\n", user)
			break
		}

		if len(str) > 0 {
			fmt.Println(str)
		}
	}
}
