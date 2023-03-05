package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tree "github.com/YoungPentagonHacker/consolephonebook/binarytree"
	db "github.com/YoungPentagonHacker/consolephonebook/database"
)

func main() {
	var action string
	var name string
	var numbers []string
	scanner := bufio.NewScanner(os.Stdin)
	var root tree.Node
	for {
		root = tree.CreateTree(db.GetUsers())
		fmt.Print("\nДля просмотра телефонной книги нажмите 1\nДля добавления новой записи в книгу нажмите 2\nДля удаления номера из телефонной книги нажмите 3\nЧтобы найти номер по имени человека нажмите 4\nЧтобы выйти из телефонной книги нажмите 0\n\n")

		if scanner.Scan() {
			action = scanner.Text()
		}

		switch action {
		case "1":
			fmt.Print("\n")
			root.PrintTree()
		case "2":
			fmt.Println("Введите имя")
			if scanner.Scan() {
				name = scanner.Text()
			}
			fmt.Println("Введите номера через пробел")
			if scanner.Scan() {
				numbers = strings.Split(scanner.Text(), " ")
			}
			db.AddNumber(name, numbers)
			fmt.Print("Человек успешно записан \n\n")
		case "3":
			fmt.Println("Введите имя человека запись о котором вы хотите удалить")
			if scanner.Scan() {
				name = scanner.Text()
			}
			if err := db.DeleteUser(name); err != nil {
				fmt.Printf("%s\n\n", err)
				break
			}
			fmt.Print("Запись успешно удалена\n\n")
		case "4":
			fmt.Println("Введите имя человека номер которого вы хотите найти")
			if scanner.Scan() {
				name = scanner.Text()
			}
			u := root.FindByName(name)
			if u == nil {
				fmt.Println("Записи не найдены")
				break
			}
			fmt.Printf("Имя:%s  Номера:%s \n", u.Value.Name, u.Value.PhoneNumbers)
		case "0":
			os.Exit(0)
		default:
			fmt.Print("\n")
			fmt.Println("такой команды не существует")
		}
	}
}
