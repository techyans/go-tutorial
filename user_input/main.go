package main

import (
	"bufio"
	"fmt"
	"go_course/note"
	"os"
	"strings"
)

func main() {
	title, content := GetNoteData()

	UserNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	UserNote.Display()

	err = UserNote.Save()
	if err != nil {
		fmt.Println("Error in file saving", err)
		return
	}

	fmt.Println("Note saved successfully")
	// fmt.Println("Title:", title)
	// fmt.Println("Content:", content)
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
