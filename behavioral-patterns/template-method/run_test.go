package template_method

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	smsOTP := &Sms{}
	o := Otp{iOtp: smsOTP}
	o.genAndSendOTP(4)

	fmt.Println()

	emailOTP := &Email{}
	o = Otp{iOtp: emailOTP}
	o.genAndSendOTP(4)
}
