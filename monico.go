package monico

type Moniter struct {
	path string
}

func NewMoniter(path string) (*Moniter, error) {
	return &Moniter{}, nil
}
