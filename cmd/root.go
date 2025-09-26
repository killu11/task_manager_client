package cmd

import (
	"bufio"
	"client/internal/client"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/shlex"
	"gopkg.in/spf13/cobra.v0"
)

var appClient = client.NewClient()

var rootCmd = &cobra.Command{
	Use:   "tmr",
	Short: "Запускает менеджер задач",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		startInteractiveSession(cmd)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func startInteractiveSession(cmd *cobra.Command) {
	fmt.Println("Добро пожаловать в интерактивный режим менеджера задач!")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("task-manager>")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Ошибка чтения команды: ", err.Error())
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		executeCobraCommand(cmd, input)
	}
}

// переделать метод, работает неправильно)
func executeCobraCommand(rootCmd *cobra.Command, input string) {
	args := splitArgs(input)
	if len(args) == 0 {
		return
	}

	if args[0] == rootCmd.Name() {
		fmt.Println("Используйте дочерние команды: create")
		return
	}

	for _, cmd := range rootCmd.Commands() {
		if args[0] == cmd.Name() {
			// ✅ ПРАВИЛЬНО: парсим оригинальные аргументы
			//if err := cmd.ParseFlags(args[1:]); err != nil {
			//	fmt.Printf("Ошибка парсинга флагов: %v\n", err)
			//	return
			//}
			cmd.SetArgs(args)

			// ✅ ПРАВИЛЬНО: передаем оригинальные аргументы
			if cmd.RunE != nil {
				if err := cmd.RunE(cmd, args[1:]); err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				}
			} else if cmd.Run != nil {
				cmd.Run(cmd, args[1:])
			} else {
				fmt.Println("Команда не имеет обработчика")
			}
			return
		}
	}

	fmt.Printf("Неизвестная команда: %s\n", args[0])
}

func splitArgs(args string) []string {
	argsSlice, err := shlex.Split(args)
	if err != nil {
		log.Println(err)
		return nil
	}
	return argsSlice
}
