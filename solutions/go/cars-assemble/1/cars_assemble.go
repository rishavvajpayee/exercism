package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    success := successRate / 100
	return  int((float64(productionRate) * success) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const COSTGROUP uint = 95000
    const COSTINDI uint = 10000
	group := uint(carsCount / 10)
    indi := uint(carsCount % 10)

    return (group * COSTGROUP ) + (indi * COSTINDI)
    
}
