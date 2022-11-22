package iterator

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  25,
	}
	user3 := &User{
		name: "c",
		age:  20,
	}
	userCollection := &UserCollection{
		users: []*User{user1, user2, user3},
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
