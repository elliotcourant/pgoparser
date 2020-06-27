package parser

import (
	"github.com/pkg/errors"
)

func (p *parser) expected(expected, found interface{}) error {
	return errors.Errorf("expected %v found %v - [%d]", expected, found, p.index)
}
