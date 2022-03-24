//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"fmt"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

func start(process func(), done chan bool) {
	process()
	done <- true
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := make(chan bool)
	ticker := time.Tick(time.Second)
	for {
		select {
		case <-done:
			fmt.Println(u)
			return true
		case <-ticker:
			u.TimeUsed++
			if u.TimeUsed > 10 && u.IsPremium == false {
				return false
			}
		}
	}
}

func main() {
	RunMockServer()
}
