package handlers

import "github.com/gin-gonic/gin"

type Markers interface {
	Get(ctx *gin.Context)
}

type MarkersImpl struct {
}

func (m MarkersImpl) Get(ctx *gin.Context) {

}

func NewMarkers() Markers {
	return MarkersImpl{}
}
