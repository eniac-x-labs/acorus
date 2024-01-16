package event

type IEventProcessor interface {
	StartUnpack() error
	ClosetUnpack() error
}
