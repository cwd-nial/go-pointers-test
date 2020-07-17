package main

import (
	"fmt"
	"strconv"
	"time"
)

const Iterations = 1000000

type ItemA struct {
	Id  string
	Qty int
}

type ItemB struct {
	Id          string
	Qty         int
	naming      string
	description string
	price       float64
	createdAt   time.Time
}

func main() {

	var ItemsArrayA []ItemA
	for i := 0; i < Iterations; i++ {
		ItemsArrayA = append(ItemsArrayA, ItemA{
			Id:  "A",
			Qty: i,
		})
	}
	var ItemPointersArrayA []*ItemA
	for i := 0; i < Iterations; i++ {
		ItemPointersArrayA = append(ItemPointersArrayA, &ItemA{
			Id:  "A",
			Qty: i,
		})
	}

	var ItemsArrayB []ItemB
	for i := 0; i < Iterations; i++ {
		ItemsArrayB = append(ItemsArrayB, ItemB{
			Id:          "A",
			Qty:         i,
			naming:      "dflkhasndlfndsf",
			description: "rpijäasd fäösadjf äöoasdjkf äasdrewmöafsd äjiläawertj fdäpj asdföjasdm",
			price:       float64(i),
			createdAt:   time.Now(),
		})
	}
	var ItemPointersArrayB []*ItemB
	for i := 0; i < Iterations; i++ {
		ItemPointersArrayB = append(ItemPointersArrayB, &ItemB{
			Id:          "A",
			Qty:         i,
			naming:      "dflkhasndlfndsf",
			description: "rpijäasd fäösadjf äöoasdjkf äasdrewmöafsd äjiläawertj fdäpj asdföjasdm",
			price:       float64(i),
			createdAt:   time.Now(),
		})
	}

	fmt.Println("Case Item A:")
	fmt.Print("By value array:")
	start := time.Now()
	batchAddByValueA(ItemsArrayA)
	end := time.Now()
	caseA1 := end.Sub(start)
	fmt.Println(caseA1)

	fmt.Print("By pointer array:")
	start = time.Now()
	batchAddByPointerA(ItemPointersArrayA)
	end = time.Now()
	caseA2 := end.Sub(start)
	fmt.Println(caseA2)

	difference := ((float64(caseA1.Nanoseconds()) / float64(caseA2.Nanoseconds())) - 1) * 100
	sign := "slower"
	if difference > 0.0 {
		sign = "faster"
	}
	fmt.Println("Pointer operation is " + strconv.FormatFloat(difference, 'f', 2, 64) + "% " + sign)

	fmt.Println("Case Item B:")
	fmt.Print("By value array:")
	start = time.Now()
	batchAddByValueB(ItemsArrayB)
	end = time.Now()
	caseB1 := end.Sub(start)
	fmt.Println(caseB1)

	fmt.Print("By pointer array:")
	start = time.Now()
	batchAddByPointerB(ItemPointersArrayB)
	end = time.Now()
	caseB2 := end.Sub(start)
	fmt.Println(caseB2)

	difference = ((float64(caseB1.Nanoseconds()) / float64(caseB2.Nanoseconds())) - 1) * 100
	sign = "slower"
	if difference > 0.0 {
		sign = "faster"
	}
	fmt.Println("Pointer operation is " + strconv.FormatFloat(difference, 'f', 2, 64) + "% " + sign)
}

func batchAddByValueA(items []ItemA) {
	for k, v := range items {
		v.Id = strconv.Itoa(k)
	}
	return
}

func batchAddByPointerA(items []*ItemA) {
	for k, v := range items {
		v.Id = strconv.Itoa(k)
	}
	return
}

func batchAddByValueB(items []ItemB) {
	for k, v := range items {
		v.Id = strconv.Itoa(k)
	}
	return
}

func batchAddByPointerB(items []*ItemB) {
	for k, v := range items {
		v.Id = strconv.Itoa(k)
	}
	return
}
