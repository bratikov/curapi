package rate

import (
	"currency/internal/scheduler"
	"fmt"
)

func Init() {
	if !scheduler.HasCron("get_rates") {
		scheduler.AddCron("get_rates", getRates, 1)
	}
}

func getRates() {
	fmt.Println("Get rates")
}
