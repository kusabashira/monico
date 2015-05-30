package monico

import (
	"os"
)

type Moniter struct {
	path string
}

func NewMoniter(path string) (*Moniter, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &Moniter{
		path: path,
	}, nil
}

func NewMoniterWithWD() (*Moniter, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return NewMoniter(wd)
}

func (m *Moniter) Path() string {
	return m.path
}
