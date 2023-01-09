package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip):", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Item Price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be an number")
			promptOptions(b)
		}
		b.addItems(name, p)
		fmt.Println("item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter an tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be an number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("the entered tip: ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file", b.name)
	default:
		fmt.Println("That was not an valid option....")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
