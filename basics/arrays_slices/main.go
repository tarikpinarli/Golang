package main

import (
	"fmt"
)

func main() {
	todoList := [2]string{"Workout", "Breakfast"}
	dynamicList := []string{"Workout", "Breakfast"}
	//todoList[2] = "new"
	dynamicList = append(dynamicList, "LAST")
	fmt.Println("len of todolist", len(todoList))
	fmt.Println("len of dynamicList", len(dynamicList))
}