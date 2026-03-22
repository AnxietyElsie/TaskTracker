package main

import (
	"TaskTracker/internal/service"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("TaskTracker CLI v0.1")
	fmt.Println("Welcome to TaskTracker! What do u want to do today?")
	fmt.Println("use    task-cli help     to get the information about commands")

	if os.Args[0] != "task-cli" {
		fmt.Println("Error: invalid command. Start your commands with 'task-cli'!")
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: invalid command")
		return
	}

	switch os.Args[1] {
	case "add":
		service.AddTask(os.Args)
	case "list":
		if len(os.Args) == 2 {
			service.ListTasks("")
			return
		}
		service.ListTasks(os.Args[2])
	case "delete":
		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		service.DeleteTask(arg)
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		service.UpdateTask(os.Args, id)
	case "mark":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		service.UpdateStatus(os.Args, id)
	case "help":
		helper()
	default:
		fmt.Print("Invalid command!")
		return
	}
}

func helper() {
	fmt.Println("What do u wanna know?")

	fmt.Println("U should start all your commands with: task-cli")
	fmt.Println("COMMAND LIST:")
	fmt.Println("task-cli add \"Your task\"  -add a new task")
	fmt.Println("task-cli delete *index*  -delete a task with this index; enter the index without *")
	fmt.Println("task-cli list  -show all tasks")
	fmt.Println("task-cli list *status*  -show tasks with this status; enter a status without *: todo/done/in-progress")
	fmt.Println("task-cli mark *status* [id]  -mark a tsk with status; enter id without [] and one status without *: todo/done/in-progress")
	fmt.Println("task-cli update [id] \"Your updated task\"  -update a description of the task; enter id without []")
}
