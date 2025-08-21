package main

import (
	"fmt"
	"sync"
)

// why receiver function needs to be pointer receiver in go
// from oops sense, we are directly  modifing the properties of a object so, in go we use pointer receiver or return pointer to modify the struct(properties)

type Koi struct {
	Name   string
	Origin string
}

type Shapes struct {
	Name string
}

type Pond struct {
	Name string
	Shapes ShapesList
	Fish []Koi
}

type KoiList []Koi

type ShapesList []Shapes

type Account struct {
	ID int
	balance int
}

type Transaction struct{
	accountID int
	txnType   string
	ammount   int
}



// -------- channels concepts-------------- 
func testChannelAsArgument(testChan <-chan int, num int ){

	msgValue := <-testChan
    
	fmt.Println("I got your message bro ", msgValue)
	sum := msgValue+num;
	fmt.Println("sum", sum)
}

func testChannelAsArgumentMultipleValueReceiver(values <-chan Transaction ){
    for value := range values{
       msg :=value
	   fmt.Println("Hey bro your msg is", msg.txnType, "and your id is ", msg.accountID, "and ammount ", msg.ammount)
	}
}

// sender go routine
func (acc *Account) accountOperation(wg *sync.WaitGroup,txnChan chan<- Transaction, operation string, txnAmmount int){
   defer wg.Done() 
   
   txnChan <- Transaction{
	accountID: acc.ID,
	txnType: operation,
	ammount: txnAmmount,
   }

   fmt.Println("I have sent my message", acc.ID)
}

func main() {




}


	// koiArrays := []Koi{
	// 	{Name: "blue",
	// 		Origin: "osaka",
	// 	}, {
	// 		Name:   "red",
	// 		Origin: "tokyo",
	// 	},
	// }

	// koiArrays2 := KoiList{
	// 	{Name: "red", Origin: "shifu"},
	// 	{Name: "orange", Origin: "yokhoma"},
	// }

	// pond := Pond{
	// 	Name: "Green",
	// 	Shapes: ShapesList{
	// 		{Name: "square"},
	// 		{Name: "round"},
	// 	},
	// 	Fish: []Koi{
	// 		{Name: "blue",Origin: "osaka",}, 
	// 		{Name:"red",Origin:"tokyo",},
	// 	},
	// }
     
    
	// fmt.Println(koiArrays2)
	// ---------- Pointer -----------
	// x:=2
	// px:= &x
	// fmt.Println("value", *px , "address of", px)

    
    //  channel 

	// testChan := make(chan int)

	// go testChannelAsArgument(testChan,101)

	// testChan <- 10
    // fmt.Println("All task done")
    
	// txnChan := make(chan Transaction)


	// go testChannelAsArgumentMultipleValueReceiver(txnChan)

	// go func ()  {
	// 	txnChan <- Transaction{accountID: 1, txnType: "deposit", ammount: 500}
	// }()
	// go func ()  {
	// 	txnChan <- Transaction{accountID: 3, txnType: "withdraw", ammount: 1000}
	// }()
	// go func ()  {
	// 	txnChan <- Transaction{accountID: 2, txnType: "deposit", ammount: 3000}
	// }()
  
	// time.Sleep(1* time.Second)
	// close(txnChan)
	// fmt.Println("All transaction done")

	// -------- object to channel message passing ------
    
	// txnChan := make(chan Transaction)
    // var wg sync.WaitGroup

	// user1 := Account{ID: 2, balance: 2000}
	// user2 := Account{ID: 1, balance: 100}
	// user3 := Account{ID: 3, balance: 5000}
    
	// go testChannelAsArgumentMultipleValueReceiver(txnChan)
    // wg.Add(3)
	// go user1.accountOperation(&wg,txnChan, "withdraw", 500)
	// go user2.accountOperation(&wg,txnChan, "deposit",1000)
	// go user3.accountOperation(&wg,txnChan, "withdraw",2000)

	
	// wg.Wait()
	// close(txnChan)
	
	// fmt.Println("All transaction done")

	