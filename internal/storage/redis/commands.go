package redis

type commandList struct {
	GET string
	SET string
}

var COMMANDS = &commandList{
	GET: "GET",
	SET: "SET",
}
