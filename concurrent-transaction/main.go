package main

import "sync"

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




func main() {
	trxChan := make(chan Transaction)
    
	accountSlice := []BankAccount{
		{accountID: 1, balance: 5000}, 
		{accountID: 2, balance: 3000},
		{accountID: 3, balance: 7000},
	}

	
	
}

