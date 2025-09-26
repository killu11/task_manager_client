package cmd

import (
	"fmt"
	"unicode/utf8"

	"gopkg.in/spf13/cobra.v0"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "создает задачу",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("description")
		status, _ := cmd.Flags().GetInt("status")
		fmt.Printf("Заголовок задачи: %s\nЕе описание: %s\n", title, desc)
		response, err := appClient.CreateTask(title, desc, status)
		if title == "" || utf8.RuneCountInString(title) < 4 {
			fmt.Println("Заголовок не может быть пустым и должен содержать не менее 4х символов")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("title", "t", "", "Заголовок задачи")
	createCmd.MarkFlagRequired("title")

	createCmd.Flags().StringP("description", "d", "", "Описание задачи")
	createCmd.MarkFlagRequired("description")

	createCmd.Flags().IntP("status", "s", 1, "ID статуса задачи")
	createCmd.MarkFlagRequired("status")
}
