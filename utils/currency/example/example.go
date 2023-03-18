package main

import (
	"context"
	"fmt"

	"github.com/lucky777strike/bottest/utils/currency"
)

func main() {
	w := currency.New()
	fmt.Println(w.ParseCurrencyValue(context.Background(), "usd"))
}
