package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	uploadEventsQueue := make(chan string, 100)

	go mainService(&wg, uploadEventsQueue)
	go feedService(&wg, uploadEventsQueue)
}

func mainService(wg *sync.WaitGroup, queue chan<- string) {
	defer wg.Done()
	http.HandleFunc("/uploadPic", uploadPic(queue))

	http.ListenAndServe(":8080", nil)
}

func feedService(wg *sync.WaitGroup, queue <-chan string) {
	defer wg.Done()
	ln, _ := net.Listen("tcp", "localhost:9090")

	// Only one listener will recieve the event, hence we need a chanle per user listenin for the feed
	var clientsQueues []chan string
	go func() {
		for {
			newEvent := <-queue
			fmt.Println("Sending to multiple people ", len(clientsQueues))
			for _, q := range clientsQueues {
				q <- newEvent
			}
		}
	}()

	for {
		fmt.Println("accepting connection")
		conn, err := ln.Accept()
		fmt.Println("accepted ")
		if err != nil {
			log.Println("Failed listening for connegtion ", err)
		}

		broadcastQueue := make(chan string, 100)
		clientsQueues = append(clientsQueues, broadcastQueue)

		go handleConnection(conn, broadcastQueue)
	}
}

func handleConnection(conn net.Conn, queue <-chan string) {
	// store incoming data
	user := make([]byte, 1024)
	_, err := conn.Read(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New user has connected ", string(user))

	for {
		event := <-queue
		conn.Write([]byte(event))
	}
}

func uploadPic(queue chan<- string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, _ := ioutil.ReadAll(r.Body)
		user, _, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		log.Printf("User %s uploaded pic %s\n", user, body)
		queue <- fmt.Sprintf("User %s uploaded pic %s\n", user, body)
		// io.WriteString(w, fmt.Sprintf("User %s uploaded pic %s\n", user, body))
	}
}
