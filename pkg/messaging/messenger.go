package messaging


type Messenger interface {
	Write(message string)
	Subscribe(onEvent func(message string))
}
