package main

import (
    "fmt"
    "time"
)

func main() {
    array := []int{}
    var i int
    for i != -1 {
        fmt.Println("Add an integer to the array. Type -1 to quit.")
        fmt.Scan(&i)
    
        if i != -1 {
            array = append(array, i)
        }
    }

    fmt.Println("Enter a target value.")
    var target int
    fmt.Scan(&target)

    fmt.Printf("Array: %v\nTarget: %d\n", array, target)

    if binarySearch(array, target) != -1 {
        fmt.Printf("Target value %d was found at index %d.\n", target, binarySearch(array, target))
    } else {
        fmt.Printf("The target value %d was not found in the array %v.\n", target, array)
    }
}

func binarySearch(array []int, target int) int {
    startTime := time.Now()
    
    l := 0
    r := len(array) - 1

    for l <= r {
        m := (l + r) / 2

        if array[m] > target {
            r = m - 1
        } else if array[m] < target {
            l = m + 1
        } else {
            return m
        }
    }
    
    elapsedTime := time.Since(startTime)
    fmt.Printf("Time taken to execute binary search: %s\n", elapsedTime)

    return -1
}
