package redis

type commandList struct {
	GET  string
	SET  string
	PING string
}

var COMMANDS = &commandList{
	GET:  "GET",
	SET:  "SET",
	PING: "PING",
}
