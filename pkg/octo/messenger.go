package octo

type Messenger interface {
	Write(message string)
	WriteAndListen(message string, onEvent func(message string))
	Subscribe(onEvent func(message string))
	Unsubscribe()
}
