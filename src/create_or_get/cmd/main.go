package main

import (
	"create-or-insert/internal/api"
	"fmt"
	"log"
)

func main() {
	code := 0
	fmt.Println("1.Добавить item.\n2.Получить item.\nВведите значение:")
	fmt.Scanf("%d", &code)
	if code != 1 && code != 2 {
		log.Fatal("incorrect code")
	}
	if code == 1 {
		fmt.Println("Введите количество item:")
		count := 0
		fmt.Scanf("%d", &count)
		err := api.PostApiRequest(count)
		if err != nil {
			log.Fatal(err)
		}

	}
	if code == 2 {
		err := api.GetApiRequest()
		if err != nil {
			log.Fatal(err)
		}
	}
}
