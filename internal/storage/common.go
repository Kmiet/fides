package storage

type dataTypes struct {
	USERS string
}

const DATABASE_NAME = "fides"
const CACHE_KEY_SEPARATOR = ":"

var DATA_TYPES = &dataTypes{
	USERS: "users",
}
