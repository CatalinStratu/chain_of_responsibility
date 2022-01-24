package Application

type dates interface {
	Execute(*Inputs)
	SetNext(dates)
}
