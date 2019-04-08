package user

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var userStorage []User = []User{
	{Name: "Orlando", Age: 27},
	{Name: "Daniel", Age: 15},
	{Name: "Monica", Age: 22},
	{Name: "Karl", Age: 34},
	{Name: "Diana", Age: 12},
}

func GetAll() []User {
	return userStorage
}

func GetOne(id int) User {
	for i, u := range userStorage {
		if id == i {
			return u
		}
	}
	return User{}
}

func Create(u ...User) {
	userStorage = append(userStorage, u...)
}
