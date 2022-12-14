package task

type Task interface {
	setName(name string)
	getFrequency() string
	getName() string
	Execute() (s bool, e bool)
	onComplete()
	onError()
}
