package repo

type cache struct{}

func InitCache() UserRepository {
	return &cache{}
}

func (c *cache) findUserById() {}
