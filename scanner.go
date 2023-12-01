package iter

import (
	"bufio"
	"io"
	"os"
)

type Scanner struct {
	*bufio.Scanner
}

func (s *Scanner) Iterate(yield func(string) error) error {
	for s.Scan() {
		if s.Err() != nil {
			return s.Err()
		}

		if e := yield(s.Text()); e != nil {
			return e
		}
	}

	return nil
}

func NewScanner(r io.Reader) *Scanner {
	s := bufio.NewScanner(r)
	s.Buffer([]byte{}, 1e12)
	return &Scanner{s}
}

type Path string

func (p Path) Iterate(yield func(string) error) (err error) {
	r, e := os.Open(string(p))
	if e != nil {
		return e
	}
	defer func() {
		e := r.Close()
		if err == nil {
			err = e
		}
	}()

	s := NewScanner(r)
	return s.Iterate(yield)
}
