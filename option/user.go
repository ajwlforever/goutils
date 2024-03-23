package option

/*
	Option模式没有固定的应用场景，通常是作为不定长参数(所以才叫option)。
	通过type把option定义为函数或接口。
*/

type User struct {
	Name string
	Age  int
}

type UserOption func(*User) // optional

func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) UserOption {
	return func(u *User) {
		u.Age = age
	}
}

func NewUser(options ...UserOption) *User {
	user := new(User)
	for _, option := range options {
		option(user)
	}
	return user
}

// 复杂的 Option 作为interface
type QueryOption interface {
	Apply([]*User) []*User
}

// where where query option
type where struct {
	Name    string
	AgeFrom int
	AgeEnd  int
}

func (w where) Apply(users []*User) []*User {
	var result []*User
	for _, user := range users {
		if user.Name == w.Name && user.Age >= w.AgeFrom && user.Age <= w.AgeEnd {
			result = append(result, user)
		}
	}
	return result
}

type Limit struct {
	Offset int
	Count  int
}

func (l Limit) Apply(users []*User) []*User {
	// check if limit
	if l.Count == 0 {
		return users
	}
	// check if offset
	if l.Offset >= len(users) {
		return nil
	}
	if l.Offset+l.Count > len(users) {
		return users[l.Offset:]
	}
	return users[l.Offset : l.Offset+l.Count]
}

func QuryUsers(users []*User, options ...QueryOption) []*User {
	res := users
	for _, option := range options {
		res = option.Apply(res)
	}
	return res

}
