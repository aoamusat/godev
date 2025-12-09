package learn

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/google/uuid"
)

type User struct {
	name    string
	email   string
	age     int
	address string
}

func main() {
	users := make(map[string]User)
	for i := range 10 {
		userId := uuid.NewString()
		users[userId] = User{
			name:    "User" + strconv.Itoa(i),
			email:   "user" + strconv.Itoa(i) + "@example.com",
			age:     20 + i,
			address: "Address" + strconv.Itoa(i),
		}
	}

	for id, user := range users {
		fmt.Printf("ID: %s\nName: %s\nEmail: %s\nAge: %d\nAddress: %s\n**************************************\n", id, user.name, user.email, user.age, user.address)
	}
}

// sortMapByKey returns a new map with the same key-value pairs as the input map,
// but with keys sorted in lexicographic order. The original map is not modified.
// Note: Go maps are unordered, so the returned map's iteration order is still undefined;
// use the sorted keys separately if you need guaranteed ordering.
func sortMapByKey(m map[string]any) map[string]any {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	sortedMap := make(map[string]any, len(m))
	for _, key := range keys {
		sortedMap[key] = m[key]
	}
	return sortedMap
}
