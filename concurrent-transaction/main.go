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

// func processTransaction(trxChan <-chan Transaction, accounts []BankAccount) {
// 	for trx := range trxChan { // Loop until channel is closed
// 		// Find the target account
// 		var account *BankAccount
// 		for i := range accounts {
// 			if accounts[i].accountID == trx.accountID {
// 				account = &accounts[i]
// 				break
// 			}
// 		}

// 		if account == nil {
// 			trx.resultChan <- "Account not found"
// 			continue
// 		}

// 		// Lock the account mutex and defer unlock
// 		account.mutex.Lock()
// 		func() {
// 			defer account.mutex.Unlock() // Ensures mutex is released even on early return

// 			switch trx.txnType {
// 			case "deposit":
// 				account.balance += trx.amount
// 				trx.resultChan <- "Deposit successful, new balance: " + fmt.Sprint(account.balance)
// 			case "withdraw":
// 				if account.balance >= trx.amount {
// 					account.balance -= trx.amount
// 					trx.resultChan <- "Withdrawal successful, new balance: " + fmt.Sprint(account.balance)
// 				} else {
// 					trx.resultChan <- "Withdrawal failed: insufficient funds"
// 				}
// 			default:
// 				trx.resultChan <- "Unknown transaction type"
// 			}
// 		}()
// 	}
// }


// func preocessTransaction(trnChan <-chan Transaction, accounts map[int]*BankAccount){

// 	for trx := range trnChan {
//         account, exist := accounts[trx.accountID]
		
// 		if !exist{
// 			trx.resultChan <-"Account not found"
// 			continue
// 		}

// 		account.mutex.Lock()
// 		func ()  {
// 		defer account.mutex.Unlock()
		
// 		switch trx.txnType{
// 		case "deposit":
// 				account.balance += trx.amount
// 				trx.resultChan<-fmt.Sprintf(
//                     "Deposit successful for, new balance: %d", account.balance,
//                 )
// 				close(trx.resultChan)
			
// 			case "withdraw":
// 				if trx.amount <= account.balance {
//                       account.balance -= trx.amount
// 					  trx.resultChan<-fmt.Sprintf(
//                     "Withdraw successful, new balance: %d", account.balance,
//                 )
// 				} else {
// 					trx.resultChan <-"Withdrawal failed: insufficient funds"
// 				}
// 				 close(trx.resultChan)
// 			default:
// 				 trx.resultChan <- "Unknown transaction type"
// 				 close(trx.resultChan)

// 		}
// 		}()
// 	}
// }

// ------------ processtransaction sends multiple values in channel for each transcation------------ 
func processTransaction(trxChan <-chan Transaction, accounts map[int]*BankAccount) {
	for trx := range trxChan {
		account, exists := accounts[trx.accountID]
		if !exists {
			trx.resultChan <- "Account not found"
			close(trx.resultChan) // close after all updates
			continue
		}

		account.mutex.Lock()
		go func(trx Transaction, account *BankAccount) {
			defer account.mutex.Unlock()
			defer close(trx.resultChan) // close when transaction fully done

			trx.resultChan <- fmt.Sprintf("Starting transaction... %d", account.accountID)

			switch trx.txnType {
			case "deposit":
				trx.resultChan <- fmt.Sprintf("Depositing %d", trx.amount)
				account.balance += trx.amount
				trx.resultChan <- fmt.Sprintf("Deposit successful, new balance: %d", account.balance)

			case "withdraw":
				trx.resultChan <- fmt.Sprintf("Attempting to withdraw %d", trx.amount)
				if account.balance >= trx.amount {
					account.balance -= trx.amount
					trx.resultChan <- fmt.Sprintf("Withdrawal successful, new balance: %d", account.balance)
				} else {
					trx.resultChan <- "Withdrawal failed: insufficient funds"
				}

			default:
				trx.resultChan <- "Unknown transaction type"
			}
		}(trx, account)
	}
}



func fanInResults(cs ...chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))
    
	output:= func (c <-chan string)  {
		for msg := range c{
			out <-msg
		}
		wg.Done()
	}
    // start separate go routines for each input channels
	for _,ch:= range cs{
		go output(ch)
	}

	go func ()  {
	 wg.Wait()
	 close(out)	
	}()
    
    return out
}


func main() {
	trxChan := make(chan Transaction)
    
	// accountSlice := []BankAccount{
	// 	{accountID: 1, balance: 5000}, 
	// 	{accountID: 2, balance: 3000},
	// 	{accountID: 3, balance: 7000},
	// }

	accounts := map[int]*BankAccount{
		1:{accountID: 1, balance: 5000}, 
		2:{accountID: 2,balance: 300},
		3:{accountID: 3, balance: 7000},
	}

	go processTransaction(trxChan, accounts)
    
	// ------------------ Blocking operation i.e. although processTransaction runs spearatly main go routine prints result one after another not true concurrency------------- 
	// result1:= make(chan string)
	// trxChan <- Transaction{accountID: 1, txnType: "deposit", amount: 1000, resultChan: result1}
	// fmt.Println(<-result1)
      
    // result3:=make(chan string)
	// trxChan <-Transaction{accountID: 3, txnType: "withdraw", amount: 10000, resultChan: result3}
	// fmt.Println(<-result3)

	// result2:= make(chan string)
	// trxChan<-Transaction{accountID: 2,txnType: "withdraw", amount: 100, resultChan: result2}
    // fmt.Println(<-result2)

	// --------------------- Fan in pattern for printing results randomaly or concurrently not sequenctial blocking

	result1 := make(chan string)
	result2 := make(chan string)
	result3 := make(chan string)

	trxChan <- Transaction{accountID: 1, txnType: "deposit", amount: 1000, resultChan: result1}
	trxChan <- Transaction{accountID: 3, txnType: "withdraw", amount: 10000, resultChan: result3}
	trxChan <- Transaction{accountID: 2, txnType: "withdraw", amount: 100, resultChan: result2}

	results := fanInResults(result1, result2, result3)

	for res:= range results{
		fmt.Println(res)
	}
	

	close(trxChan)
}

