package repository

import (
	"go-gin-start/app/ent"
	"go-gin-start/app/ent/user"
	"go-gin-start/app/util"
)

type User struct{}

/**
 * Get One Demo
**/
func (User) GetOne(id int) (*ent.User, error) {

	// get
	entUser, err := util.DBC.User.
		Query().
		Where(user.IDEQ(id)).
		First(util.Ctx)

	// err
	if err != nil {
		return nil, err
	}

	return entUser, nil
}