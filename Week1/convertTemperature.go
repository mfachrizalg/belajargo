package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in Celcius : ")
	var celcius float32
	fmt.Scan(&celcius)
	for {
		fmt.Println("Enter 1 to convert to Fahrenheit")
		fmt.Println("Enter 2 to convert to Kelvin")
		fmt.Println("Enter 3 to convert to Reamur")
		var choice uint
		fmt.Print("Enter your choice : ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fahrenheit := (celcius * 9 / 5) + 32
			fmt.Println("Temperature in Fahrenheit : ", fahrenheit, "°F")
			break
		case 2:
			kelvin := celcius + 273
			fmt.Println("Temperature in Kelvin : ", kelvin, "°K")
			break
		case 3:
			reamur := celcius * 4 / 5
			fmt.Println("Temperature in Reamur : ", reamur, "°R")
			break
		default:
			fmt.Println("Invalid choice")
			break
		}
		fmt.Print("Do you want to convert again? (y/n) ")
		var option string
		fmt.Scan(&option)
		if option == "n" || option == "N" {
			fmt.Print("Goodbye")
			break
		}
	}
}
