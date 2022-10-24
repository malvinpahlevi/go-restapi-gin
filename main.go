package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/getEmployee", getEmployee)
	router.GET("/getPahlawan", getPahlawan)
	router.GET("/getPokemon", getPokemon)
	router.GET("/getNews", getNews)

	router.Run("localhost:8080")
}

func getEmployee(c *gin.Context) {
	var employee Employee
	var arr_employee []Employee
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name, email, department FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {

		if err := rows.Scan(
			&employee.Id,
			&employee.FirstName,
			&employee.LastName,
			&employee.Email,
			&employee.Department); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_employee = append(arr_employee, employee)
		}
	}

	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = arr_employee

	c.JSON(http.StatusOK, response)
}

func getPahlawan(c *gin.Context) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	response, err := myClient.Get("https://indonesia-public-static-api.vercel.app/api/heroes")
	if err != nil {
		panic(err.Error())
		c.Status(http.StatusServiceUnavailable)
		return
	}

	// if there was no error, you should close the body
	defer response.Body.Close()

	// hence this condition is moved into its own block
	if response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))

	var responsePahlawan ResponsePahlawan
	var arr_pahlawan []Pahlawan
	json.Unmarshal(res, &arr_pahlawan)

	responsePahlawan.Status = http.StatusOK
	responsePahlawan.Message = "Success"
	responsePahlawan.Data = arr_pahlawan

	c.JSON(200, responsePahlawan)
}

func getPokemon(c *gin.Context) {

	var myClient = &http.Client{Timeout: 10 * time.Second}
	response, err := myClient.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		panic(err.Error())
		c.Status(http.StatusServiceUnavailable)
		return
	}

	// if there was no error, you should close the body
	defer response.Body.Close()

	// hence this condition is moved into its own block
	if response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePokemon ResponsePokemon
	json.Unmarshal(responseData, &responsePokemon)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    responsePokemon,
		"total":   len(responsePokemon.Pokemon),
	})
}

func getNews(c *gin.Context) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	response, err := myClient.Get("https://api-berita-indonesia.vercel.app/cnn/nasional")
	if err != nil {
		panic(err.Error())
		c.Status(http.StatusServiceUnavailable)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))

	var responseObject ResponseNews
	err = json.Unmarshal(res, &responseObject)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, responseObject)
}
