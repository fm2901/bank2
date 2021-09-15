package stats

import (
	"fmt"
	"github.com/fm2901/bank/pkg/bank/types"
)


func ExampleAvg_positive() {
	fmt.Println(Avg([]types.Payment{{Amount: 10},{Amount: 20},}))
	// Output:
	// 15
}

func ExampleAvg_null() {
	fmt.Println(Avg([]types.Payment{}))
	// Output:
	// 0
}