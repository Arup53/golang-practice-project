package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	accountID int
	balance int
	mutex   sync.Mutex
}

type Transaction struct {
    accountID int
	txnType  string
    amount    int
	resultChan chan string
}

func processTransaction(trxChan <-chan Transaction, accounts []BankAccount) {
	for trx := range trxChan { // Loop until channel is closed
		// Find the target account
		var account *BankAccount
		for i := range accounts {
			if accounts[i].accountID == trx.accountID {
				account = &accounts[i]
				break
			}
		}

		if account == nil {
			trx.resultChan <- "Account not found"
			continue
		}

		// Lock the account mutex and defer unlock
		account.mutex.Lock()
		func() {
			defer account.mutex.Unlock() // Ensures mutex is released even on early return

			switch trx.txnType {
			case "deposit":
				account.balance += trx.amount
				trx.resultChan <- "Deposit successful, new balance: " + fmt.Sprint(account.balance)
			case "withdraw":
				if account.balance >= trx.amount {
					account.balance -= trx.amount
					trx.resultChan <- "Withdrawal successful, new balance: " + fmt.Sprint(account.balance)
				} else {
					trx.resultChan <- "Withdrawal failed: insufficient funds"
				}
			default:
				trx.resultChan <- "Unknown transaction type"
			}
		}()
	}
}





func main() {
	trxChan := make(chan Transaction)
    
	accountSlice := []BankAccount{
		{accountID: 1, balance: 5000}, 
		{accountID: 2, balance: 3000},
		{accountID: 3, balance: 7000},
	}

	
	
}

