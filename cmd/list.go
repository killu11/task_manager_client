package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/spf13/cobra.v0"
)

var headers = []string{"Задача", "Описание", "Статус", "Дата создания"}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Выводит список всех существующих задач",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := appClient.GetTasks()

		if err != nil {
			fmt.Println(err)
			return
		}
		if len(response) == 0 {
			fmt.Println("У вас пока нет задач!")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header(headers)
		for _, task := range response {
			if table.Append(task) != nil {
				fmt.Println("Ошибка вывода задач")
			}
		}
		if err = table.Render(); err != nil {
			fmt.Printf("list-table render error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
