package cmd

import (
	"fmt"
	"gopkg.in/spf13/cobra.v0"
)

var signInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "Входит в учетную запись",
	Run: func(cmd *cobra.Command, args []string) {
		if appClient.Token != "" {
			fmt.Println("Вы уже в системе")
			return

		}

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		response, err := appClient.SignIn(username, password)

		if err != nil {
			fmt.Println(err)
			return
		}

		appClient.Token = response.Token
		fmt.Printf("Добро пожаловать, %v\n", response.Username)
	},
}

func init() {
	rootCmd.AddCommand(signInCmd)
	signInCmd.Flags().StringP("username", "u", "", "Имя пользователя")
	signInCmd.MarkFlagRequired("username")

	signInCmd.Flags().StringP("password", "p", "", "Пароль пользователя")
	signInCmd.MarkFlagRequired("password")
}
