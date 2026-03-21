package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

const (
	format    = "2006-01-02"
	tasksFile = "tasks.json"
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
		addTask(os.Args)
	case "list":
		if len(os.Args) == 2 {
			listTasks("")
			return
		}
		listTasks(os.Args[2])
	case "delete":
		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		deleteTask(arg)
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		updateTask(os.Args, id)
	case "mark":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error reading third argument: ", err)
			return
		}
		updateStatus(os.Args, id)
	case "help":
		helper()
	default:
		fmt.Print("Invalid command!")
		return
	}
}

func addTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Not enough arguments")
		return
	}

	now := time.Now().Format(format)

	tasks := loadTasks()
	newTask := Task{
		ID:          getNextID(tasks),
		Description: args[2],
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)

	err := saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:  ", err)
		return
	}

	fmt.Println("Task added!♥")
	fmt.Printf("%s (ID: %d)\n", args[2], newTask.ID)
}

func loadTasks() []Task {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error task file is not exist: ", err)
			return []Task{}
		}
		fmt.Println("Error reading tasks file: ", err)
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error JSON parsing: ", err)
		return []Task{}
	}

	return tasks
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error formating JSON: ", err)
		return err
	}
	err = os.WriteFile(tasksFile, data, 0644)
	if err != nil {
		fmt.Println("Error: failed to save task", err)
		return err
	}
	return nil
}

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func listTasks(status string) {
	valid := map[string]bool{
		"todo":        true,
		"done":        true,
		"in-progress": true,
	}

	if status != "" && !valid[status] {
		fmt.Println("Unknown status!")
		return
	}

	found := false
	tasks := loadTasks()

	for _, t := range tasks {
		if status == "" || t.Status == status {
			fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			found = true
		}
	}

	if !found {
		fmt.Println("No tasks found")
	}
}

func deleteTask(id int) error {
	tasks := loadTasks()
	for i, t := range tasks {
		if t.ID == id {
			newTaskList := append(tasks[:i], tasks[i+1:]...)
			err := saveTasks(newTaskList)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			fmt.Printf("Task Deleted! [ID: %d]\n", t.ID)
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}

func updateStatus(args []string, id int) error {
	tasks := loadTasks()

	if len(args) < 2 {
		return fmt.Errorf("To update your task's status use format: 'task-cli <mark> [id]'")
	}

	for i, t := range tasks {
		if t.ID == id {
			switch args[2] {
			case "done":
				tasks[i].Status = "done"
			case "in-progress":
				tasks[i].Status = "in-progress"
			default:
				return fmt.Errorf("Unknown status")
			}
			tasks[i].UpdatedAt = time.Now().Format(format)
			err := saveTasks(tasks)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			fmt.Println("Task status updated!")
			return nil
		}
	}
	return fmt.Errorf("ID not found")
}

func updateTask(args []string, id int) error {
	tasks := loadTasks()

	if len(args) < 4 {
		return fmt.Errorf("To update your task's status use format: 'task-cli update [id] \"Your changed task\"'")
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = args[3]
			err := saveTasks(tasks)
			if err != nil {
				return fmt.Errorf("Error saving task list: %v", err)
			}
			tasks[i].UpdatedAt = time.Now().Format(format)
			fmt.Println("Your task successfully updated!")
			return nil
		}
	}
	return fmt.Errorf("ID not found")
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
