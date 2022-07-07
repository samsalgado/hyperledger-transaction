package main
import "fmt"
type transfer interface {
	sendToken() int
}
type accountInGreen interface {
	hasTokens() bool 

}
type wallet struct {
	inGreen bool
	balance int
}
func (b wallet) hasTokens() bool {
	return b.inGreen
}
func (b wallet) sendToken() int {
	return b.balance - 1 
}
func Transfer(b transfer) {
	if balanceNotZero, ok := b.(accountInGreen); ok {
		fmt.Println(balanceNotZero.hasTokens())
		fmt.Println("Decremented balance is", b.sendToken())
	}
}
func main() {
	b:= wallet{inGreen: true, balance:32}
	Transfer(b)
}