package vo

/**
 * Value Object
 * 这是一个 VO 类型数据，即 Value Object，系统中的价值数据或值类型数据
 * 使用 VO 的目的在于传递和转换一些非 PO（Persistent Object）持久化对象中的数据结构
 * 通常我们需要将 PO 类型数据，即 ent 数据转换为 VO 类型数据传输给其他应用层或前端
 */

// User
type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// UserOnlyUserName
type UserOnlyUserName struct {
	UserName string `json:"user_name"`
}
