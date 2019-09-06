// Package bank provides a concurrency-safe bank with one account
package bank

var deposits = make(chan int) //send amount to deposit
var balances = make(chan int) //receive balance
