package session

import (
	"crypto/rand"
	"fmt"
)

func Token() string {
	b := make([]byte, 24)

	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", b)
}
