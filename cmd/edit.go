package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

func EditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit <name>",
		Short: "Edit a task using the default text editor",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "vi" // Default to vi if EDITOR env is not set
			}
			tmpFile := fmt.Sprintf("/tmp/%s_task_edit", name)
			f, err := os.Create(tmpFile)
			if err != nil {
				fmt.Println("Error creating temp file:", err)
				return
			}
			defer os.Remove(tmpFile)
			defer f.Close()

			cmd := exec.Command(editor, tmpFile)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Error opening editor:", err)
			}
			fmt.Println("Editing completed for task:", name)
		},
	}
}