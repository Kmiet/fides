package repo

type UserRepository interface {
	findUserByID(id string)
}
