package facade

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account:      newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (f *WalletFacade) addMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := f.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = f.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	f.wallet.creditBalance(amount)
	f.notification.sendWalletCreditNotification()
	f.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (f *WalletFacade) deductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := f.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = f.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = f.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	f.notification.sendWalletDebitNotification()
	f.ledger.makeEntry(accountId, "credit", amount)
	return nil
}
