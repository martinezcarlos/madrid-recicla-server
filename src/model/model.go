package model

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Properties struct {
	Title string `json:"title"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Properties  Properties  `json:"properties"`
	Coordinates Coordinates `json:"coordinates"`
}

type BasicCollection struct {
	Type      string      `json:"type"`
	Locations []*Location `json:"locations"`
}

type ContainersFeedXml struct {
	XMLName    xml.Name        `xml:"http://www.w3.org/2005/Atom feed"`
	Containers []*ContainerXml `xml:"http://www.w3.org/2005/Atom entry,omitempty"`
}

type ContainerXml struct {
	XMLName     xml.Name     `xml:"http://www.w3.org/2005/Atom entry,omitempty"`
	Title       string       `xml:"title"`
	GeoRssPoint *GeoRssPoint `xml:"http://www.georss.org/georss point,omitempty"`
}

type GeoRssPoint struct {
	Lat float64
	Lon float64
}

func (v GeoRssPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(fmt.Sprintf("%f %f", v.Lat, v.Lon), start)
	return nil
}

func (c *GeoRssPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	res := strings.Fields(v)
	if len(res) != 2 {
		log.Println("Error parsing GeoRssPoint: ", v)
		return nil
	}
	lat, err := strconv.ParseFloat(res[0], 64)
	if err != nil {
		log.Println("Error parsing Latitude: ", err)
		return nil
	}
	lon, err := strconv.ParseFloat(res[1], 64)
	if err != nil {
		log.Println("Error parsing Longitude: ", err)
		return nil
	}
	*c = GeoRssPoint{Lat: lat, Lon: lon}
	return nil
}
