package facade

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (c *SecurityCode) checkCode(incomingCode int) error {
	if c.code != incomingCode {
		return errors.New("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
