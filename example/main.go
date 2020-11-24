package main

import (
	"fmt"
	"github.com/Bhinneka/ammana-go"
	"time"
)

const (
	BaseURL = "https://paylater-dev.ammana.id/api/v1"
	ClientID = "da89901f-4b2e-438b-b707-f93645b39631"
	ClientSecret = "3b5b8d854e53662d91a615d7705d61fe"
	DefaultTimeout = 10 * time.Second
)

func main() {
	fmt.Print("AMMANA")

	newAmmana := ammana.New(BaseURL, ClientID, ClientSecret, DefaultTimeout)
	getCardDetail(newAmmana)
}

func getCardDetail(amm ammana.AmmanaService) {
	var reqCardDetail ammana.GetCardDetailRequest
	reqCardDetail.PhoneNumber = "085912347789"

	result := amm.GetCardDetail(reqCardDetail)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}

	getCardDetailResponse, ok := result.Result.(ammana.GetCardDetailResponse)
	if !ok {
		fmt.Println("Result is not Get Card Detail Response")
		return
	}
	fmt.Println(getCardDetailResponse)
}
