package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "tasks.txt"

// Add task
func addTask(task string) {
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(task + "\n")
	fmt.Println("Task added!")
}

// View tasks
func viewTasks() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No tasks found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}
}

// Delete all tasks
func clearTasks() {
	os.Remove(fileName)
	fmt.Println("All tasks cleared!")
}

func main() {
	var choice int
	var task string

	fmt.Println("==== GO CLI TODO APP ====")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Clear All Tasks")
	fmt.Print("Enter choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter task: ")
		reader := bufio.NewReader(os.Stdin)
		task, _ = reader.ReadString('\n')
		addTask(task)
	case 2:
		viewTasks()
	case 3:
		clearTasks()
	default:
		fmt.Println("Invalid choice")
	}
}