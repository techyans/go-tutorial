package main

import (
	"bufio"
	"fmt"
	"go_course/note"
	"go_course/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	printSomething("Arshad")
	printSomething(1)
	printSomething(1.3)

	title, content := GetNoteData()

	todotext := GetUserInput("Todo: ")

	todo, err := todo.New(todotext)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	err = outPutData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	UserNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outPutData(UserNote)
	// err = UserNote.Save()
	// if err != nil {
	// 	fmt.Println("Error in file saving", err)
	// 	return
	// }

	// fmt.Println("Note saved successfully")
	// fmt.Println("Title:", title)
	// fmt.Println("Content:", content)
}

func printSomething(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("Integer: ", data)
	case float64:
		fmt.Println("Float: ", data)
	case string:
		fmt.Println("String: ", data)
	}
}

func outPutData(data outputable) error {
	data.Display()
	return SaveData(data)
}
func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}
func GetNoteData() (string, string) {
	titel := GetUserInput("Note title:")

	content := GetUserInput("Note content:")

	return titel, content
}

func GetUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
