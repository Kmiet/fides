package users

import (
	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/services/users/repo"
)

type UserService interface {
	// Create
	register() (int, error)
	// Read
	findUserWithId(int) (int, error)
	findUserWithEmail(string) (string, error)
	// Update
	softDeleteUserWithId(int) (int, error)
	// Delete
	deleteUserWithId(int) (int, error)
}

type service struct {
	mqProducer amqp.Channel
	cache      repo.UserRepository
	db         repo.UserRepository
}

func InitService(cache repo.UserRepository, db repo.UserRepository, producer amqp.Channel) UserService {
	return &service{
		cache:      cache,
		db:         db,
		mqProducer: producer,
	}
}

// Create
func (s *service) register() (int, error) {
	return 0, nil
}

// Read
func (s *service) findUserWithId(id int) (int, error) {
	return id, nil
}

func (s *service) findUserWithEmail(email string) (string, error) {
	return email, nil
}

// Update
func (s *service) softDeleteUserWithId(id int) (int, error) {
	return id, nil
}

// Delete
func (s *service) deleteUserWithId(id int) (int, error) {
	return id, nil
}
