package mapper

import (
	"encoding/xml"
	"log"
	"sync"

	"github.com/madrid-recicla/server/src/model"
)

var geoRssLock = &sync.Mutex{}

type GeoRssMapper struct {
}

var geoRssMapper *GeoRssMapper

func GetGeoRssMapper() *GeoRssMapper {
	if geoRssMapper == nil {
		geoRssLock.Lock()
		defer geoRssLock.Unlock()
		if geoRssMapper == nil {
			geoRssMapper = &GeoRssMapper{}
		}
	}
	return geoRssMapper
}

func (m *GeoRssMapper) MapToModelContainers(geoRssBytes *[]byte) (*[]*model.Location, error) {
	var cXml model.ContainersFeedXml
	if err := xml.Unmarshal(*geoRssBytes, &cXml); err != nil {
		log.Println(err)
		return nil, err
	}

	var container *model.Location
	var containers []*model.Location
	for _, xmlContainer := range cXml.Containers {
		if xmlContainer.GeoRssPoint == nil {
			continue
		}
		container = &model.Location{
			Properties: model.Properties{
				Title: xmlContainer.Title,
			},
			Coordinates: model.Coordinates{
				Latitude:  xmlContainer.GeoRssPoint.Lat,
				Longitude: xmlContainer.GeoRssPoint.Lon,
			},
		}
		containers = append(containers, container)
	}
	log.Printf("Successfully parsed %d containers", len(containers))
	return &containers, nil
}
