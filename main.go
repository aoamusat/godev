package main

import (
	"log"
	"math/rand/v2"
	"sync"
	"time"

	grt "a4m.dev/godev/advance/goroutines"
)

func main() {
	startTime := time.Now()
	randomId := rand.IntN(100)
	ch := make(chan *grt.MessageData, 1024)
	var wg sync.WaitGroup
	wg.Add(2)
	user := grt.GetUserById(randomId)

	go grt.GetUserChats(user.ID, ch, &wg)
	go grt.GetUserFriends(user.ID, ch, &wg)
	wg.Wait()
	close(ch)

	for msg := range ch {
		if msg.Chats != nil {
			log.Printf("User Chats: %v\n", msg.Chats)
		}
		if msg.Friends != nil {
			log.Printf("User Friends: %v\n", msg.Friends)
		}
	}
	log.Printf("\n\nTotal execution time: %.2f seconds\n", time.Since(startTime).Seconds())
}
