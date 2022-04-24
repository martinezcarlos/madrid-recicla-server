package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madrid-recicla/server/src/model"
	geojson "github.com/paulmach/go.geojson"
)

type ContainerHandler interface {
	Load() gin.HandlerFunc
	ListAll() gin.HandlerFunc
	DeleteAll() gin.HandlerFunc
}

type requestHeader struct {
	ResponseFormat string `header:"Response-Format"`
}

func respondWithError(c *gin.Context, code int, message interface{}, e error) {
	log.Println(message, e)
	c.Error(e)
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func formatAsGeoJson(locations *[]*model.Location) *geojson.FeatureCollection {
	fc := geojson.NewFeatureCollection()
	for _, cont := range *locations {
		point := geojson.
			NewPointFeature([]float64{cont.Coordinates.Longitude, cont.Coordinates.Latitude})
		point.SetProperty("title", cont.Properties.Title)
		fc.AddFeature(point)
	}
	return fc
}

func formatAsBasic(locations *[]*model.Location) *model.BasicCollection {
	return &model.BasicCollection{
		Type:      "BasicCollection",
		Locations: *locations,
	}
}
