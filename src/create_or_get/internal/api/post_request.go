package api

import (
	"create-or-insert/internal/model"
	cryptorand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	mathrand "math/rand"
	"net/http"
	"strings"
)

func PostApiRequest(count int) error {
	if count < 1 {
		return errors.New("incorrect number of items")
	}

	for i := 0; i < count; i++ {
		item, err := GenerateItem()
		if err != nil {
			return err
		}
		arr := fmt.Sprintf("{ \"caption\": \"%v\", \"weight\": %v, \"number\": %v }", item.Caption, item.Weight, item.Number)
		r := strings.NewReader(arr)
		resp, err := http.Post("http://localhost:8080/item", "application/json", r)
		if err != nil {
			return err
		}
		err = ParsingResponse(resp, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParsingResponse(resp *http.Response, item *model.Items) error {
	decoder := json.NewDecoder(resp.Body)
	_, err := decoder.Token()
	if err != nil {
		return err
	}
	for decoder.More() {
		_, err := decoder.Token()
		if err != nil {
			return err
		}
		status := 0
		err = decoder.Decode(&status)
		if err != nil {
			return err
		}
		if status == 200 {
			fmt.Printf("Item добавлен\n caption: %v\n\n", item.Caption)
		}

	}
	defer resp.Body.Close()
	return nil
}

func GenerateItem() (*model.Items, error) {
	uniqCaption, err := getUniqCode()
	if err != nil {
		return nil, err
	}

	item := model.Items{
		Caption: uniqCaption,
		Weight:  mathrand.Float32(),
		Number:  mathrand.Int(),
	}
	return &item, nil
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
