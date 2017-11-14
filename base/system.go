package base

type System interface {
	Update()
	SendMessage(s *System, str string)
	RecieveMessage()
}
