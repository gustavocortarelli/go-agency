package utils

import (
	"fmt"
	"github.com/gustavocortarelli/go-agency/internal/model"
	"github.com/gustavocortarelli/go-agency/internal/service/costumer"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
const numberBytes = "0123456789"

func RandStringBytes(n int, numbersOnly bool) string {
	b := make([]byte, n)
	for i := range b {
		if numbersOnly {
			b[i] = numberBytes[rand.Intn(len(numberBytes))]
		} else {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
	}
	return string(b)
}

func GetRandomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GenerateAndInsertData(numberOfRecords int) {
	var addresses []model.CostumerAddress
	fmt.Println("Create addresses:", time.Now().Format("2006-01-02 03:04:05"))
	for i := 0; i < numberOfRecords; i++ {
		addresses = append(addresses, model.CostumerAddress{
			CityID:  int64(GetRandomNumber(1, 9)),
			Address: RandStringBytes(GetRandomNumber(10, 120), false),
			ZipCode: RandStringBytes(GetRandomNumber(6, 10), true),
		})
	}
	fmt.Println("All addresses created:", time.Now().Format("2006-01-02 03:04:05"))

	costumerInsData := model.Costumer{
		Name:      fmt.Sprintf("NAME %d", GetRandomNumber(1, 9)),
		Surname:   fmt.Sprintf("SURNAME [%d]", numberOfRecords),
		DocNumber: "DOCNUMBER",
		Birthdate: model.Date(time.Now()),
		Addresses: addresses,
	}
	fmt.Println("[CreateBatchSize: 1000] inserting costumer:", time.Now().Format("2006-01-02 03:04:05"))
	_, err := costumer.Create(costumerInsData)
	fmt.Println("Finishes insert:", time.Now().Format("2006-01-02 03:04:05"), err)
}
