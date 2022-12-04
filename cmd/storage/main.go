package main

import (
	"fmt"
	"github.com/Kiatsyndesi/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("My coursework", []byte("I hate university"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
	fmt.Println(restoredFile)
}
