package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/ivmax96/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file := st.Upload("text.txt", []byte("hello"))
	_ = file

	rf, err := st.GetByID(uuid.New())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rf)
}
