package users

import (
	"github.com/Kmiet/fides/internal/net/amqp"
	"github.com/Kmiet/fides/services/users/repo/cache"
	"github.com/Kmiet/fides/services/users/repo/db"
)

type UserService interface {
	// Create
	Register(email string) (string, error)
	// Read
	FindUserWithID(id string) (interface{}, error)
	findUserWithEmail(string) (string, error)
	// Update
	softDeleteUserWithID(int) (int, error)
	// Delete
	deleteUserWithID(int) (int, error)
}

type service struct {
	cache cache.Repository
	db    db.Repository
	mq    amqp.Producer
}

func InitService(cacheRepo cache.Repository, dbRepo db.Repository, producer amqp.Producer) UserService {
	return &service{
		cache: cacheRepo,
		db:    dbRepo,
		mq:    producer,
	}
}

// Create
func (s *service) Register(email string) (string, error) {
	id, err := s.db.RegisterNewUser(email)
	if err != nil {
		return "", err
	}
	s.cache.SetUserByID(id, email)
	return id, nil
}

// Read
func (s *service) FindUserWithID(id string) (interface{}, error) {
	// if user, err := s.cache.GetUserByID(id); err == nil {
	// 	return user, err
	// }

	return s.db.FindUserByID(id)
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
