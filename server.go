package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/ddosify/go-faker/faker"
)

type product struct {
	Id			int64	`json:"id" query:"id"`
	Name 		string 	`json:"name"`
	Price		float64 `json:price"`
	Quantity 	int 	`json:"quantity"`
	Image		string	`json:"image_url"`
}

var products []product

func initProducts(totalProducts int) {
	faker := faker.NewFaker()

	fmt.Printf("Create %d products, please wait...", totalProducts)
	for i := range totalProducts {
		float1, error1 := strconv.ParseFloat(faker.RandomPrice(), 64)
		if error1 != nil {
			fmt.Println(error1)
		}
		product := product{
			int64(i),
			faker.RandomProductName(),
			float1,
			faker.RandomInt(),
			faker.RandomImageURL(),
		}
		products = append(products, product)
	}
}

func getProduct(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	// Loop over products
	for i,product := range products {
		if products[i].Id == id {
			fmt.Println(product)
			return c.JSON(http.StatusOK, product)
		}
	}

	return c.JSON(http.StatusNotFound, "not found")
}

func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func main() {
	initProducts(50)
	
	e := echo.New()
	e.GET("/", getAllProducts)
	e.GET("/:id", getProduct)

	e.Logger.Fatal(e.Start(":8080"))
}