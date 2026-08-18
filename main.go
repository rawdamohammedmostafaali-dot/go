package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID = 1

func addTask(title string) {
	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, task)
	nextID++

	fmt.Println("✅ Task added successfully!")
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found.")
		return
	}

	fmt.Println("\n========== TASKS ==========")

	for _, task := range tasks {
		status := "❌"

		if task.Completed {
			status = "✅"
		}

		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}

	fmt.Println("============================")
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("✅ Task completed!")
			return
		}
	}

	fmt.Println("❌ Task not found.")
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("🗑️ Task deleted!")
			return
		}
	}

	fmt.Println("❌ Task not found.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n====== GO TODO APP ======")
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Println("==========================")

		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("❌ Please enter a number.")
			continue
		}

		switch choice {

		case 1:
			fmt.Print("Enter task title: ")

			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			if title == "" {
				fmt.Println("❌ Task title cannot be empty.")
				continue
			}

			addTask(title)

		case 2:
			showTasks()

		case 3:
			fmt.Print("Enter task ID: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("❌ Invalid ID.")
				continue
			}

			completeTask(id)

		case 4:
			fmt.Print("Enter task ID: ")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("❌ Invalid ID.")
				continue
			}

			deleteTask(id)

		case 5:
			fmt.Println("👋 Goodbye!")
			return

		default:
			//00
			//000
			fmt.Println("❌ Invalid choice.")
		}
	}
}
