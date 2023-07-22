package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"travel-api/internal/model"
	"travel-api/internal/repository"
)

type Marker interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type MarkerImpl struct {
	markerRepository repository.Marker
}

func (m MarkerImpl) Update(ctx *gin.Context) {
	var marker model.Marker
	err := ctx.ShouldBind(&marker)
	if err != nil {
		panic(err)
	}
	updatedMarker, err := m.markerRepository.Update(marker)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, updatedMarker)
}

func (m MarkerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := m.markerRepository.Delete(id)
	if err != nil {
		panic(err)
	}
	ctx.Status(http.StatusNoContent)
}

func (m MarkerImpl) Create(ctx *gin.Context) {
	var marker model.Marker
	err := ctx.ShouldBind(&marker)
	if err != nil {
		panic(err)
	}
	createdMarker, err := m.markerRepository.Create(marker)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusCreated, createdMarker)
}

func (m MarkerImpl) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	profile, err := m.markerRepository.Get(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, profile)
}

func NewMarker() Marker {
	return MarkerImpl{
		markerRepository: repository.NewMarkerMongoRepository(),
	}
}
