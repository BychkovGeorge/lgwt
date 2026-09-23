package pointersAndErrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Deposit(20)
		got := wallet.Balance()
		want := Bitcoin(30.00)
		if got != want {
			t.Errorf("Balance is %v, but has to be %v", got, want)
		}
	})

	t.Run("valid withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(300)}
		err := wallet.Withdraw(Bitcoin(100))
		got := wallet.Balance()
		want := Bitcoin(200.00)
		if err != nil {
			t.Errorf("Wrong balance calculation")
		}
		if got != want {
			t.Errorf("Balance is %v, but has to be %v", got, want)
		}
	})

	t.Run("invalid withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(300)}
		err := wallet.Withdraw(Bitcoin(400))
		if err == nil {
			t.Errorf("Wrong balance calculation")
		}
	})
}

func TestBitcoin(t *testing.T) {
	t.Run("Test bitcoin", func(t *testing.T) {
		bitcoin := Bitcoin(40)
		got := bitcoin.String()
		want := "40.00 BTC"
		if got != want {
			t.Errorf("Bitcoin value has to be %v, but is %v", want, got)
		}
	})
}
