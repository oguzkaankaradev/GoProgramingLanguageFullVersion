package Variable

import (
	"fmt"
	"math"
)

func investimentCalculator(investmentAmount int, expectedReturnRate float64, years int) {
	// investmentAmount = 1000
	// expectedReturnRate = 5.5
	// years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
