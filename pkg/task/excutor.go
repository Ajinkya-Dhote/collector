package task

type Executor interface {
	// SetName(name string)
	// SetFrequency(frq string)
	Register(o Task)
	DeRegister(o Task)
	Start()
	// OnComplete()
	// OnError()
}
