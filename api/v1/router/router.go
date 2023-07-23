package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"net/http"
	"travel-api/api/v1/handlers"
	"travel-api/internal/config"
	"travel-api/locale"
)

type Gin interface {
	CreateRouter() *gin.Engine
	ErrorHandler(c *gin.Context, err any)
}

type GinImpl struct {
	markerHandler  handlers.Marker
	countryHandler handlers.Country
}

func (r GinImpl) CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(r.ErrorHandler))
	router.Use(CORSMiddleware())
	v1Routes := router.Group("/api/v1")
	{
		markerRoutes := v1Routes.Group("/marker")
		{
			markerRoutes.POST("", r.markerHandler.Create)
			markerRoutes.GET("/:id", r.markerHandler.Get)
			markerRoutes.PUT("/:id", r.markerHandler.Update)
			markerRoutes.DELETE("/:id", r.markerHandler.Delete)
		}
		countryRoutes := v1Routes.Group("/country")
		{
			countryRoutes.POST("", r.countryHandler.Create)
			countryRoutes.GET("/:id", r.countryHandler.Get)
			countryRoutes.GET("", r.countryHandler.GetAll)
			countryRoutes.PUT("/:id", r.countryHandler.Update)
			countryRoutes.DELETE("/:id", r.countryHandler.Delete)
			countryRoutes.GET("/:id/markers", r.countryHandler.GetMarkers)
		}
	}
	err := router.Run(":" + config.GetProperty("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (r GinImpl) ErrorHandler(c *gin.Context, err any) {
	switch err.(type) {
	case validator.ValidationErrors:
		errorsAmount := len(err.(validator.ValidationErrors))
		fieldErrors := make([]gin.H, errorsAmount)
		for i, fieldError := range err.(validator.ValidationErrors) {
			fieldErrors[i] = gin.H{"namespace": fieldError.Namespace(), "field": fieldError.Field(), "tag": fieldError.Tag()}
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": fieldErrors})
	case mongo.WriteException:
		fieldErrors := make([]gin.H, len(err.(mongo.WriteException).WriteErrors))
		for i, writeError := range err.(mongo.WriteException).WriteErrors {
			doc, err := bsonx.ReadDoc(writeError.Raw)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.(error).Error()})
			}
			duplicatedKey := doc[doc.IndexOf("keyValue")].Value.Document()[0].Key
			fieldErrors[i] = gin.H{"message": locale.GetMessageLocaleFromRequest("Duplicated Key", c, map[string]string{
				"key": duplicatedKey,
			})}
		}
		status := http.StatusBadRequest
		if err.(mongo.WriteException).HasErrorCode(11000) {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{"errors": fieldErrors})
	case error:
		if err.(error).Error() == mongo.ErrNoDocuments.Error() {
			id := c.Param("id")
			c.JSON(http.StatusNotFound, gin.H{"message": locale.GetMessageLocaleFromRequest("Not Found", c, map[string]string{
				"Name": id,
			})})
		} else if err.(error).Error() == primitive.ErrInvalidHex.(error).Error() {
			id := c.Param("id")
			c.JSON(http.StatusBadRequest, gin.H{"message": locale.GetMessageLocaleFromRequest("Invalid Hex", c, map[string]string{
				"Hex": id,
			})})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.(error).Error()})
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.(error).Error()})
	}
}
func NewGin() GinImpl {
	return GinImpl{
		markerHandler:  handlers.NewMarker(),
		countryHandler: handlers.NewCountry(),
	}
}
