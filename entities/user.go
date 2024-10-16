package entities

type User struct {
	ID    uint   // The user's unique ID (this could be set in the repository layer)
	Name  string // The user's name
	Email string // The user's unique email
	Age   int    // The user's age
}
