package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/danielwetan/kdigital-backend/helpers"
	"github.com/danielwetan/kdigital-backend/models"
	uuid "github.com/satori/go.uuid"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	helpers.Headers(&w)

	if r.Method == "POST" {
		r.ParseForm()

		itemID := "id1"
		itemPrice, _ := strconv.Atoi(r.FormValue("item_price"))
		itemQuantity, _ := strconv.Atoi(r.FormValue("item_quantity"))
		itemName := r.FormValue("item_name")
		email, phone := r.FormValue("email"), r.FormValue("phone")
		grossAmount := itemPrice * itemQuantity

		uuidDigit := uuid.NewV4().String()
		unixTime := int(time.Now().Unix())
		orderID := "GPY" + strconv.Itoa(unixTime) + uuidDigit[30:]

		transactionDetails := &models.TransactionDetails{
			OrderID:     orderID,
			GrossAmount: grossAmount,
		}

		customerDetails := &models.CustomerDetails{
			FirstName: "Daniel",
			LastName:  "Saputra",
			Email:     email,
			Phone:     phone,
		}

		order := &models.Orders{
			PaymentType:        "gopay",
			TransactionDetails: *transactionDetails,
			ItemDetails: []models.ItemDetails{
				models.ItemDetails{
					ID:       itemID,
					Price:    itemPrice,
					Quantity: itemQuantity,
					Name:     itemName,
				},
			},
			CustomerDetails: *customerDetails,
			Gopay: models.Gopay{
				EnableCallback: true,
				CallbackURL:    "google.com",
			},
		}
		orderByte, _ := json.Marshal(order)
		orderString := string(orderByte)
		gopayCharge(orderString)
		transactionStatus := gopayStatus(orderID)

		res := helpers.ResponseMsg(true, transactionStatus)
		json.NewEncoder(w).Encode(res)
		// stdout := helpers.GenerateStdout(register, "application/json", statusCode, res)
	} else {
		statusCode := http.StatusBadRequest
		body := "Invalid HTTP method"
		res := helpers.ResponseMsg(false, body)
		json.NewEncoder(w).Encode(res)
		stdout := helpers.GenerateStdout(register, "application/json", statusCode, res)
		fmt.Println(stdout)
	}

}

func gopayCharge(orderString string) {
	reqBody := strings.NewReader(orderString)
	req, _ := http.NewRequest(
		"POST",
		"https://api.sandbox.midtrans.com/v2/charge",
		reqBody,
	)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1hUzFFYzRLcGNaTjhSRklYTmtiYzItNGo6")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("\nGOPAY CHARGE")
	fmt.Printf("body: %s\n\n", data)
}

func gopayStatus(orderID string) string {
	url := "https://api.sandbox.midtrans.com/v2/" + orderID + "/status"
	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1hUzFFYzRLcGNaTjhSRklYTmtiYzItNGo6")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("GOPAY STATUS")
	fmt.Printf("body: %s\n", data)
	return string(data)
}
