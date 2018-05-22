package cli

// Interface that must be implemented by actions. This is interface is used by the RegisterAction function. The Run
// method of the implementing type will be called if the given route matches it.
type Runner interface {
	Run() error
}

type RunFunc func() error

func (rf RunFunc) Run() error {
	return rf()
}