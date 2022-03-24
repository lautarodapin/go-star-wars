package people

import (
	"fmt"
	"net/http"
	"star-wars-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var LIST_URL = fmt.Sprintf("%s/people/", utils.BASE_URL)
var DETAIL_URL = LIST_URL + "%s/"

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
		path := ctx.Request.URL.Path
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		var apiResponse utils.ApiResponse[Person]
		err := utils.MakeRequest(fmt.Sprintf("%s?page=%d", LIST_URL, page), &apiResponse)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.Response[any]{
				Status: "error",
				Detail: err.Error(),
			})
			return
		}
		if apiResponse.Next != nil {
			next := fmt.Sprintf("%s?page=%d", path, page+1)
			apiResponse.Next = &next
		}
		if apiResponse.Previous != nil {
			previous := fmt.Sprintf("%s?page=%d", path, page-1)
			apiResponse.Previous = &previous
		}

		ctx.JSON(http.StatusOK, utils.Response[utils.ApiResponse[Person]]{
			Status: "success",
			Data:   apiResponse,
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
