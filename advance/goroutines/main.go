package goroutines

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type MessageData struct {
	Chats   []string
	Friends []string
}

func GetUserById(id int) User {
	return User{
		ID:    uuid.New(),
		Name:  fmt.Sprintf("User%d", id),
		Email: fmt.Sprintf("user%d@example.com", id),
	}
}

func GetUserChats(userID uuid.UUID, ch chan<- *MessageData, wg *sync.WaitGroup) {
	// Simulate fetching user chats with a delay
	defer wg.Done()
	time.Sleep(3 * time.Second)
	chats := []string{
		"Chat with Alice",
		"Chat with Bob",
		"Chat with Charlie",
	}
	ch <- &MessageData{Chats: chats}
}

func GetUserFriends(userID uuid.UUID, ch chan<- *MessageData, wg *sync.WaitGroup) {
	// Simulate fetching user friends with a delay
	defer wg.Done()
	time.Sleep(4 * time.Second)
	friends := []string{
		"John Doe",
		"Jane Smith",
		"Bob Johnson",
	}
	ch <- &MessageData{Friends: friends}
}

func UpdateUserEmail(userID uuid.UUID, newEmail string) {
	// Simulate updating user email with a delay
	time.Sleep(2 * time.Second)
	fmt.Printf("User %s email updated to %s\n", userID, newEmail)
}
