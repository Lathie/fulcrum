package base

type System interface {
	Update() bool
	SendMessage(s *System, str string) bool
	RecieveMessage() bool
}
