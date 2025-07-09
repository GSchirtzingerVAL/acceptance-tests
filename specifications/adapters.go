package specifications

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type MeanGreeterAdapter func(name string) string

func (m MeanGreeterAdapter) Curse(name string) (string, error) {
	return m(name), nil
}
