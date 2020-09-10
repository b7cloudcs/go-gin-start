package transform

import (
	"jinrongdati/app/res"
	"jinrongdati/ent"
)

func OneUser(one *ent.User) *res.User {

	return &res.User{
		UserName: "",
		Password: "",
	}
}

func SomeUser(some []*ent.User) []*res.User {

	// vars
	ranges := make([]*res.User, 0)

	// range
	for _, one := range some {
		ranges = append(ranges, OneUser(one))
	}

	return ranges
}
