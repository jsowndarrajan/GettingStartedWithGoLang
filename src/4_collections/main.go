package main

import "fmt"

func main() {
	printDays()
	printMonths()
}

func printDays() {
	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Length:", len(days))
	fmt.Println("Capacity:", cap(days))
	for i := 0; i < len(days); i++ {
		fmt.Println(i, days[i])
	}
}

func printMonths() {
	var months []string
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "January")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "February")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "March")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "April")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "May")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "June")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "July")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "August")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "September")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "October")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "November")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
	months = append(months, "December")
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	for i, month := range months {
		fmt.Println(i, month)
	}

	// for _, month := range months {
	// 	fmt.Println(month)
	// }

	// for i := range months {
	// 	fmt.Println(months[i])
	// }
}
