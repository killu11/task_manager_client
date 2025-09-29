package cmd

import (
	"fmt"

	"gopkg.in/spf13/cobra.v0"
)

var updateStatusCmd = &cobra.Command{
	Use:   "update-status",
	Short: "Обновляет статус задачи",
	Long:  "Обновляет статус задачи по её заголовку и номеру статуса",
	Run: func(cmd *cobra.Command, args []string) {
		if !appClient.TokenExist() {
			fmt.Println("Необходимо войти или создать учетную запись!")
			return
		}
		title, _ := cmd.Flags().GetString("title")
		statusID, _ := cmd.Flags().GetInt("status")

		if title == "" {
			fmt.Println("Заголовок задачи не может быть пустым!")
			return
		}
		if statusID > 3 || statusID < 1 {
			fmt.Println("Значение статуса может быть в пределах от 1 до 3")
			return
		}
		respString, err := appClient.UpdateStatusByTitle(title, statusID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(respString)
	},
}

func init() {
	rootCmd.AddCommand(updateStatusCmd)
	updateStatusCmd.Flags().StringP("title", "t", "", "Используется для заголовка задачи")
	updateStatusCmd.MarkFlagRequired("title")

	updateStatusCmd.Flags().IntP("status", "s", 1, "Используется для номера статуса задачи")
	updateStatusCmd.MarkFlagRequired("status")
}
