package events

type userServiceEvents struct {
	USER_DELETED EventName
}

var USER_SERVICE = &userServiceEvents{
	USER_DELETED: "USER_DELETED",
}
