package option

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	user := NewUser(
		WithName("zhangsan"),
		WithAge(18),
	)
	fmt.Println(user)
}

func TestWhere(t *testing.T) {
	users := []*User{
		{
			Name: "zhangsan",
			Age:  12,
		},
		{
			Name: "zhangsan",
			Age:  13,
		},
		{
			Name: "zhangsan",
			Age:  18,
		},
		{
			Name: "zhangsan",
			Age:  19,
		},
	}
	result := where{
		Name:    "zhangsan",
		AgeFrom: 13,
		AgeEnd:  18,
	}.Apply(users)
	for _, r := range result {
		fmt.Println(r)
	}

}

func TestLimit(t *testing.T) {
	users := []*User{
		{
			Name: "zhangsan",
			Age:  12,
		},
		{
			Name: "zhangsan",
			Age:  13,
		},
		{
			Name: "zhangsan",
			Age:  18,
		},
		{
			Name: "zhangsan",
			Age:  19,
		},
	}
	result := where{
		Name:    "zhangsan",
		AgeFrom: 12,
		AgeEnd:  19,
	}.Apply(users)

	result = Limit{
		Offset: 1,
		Count:  2,
	}.Apply(result)

	for _, r := range result {
		fmt.Println(r)
	}
}

func TestQuery(t *testing.T) {
	users := []*User{
		{
			Name: "zhangsan",
			Age:  12,
		},
		{
			Name: "zhangsan",
			Age:  13,
		},
		{
			Name: "zhangsan",
			Age:  18,
		},
		{
			Name: "zhangsan",
			Age:  19,
		},
	}
	q1 := where{
		Name:    "zhangsan",
		AgeFrom: 12,
		AgeEnd:  19,
	}

	q2 := Limit{
		Offset: 1,
		Count:  2,
	}
	result := QuryUsers(users, q1, q2)

	for _, r := range result {
		fmt.Println(r)
	}
}
