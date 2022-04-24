package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/madrid-recicla/server/src/mapper"
	"github.com/madrid-recicla/server/src/repository"
)

type clothesHandler struct {
	repository repository.ContainerRepository
}

func NewClothesHandler() ContainerHandler {
	return &clothesHandler{repository.NewClothesRepository()}
}

func (h *clothesHandler) Load() gin.HandlerFunc {
	return func(c *gin.Context) {
		geoRssBytes := getFileContent(c)
		if geoRssBytes == nil {
			return
		}

		containers, err := mapper.GetGeoRssMapper().MapToModelContainers(geoRssBytes)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error while loading ClothesContainers", err)
			return
		}

		if _, err = h.repository.DeleteAll(); err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error while deleting ClothesContainers", err)
			return
		}
		numIns, err := h.repository.AddAll(containers)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error while adding ClothContainers", err)
			return
		}
		c.String(http.StatusOK, "%d ClothesContainers loaded", numIns)
	}
}

func (h *clothesHandler) ListAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		clothContainers, err := h.repository.ListAll()
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error while fetching ClothContainers", err)
			return
		}

		h := requestHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println("Error while binding header", err)
		}

		if strings.EqualFold(h.ResponseFormat, "BasicCollection") {
			c.JSON(http.StatusOK, formatAsBasic(clothContainers))
		} else {
			c.Writer.Header().Set("Content-Type", "application/geo+json")
			c.JSON(http.StatusOK, formatAsGeoJson(clothContainers))
		}
	}
}

func (h *clothesHandler) DeleteAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.repository.DeleteAll()
	}
}

func getFileContent(c *gin.Context) *[]byte {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error processing the file", err)
		return nil
	}
	file, err := fileHeader.Open()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Error processing the file", err)
		return nil
	}
	defer file.Close()
	var buf bytes.Buffer
	io.Copy(&buf, file)
	content := buf.Bytes()
	buf.Reset()
	return &content
}
