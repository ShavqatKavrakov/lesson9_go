package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{ Balance: 1_000_00, 
		 Active:  true ,
		 },
	}))
	fmt.Println(Total([]types.Card{
		{ Balance: 2_000_00, 
		 Active:  true ,
		 },
	}))
		fmt.Println(Total([]types.Card{
		{ Balance: -1_000_00, 
		 Active:  true ,
		 },
	}))
	fmt.Println(Total([]types.Card{
		{ Balance: 3_000_00, 
		 Active:  false ,
		 },
	}))
	
	//Output:
	//100000
	//200000
	//0
	//0
}
