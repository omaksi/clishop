package rdg

import "testing"

func TestGetUsers(t *testing.T) {
	users := GetUsers()
	if len(users) == 0 {
		t.Error("No users found")
	}
}
