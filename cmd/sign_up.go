package cmd

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"gopkg.in/spf13/cobra.v0"
)

var PasswordRegExp = regexp.MustCompile(`^[A-Za-z0-9]{7,}$`)

var signUpCmd = &cobra.Command{
	Use:   "sign-up",
	Short: "Создает новую учетную запись",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if err := validCreds(username, password); err != nil {
			fmt.Printf("Предупреждение: %v\n", err)
			return
		}

		response, err := appClient.SignUp(username, password)

		if err != nil {
			fmt.Println(err)
			return
		}
		appClient.Token = response.Token
		fmt.Printf("Ваша учетная запись была успешно создана, %s!\n", response.Username)
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)
	signUpCmd.Flags().StringP("username", "u", "", "Имя пользователя")
	signUpCmd.MarkFlagRequired("username")

	signUpCmd.Flags().StringP("password", "p", "", "Пароль пользователя")
	signUpCmd.MarkFlagRequired("password")
}

func validCreds(username, password string) error {
	if utf8.RuneCountInString(username) < 5 {
		return fmt.Errorf("имя пользователя слишком короткое")
	}

	if !PasswordRegExp.MatchString(password) {
		return fmt.Errorf(`Неверный формат пароля.
Требования:
- Длина не менее 7 символов
- Должен содержать хотя бы одну букву (A-Z, a-z)
- Должен содержать хотя бы одну цифру (0-9)
- Разрешены только английские буквы и цифры
- Специальные символы и пробелы не допускаются
`)
	}

	return nil
}
