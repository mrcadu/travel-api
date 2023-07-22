package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"travel-api/internal/model"
	"travel-api/internal/repository"
)

type Country interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetMarkers(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type CountryImpl struct {
	countryRepository repository.Country
}

func (c CountryImpl) GetMarkers(ctx *gin.Context) {
	id := ctx.Param("id")
	markers, err := c.countryRepository.GetMarkers(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, markers)
}

func (c CountryImpl) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	country, err := c.countryRepository.Get(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, country)
}

func (c CountryImpl) GetAll(ctx *gin.Context) {
	countries, err := c.countryRepository.GetAll()
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, countries)
}

func (c CountryImpl) Update(ctx *gin.Context) {
	var country model.Country
	err := ctx.ShouldBind(&country)
	if err != nil {
		panic(err)
	}
	updatedCountry, err := c.countryRepository.Update(country)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, updatedCountry)
}

func (c CountryImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := c.countryRepository.Delete(id)
	if err != nil {
		panic(err)
	}
	ctx.Status(http.StatusNoContent)
}

func (c CountryImpl) Create(ctx *gin.Context) {
	var country model.Country
	err := ctx.ShouldBind(&country)
	if err != nil {
		panic(err)
	}
	createdCountry, err := c.countryRepository.Create(country)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusCreated, createdCountry)
}

func NewCountry() Country {
	return CountryImpl{
		countryRepository: repository.NewCountryMongoRepository(),
	}
}
