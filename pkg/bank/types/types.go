package types

//Money представляет собой денежную единицу в минимальных единицах(копейки, дирамы)
type Money int64

//Currency представляет код валюты
type Currency string

//Код валюты
const (
	TJS Currency="TJS"
	RUB Currency="RUB"
	USD Currency="USD"
)

//PAN представляет собой номер карты
type PAN string

//Card представляет информацию о платёжной карте
type Card struct {
	ID int
	PAN PAN 
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}

type Payment struct {
	ID int
	Amount Money
}
