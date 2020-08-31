package events

type EventName = string
type TopicName = string

type topicList struct {
	USER_SERVICE TopicName
}

var TOPICS = &topicList{
	USER_SERVICE: "user-service",
}
