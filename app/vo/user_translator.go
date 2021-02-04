package vo

import "go-gin-start/ent"

/**
 * Value Object Translator
 * 关于 VO 类型数据的相关转换器，通常实现 PO2VO、VO2VO、ANY2MAP
 */

// One Po -> One Vo
func (*User) FromPo(item *ent.User) *User {
	return &User{
		UserName: item.UserName,
		Password: item.Password,
	}
}

// Some Po -> Some Vo
func (o *User) FromSomePo(items []*ent.User) []*User {

	// vars
	ranges := make([]*User, 0)

	// range
	for _, item := range items {
		ranges = append(ranges, o.FromPo(item))
	}

	return ranges
}

// One Vo -> One Vo
func (*UserOnlyUserName) FromVoUser(item *User) *UserOnlyUserName {
	return &UserOnlyUserName{
		UserName: item.UserName,
	}
}