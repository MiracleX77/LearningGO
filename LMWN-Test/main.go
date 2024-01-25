package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Case struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             string `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}
type DataCase struct {
	Data []Case `json:"Data"`
}
type Response struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

func main() {
	app := gin.Default()
	app.GET("/covid/summary", func(c *gin.Context) {
		data, err := getData()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response := summaryData(data)
		c.JSON(200, response)
	})
	app.Run(":8080")
}

func getData() (DataCase, error) {
	url := "https://static.wongnai.com/devinterview/covid-cases.json"
	response, err := http.Get(url)
	if err != nil {
		return DataCase{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return DataCase{}, err
	}
	var data DataCase
	if err := json.Unmarshal(body, &data); err != nil {
		return DataCase{}, err
	}
	return data, nil

}
func summaryData(data DataCase) Response {
	var response Response
	response.Province = make(map[string]int)
	response.AgeGroup = make(map[string]int)
	for _, caseCovid := range data.Data {
		if caseCovid.Province == "" {
			response.Province["N/A"]++
		} else {
			response.Province[caseCovid.Province]++
		}
		if caseCovid.Age == 0 {
			response.AgeGroup["N/A"]++
		} else {
			if caseCovid.Age > 0 && caseCovid.Age <= 30 {
				response.AgeGroup["0-30"]++
			} else if caseCovid.Age >= 31 && caseCovid.Age <= 60 {
				response.AgeGroup["31-60"]++
			} else {
				response.AgeGroup["61+"]++
			}
		}
	}
	return response
}
