package hasher

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pass := "password"
	hash, err := HashPassword(pass)
	if err != nil {
		t.Errorf("Error #{err}:")
	}
	fmt.Println(hash)
}

func TestCheckPasswordHash(t *testing.T) {
	pass := "password"
	hash := "$2a$08$cuzJYqNhKhGY2bxYCHExV.kUwUwFSQUrwHZGISR7TXveseNozjpru"
	b := CheckPasswordHash(pass, hash)
	fmt.Println(b)
	// if !b {
	// 	t.Errorf("Error in CheckPasswordHash")
	// }
}
