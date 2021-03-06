package types

//Money представляет собой денежную сумму в минимальных единицах{центы, копейки, дирамы и т.д}
type Money int64

//Currency представляет код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет информатцию о платёжной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
	Number   string
}
type Payment struct {
	ID     int
	Amount Money
}
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money //баланс в дирамах
	Active  bool
}