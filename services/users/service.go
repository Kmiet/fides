package users

import (
	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/services/users/repo"
)

type UserService interface {
	// Create
	register() (int, error)
	// Read
	findUserWithID(int) (int, error)
	findUserWithEmail(string) (string, error)
	// Update
	softDeleteUserWithID(int) (int, error)
	// Delete
	deleteUserWithID(int) (int, error)
}

type service struct {
	cache      repo.UserRepository
	db         repo.UserRepository
	mqProducer amqp.Producer
}

func InitService(cache repo.UserRepository, db repo.UserRepository, producer amqp.Producer) UserService {
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
func (s *service) findUserWithID(id int) (int, error) {
	return id, nil
}

func (s *service) findUserWithEmail(email string) (string, error) {
	return email, nil
}

// Update
func (s *service) softDeleteUserWithID(id int) (int, error) {
	return id, nil
}

// Delete
func (s *service) deleteUserWithID(id int) (int, error) {
	return id, nil
}
