package assembly

type Assembly struct {
	stages []Func
}

type Func func() error

func New() *Assembly {
	return &Assembly{}
}

func (a *Assembly) If(condition bool, functions ...Func) *Assembly {
	if condition {
		a.stages = append(a.stages, functions...)
	}
	return a
}

func (a *Assembly) Do() error {
	for _, f := range a.stages {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
