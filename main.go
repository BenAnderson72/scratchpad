package scratchpad

import (
	"fmt"
	"os"
)

// init with message "i am initialised"
func init() {
	fmt.Println("i am initialised")
}

// main method with message "main called". Reads & prints arguments
func main() {
	fmt.Println("main called")
	fmt.Println(os.Args)

    // vc := new(VehicleCollection);

    // vc.Populate()

}

// method to generate nth fibonacci number
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// method to generate nth fibonacci number using DP
func fibonacciDP(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// enum to describe vehicle type
type VehicleType int

const (
	Car VehicleType = iota
	Bike
	Motorcycle
	Bus
	Skateboard
)

// method to describe vehicle type
func (v VehicleType) String() string {
	switch v {
	case Car:
		return "Car"
	case Bike:
		return "Bike"
	case Motorcycle:
		return "Motorcycle"
	case Bus:
		return "Bus"
	case Skateboard:
		return "Skateboard"
	default:
		return "Unknown"
	}
}

// vehicle class with name, age, location and VehicleType
type Vehicle struct {
	Name     string
	Age      int
	Location string
	Type     VehicleType
}

// collection of vehicles
type VehicleCollection struct {
	Vehicles []Vehicle
}

// populate VehicleCollection
func (v *VehicleCollection) Populate(vehicles []Vehicle) {
	v.Vehicles = vehicles
}

// method reads vehicles.csv file
// func readVehicles() []VehicleType {
// 	file, err := os.Open("vehicles.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	var vehicles []VehicleType
// 	for {
// 		line, err := file.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			panic(err)
// 		}
// 		vehicles = append(vehicles, VehicleType(line[0]))
// 	}
// 	return vehicles
// }

// method reads vehicles.csv file and populate VehicleCollection
