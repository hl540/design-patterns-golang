package facade

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 123456)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 123456, 10)
	if err != nil {
		t.Fatalf("Error:%s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 123456, 5)
	if err != nil {
		t.Fatalf("Error:%s\n", err.Error())
	}
}
