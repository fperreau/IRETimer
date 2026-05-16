package main

import (
	"os"
	"fmt"
	"IRETimer/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		return
	}

	switch os.Args[1] {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task init <sqlite_path>")
			return
		}
		cmd.InitDatabase(os.Args[2])
	case "list":
		cmd.ListTasks()
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: task add <name> --cron <cron> --script <script>")
			return
		}
		cmd.AddTask(os.Args[2], os.Args[4], os.Args[6])
	case "edit":
		cmd.EditTask(os.Args[2])
	case "enable":
		cmd.EnableTask(os.Args[2])
	case "disable":
		cmd.DisableTask(os.Args[2])
	case "del":
		cmd.DeleteTask(os.Args[2])
	case "log":
		cmd.ShowLogs()
	default:
		fmt.Println("Unknown command")
	}
}