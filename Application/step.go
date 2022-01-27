package Application

type step interface {
	Execute(*Inputs) error
	SetNext(step)
}
