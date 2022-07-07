package main
import "fmt"

type Token interface{
	addToken() int 
}
type wallet struct {
	balance int 
}
func (w wallet) addToken() int {
	return w.balance + 1
}
func NewBalance(w Token) {
	fmt.Println("New balance is",w.addToken())
}
func main() {
	w:= wallet{balance: 0}
	NewBalance(w)
}
