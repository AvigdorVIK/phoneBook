package main

import (
	"fmt"
)

type Record struct {
	Name   string
	Number string
	Email  string
}

//contactList := make(map[string]Record)

func addContact(name, number, email string, contactList map[string]Record) {
	contakt := Record{
		Name:   name,
		Number: number,
		Email:  email,
	}
	contactList[name] = contakt
	fmt.Println("Додано контакт:", contakt)
}

func updateContact(name, number, email string, contactList map[string]Record) {
	contakt, ok := contactList[name]
	if !ok {
		fmt.Println("нажаль, контакт:", name, " не знайден.")
		return
	}
	contakt.Number = number
	contakt.Email = email
	contactList[name] = contakt
	fmt.Println("контакт:", contakt, "оновлений!")

}

func deleteContact(name, number, email string, contactList map[string]Record) {
	_, ok := contactList[name]
	if !ok {
		fmt.Println("нажаль, контакт:", name, " не знайден.")
		return
	}
	delete(contactList, name)
	fmt.Println("контакт:", name, " було видалено!")
}

func main() {

	contactList := make(map[string]Record)

	for {
		fmt.Println("Введіть одну з наступних команд: (add, update, delete, list, exit)")
		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "add":
			fmt.Println("Введіть ім'я, телефон и email через пробел:")
			var name, phone, email string
			fmt.Scan(&name, &phone, &email)
			addContact(name, phone, email, contactList)
		case "update":
			fmt.Println("Введіть ім'я, телефон и email через пробел:")
			var name, phone, email string
			fmt.Scan(&name, &phone, &email)
			updateContact(name, phone, email, contactList)
		case "delete":
			fmt.Println("Введите ім'я:")
			var name string
			fmt.Scan(&name)
			deleteContact(name, "", "", contactList)
		case "list":
			fmt.Println("Список контактів:")
			for name, contact := range contactList {
				fmt.Println(name, contact)
			}
		case "exit":
			fmt.Println("Допобачення!")
			return
		default:
			fmt.Println("Невідома команда.")
		}
	}
}
