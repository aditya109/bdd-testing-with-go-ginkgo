package main

type User struct {
    ID        uint64
    FirstName string
    LastName  string
    Age       uint16
}

type UserStorage struct {
    DB []User
}

// AddUser Add user information to the UserStorage database
func (us *UserStorage) AddUser(user User) {
    us.DB = append(us.DB, user)
}

// FindUserByID find the user by ID
func (us *UserStorage) FindUserByID(id uint64) User {
    for _, user := range us.DB {
        if user.ID == id {
            return user
        }
    }
    return User{}
}

// Count total users in the database
func (us *UserStorage) Count() int {
    return len(us.DB)
}