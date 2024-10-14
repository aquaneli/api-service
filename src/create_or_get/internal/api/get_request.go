package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetApiRequest() error {
	fmt.Println("Введите caption:")
	var str string
	fmt.Scanf("%s", &str)

	resp, err := http.Get("http://localhost:8080/item/" + str)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(string(data))
	return nil
}
