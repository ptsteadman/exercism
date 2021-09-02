package account

import "sync"

// Account is a bank account with a balance and open/closed status that can safely be updated across
// multiple goroutines
type Account struct {
	balance int64
	closed  bool
	mutex   sync.Mutex
}

// Open returns a pointer to a new Account struct with an initial balance
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
		closed:  false,
	}
}

// Balance returns the current balance of an account if it is open. When calling Balance,
// the ok return param should be checked
func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed == true {
		return 0, false
	}
	return a.balance, true
}

// Close attempts to set the account to "closed" (which prevents further transactions)
// and returns the current balance as a payout amount
// the ok return param should be checked
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.closed == true {
		return 0, false
	}
	payoutAmount := a.balance
	a.balance = 0
	a.closed = true
	return payoutAmount, true
}

// Deposit adds the passed amount to the account's balance, if it is an open account.
// Negative values can be used to do a withdrawl, ok will be false if the withdrawl is more than
// the accounts current balance.
// the ok return param should be checked
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.closed == true {
		return 0, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}
