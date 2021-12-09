package types

type Money int64
type Currency string
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)
type PAN string
type Card struct {
	ID uint
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}
type Category string
type Status string
const (
	StatusOk Status="OK"
	StatusFail Status="FAIL"
	StatusInProgress Status="INPROGRESS"
)
type Payment struct {
	ID uint
	Amount Money
	Category Category
	Status Status
}
type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}