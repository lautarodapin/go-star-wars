package people

import (
	"fmt"
	"net/http"
	"star-wars-api/utils"

	"github.com/gin-gonic/gin"
)

const LIST_URL = "https://swapi.dev/api/people/"
const DETAIL_URL = "https://swapi.dev/api/people/%s/"

type Person struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld *string   `json:"homeworld"`
	Films     *[]string `json:"films"`
	Vehicles  *[]string `json:"vehicles"`
	Starships *[]string `json:"starships"`
	Species   *[]string `json:"species"`
	Detail    *string   `json:"url"`
}

func PeopleList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var persons []Person
		next := LIST_URL
		i := 1
		for next != "" {
			var apiResponse utils.ApiResponse[Person]
			err := utils.MakeRequest(fmt.Sprintf("%s?page=%d", LIST_URL, i), &apiResponse)
			fmt.Printf("Iteracion %d", i)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.Response[any]{
					Status: "error",
					Detail: err.Error(),
				})
				return
			}
			if apiResponse.Next == nil {
				next = ""
			} else {
				next = *apiResponse.Next
			}
			persons = append(persons, apiResponse.Results...)

			i++

		}
		ctx.JSON(http.StatusOK, utils.Response[[]Person]{
			Status: "success",
			Data:   persons,
		})
	}
}

func PeopleDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var person Person
		id := ctx.Param("id")
		err := utils.MakeRequest(fmt.Sprintf(DETAIL_URL, id), &person)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.Response[any]{
				Status: "error",
				Detail: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, utils.Response[Person]{
			Status: "success",
			Data:   person,
		})
	}

}
