package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("Fuel left:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default: 
      fuel = 0
  }
  fmt.Println("Fuel cost:", fuel)
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome, passengers, to", planet + "!")
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  // Test your functions!
  // fuelGauge(1000000)
  // calculateFuel("Mercury")
  // greetPlanet("Mercury")
  // cantFly()
  // flyToPlanet("Mars", 10000000)

  // Create `planetChoice` and `fuel`
  var fuel int = 1000000
  var planetChoice string = "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

}
