package application

type step interface {
	Execute(*Inputs) error
	SetNext(step)
}
