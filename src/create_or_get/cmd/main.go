package main

import (
	cryptorand "crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	mathrand "math/rand"
	"net/http"
	"strings"
)

type Items struct {
	caption string
	weight  float32
	number  int
}

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
		for i := 0; i < count; i++ {
			item := GenerateItem()
			arr := fmt.Sprintf("{ \"caption\": \"%v\", \"weight\": %v, \"number\": %v }", item.caption, item.weight, item.number)
			r := strings.NewReader(arr)
			resp, err := http.Post("http://localhost:8080/item", "application/json", r)
			if err != nil {
				log.Fatal(err)
			}

			decoder := json.NewDecoder(resp.Body)
			_, err = decoder.Token()
			if err != nil {
				log.Fatal(err)
			}
			for decoder.More() {
				_, err := decoder.Token()
				if err != nil {
					log.Fatal(err)
				}
				status := 0
				err = decoder.Decode(&status)
				if err != nil {
					log.Fatal(err)
				}
				if status == 200 {
					fmt.Printf("Item добавлен\n caption: %v\n\n", item.caption)
				}

			}

			defer resp.Body.Close()

		}
	}

	if code == 2 {
		fmt.Println("Введите caption:")
		var str string
		fmt.Scanf("%s", &str)

		resp, err := http.Get("http://localhost:8080/item/" + str)
		if err != nil {
			log.Fatal(err)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(data))
	}

}

func GenerateItem() Items {
	uniqCaption, err := getUniqCode()
	if err != nil {
		log.Fatal(err)
	}

	item := Items{
		caption: uniqCaption,
		weight:  mathrand.Float32(),
		number:  mathrand.Int(),
	}
	return item
}

func getUniqCode() (string, error) {
	b := make([]byte, 16)
	_, err := cryptorand.Read(b)
	if err != nil {
		return "", err
	}

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}
