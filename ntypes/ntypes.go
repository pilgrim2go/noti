package ntypes

type Notification interface {
	Send() error
}

type Messager interface {
	Message() string
	SetMessage(string)
}
