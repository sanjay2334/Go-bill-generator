package main

import (
	"fmt"
	"sort"
	"strings"
)

func sayGreeting(n string) {
	fmt.Println("Good day! to you", n)
}

func sayBye(n string) {
	fmt.Println("Good bye!", n)
}

func Learn() {
	var nameOne string = "hello"
	var nameTwo = "Jeeva"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)
	var ageOne int8 = 20
	var ageTwo = 23
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	var s string
	var age int = 34
	var name = "sanjay"
	fmt.Print("Hello\n")
	fmt.Scan(s)
	fmt.Scanf(s)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("my age is %T and my name is %q \n", age, name)
	fmt.Printf("you scored %0.1f points \n", 255.55)

	var str = fmt.Sprintf("my age is %v and my name is %q \n", age, name)

	fmt.Println(str)

	//slices[] and arrays[3]
	var ages [3]int = [3]int{13, 25, 34}
	names := [4]string{"jeeva", "sanjay", "mario", "steve"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var score = []int{100, 200, 300}
	score[2] = 400

	fmt.Println(score, len(score))

	score = append(score, 900)

	fmt.Println(score, len(score))

	rangeOne := names[1:3]

	rangeOne = append(rangeOne, "sanju")

	fmt.Println(rangeOne)

	//strings Package
	hello := "hello"
	fmt.Println("Hello")
	fmt.Println(strings.Contains(hello, "hello"))
	fmt.Println(strings.ReplaceAll(hello, "hello", "Hi"))
	fmt.Println(hello)
	fmt.Println(strings.ToUpper(hello))
	fmt.Println(strings.Index(hello, "m"))
	fmt.Println(strings.Split(hello, ""))

	//sort Package
	agess := []int{120, 30, 40, 20, 60, 70, 80, 90}
	sort.Ints(agess)
	fmt.Println(agess)

	index := sort.SearchInts(agess, 800)
	fmt.Println(index)

	namess := []string{"sanjay", "jeeva"}
	sort.Strings(namess)
	fmt.Println(namess)

	//loops

	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i is", i)
	}

	namesss := []string{"sanjay", "jeeva", "mario"}

	for i := 0; i < len(namesss); i++ {
		fmt.Println("the Name", namesss[i])
	}

	for index, value := range namesss {
		fmt.Println("name is", value, index)
	}

	//funcss

	sayGreeting("sanjay")
	sayBye("jeeva")
	Score("jeeva")

	for _, v := range points {
		fmt.Println(v)
	}

	//maps
	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   3.99,
		"pudding": 8.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key

	phoneBook := map[int]string{
		123456789: "sanjay",
		987654321: "jeeva",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[123456789])

	phoneBook[123456789] = "Mario"

	fmt.Println(phoneBook)
}
