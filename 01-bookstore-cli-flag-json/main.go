package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Функция считывает данные книги и созраняет их в файле

type Books struct {
	Users []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"age"`
}

func showAll() {
	jsonFile, err := os.Open("./data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var books Books
	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(books)
}

var showAllFlag bool

func main() {

	menu := "\t1) Вывести меню\n" +
		"\t2) Вывести все книги магазина\n" +
		"\t3) Добавить книгу в магазин\n" +
		"\t0) Выход из программы\n"

	fmt.Print("Это программа для вывода содержимого файла JSON" +
		"с данными о книгах в книжном магазине\n" +
		"Для продолжения работы выберите нужный пункт меню\n" + menu)

	for {
		var key = -1
		fmt.Scanf("%d\n", &key)

		switch key {
		case 0:
			fmt.Printf("Выход...")
			os.Exit(0)
		case 1:
			fmt.Printf("Выберите нужный пункт меню:\n" + menu)
		case 2:
			fmt.Printf("Вывод:")
			showAll()
		case 3:
			//
		default:
			fmt.Printf("Ввод данных был произведен неверно! Повторите ввод: ")
		}
	}
}
