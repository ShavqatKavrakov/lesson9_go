package card

import "bank/pkg/bank/types"

//Total вычисляет общую сумму средств во всех картах
//Отрицательную сумму и сумму в заблокированых картах игнорируются.
func Total(cards []types.Card) types.Money {
	sum:=types.Money(0)
	for _,card:=range cards {
		if !card.Active {
			continue
		}else if card.Balance<0 {
			continue
		}
		sum+=card.Balance;
	}
  return sum	
}