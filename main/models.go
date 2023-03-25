/*
Models to represent Receipt and individual Item in Receipt
*/
package main

/*
Description: Represents Item provided in Receipt as list
*/
type Item struct {
	ShortDescription string
	Price            string
}

//Represents Receipt struct provided in requestbody
type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}
