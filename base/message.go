package base

//Messages contain 3 things
//Sys is a string describing where the message came from
//Message is a string with the message that the current system must read
//Code is specific codes that register behaviors
type Message struct {
	Sys     string
	Message string
	Code    int
}
