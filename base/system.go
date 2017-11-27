package base

//A system must implement these methods:
//Update - For the general update loops
//SendMessage - Systems must be able to send messages to other systems
//RecieveMessage - Systems must be able ro process messages
type System interface {
	Update() bool
	RecieveMessage() bool
}
