package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	errNoData  = errors.New("no data")
	errTimeout = errors.New("timeout")
)

func sendData(numbers chan int) {
	num := []int{1, 2, 3, 4, 5, 1, 2, 3, 1}
	for _, v := range num {
		numbers <- v
	}
	close(numbers)
}
func main() {
	num := make(chan int, 1)
	go sendData(num)
	mostCommon, err := minFromChann(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mostCommon)
	}
}
func minFromChann(numbers chan int) (int, error) {

	var nilSlace []int
	var min int
	/*
		for i := range numbers {
			fmt.Println(i)
		}
	*/
L:
	for {
		select {
		case <-time.After(1 * time.Second):
			return 0, errTimeout
		case rez, ok := <-numbers:
			if ok {
				nilSlace = append(nilSlace, rez)
				fmt.Println(rez)
			} else {
				if len(nilSlace) == 0 {
					return 0, errNoData
				} else {
					fmt.Println("kanal zatvoren")
					break L
				}
			}
		}

	}
	for index, value := range nilSlace {
		fmt.Println(index, value)
		if index == 0 {
			min = value
		}
		if value < min {
			min = value
		}
	}
	return min * 707, nil
}
