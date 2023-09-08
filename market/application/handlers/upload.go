package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"io"
	"market/market/domain"
	"market/market/domain/dataProviders/xtb"

	"net/http"
	"os"
)

func XtbUpload(provider domain.Provider[*xtb.CSV]) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := c.Request.ParseMultipartForm(5 * 1024 * 1024); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("parse form err: %s", err))
			return
		}

		ctx := c.Request.Context()

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err))
			return
		}

		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		xmlFile, err := os.Open(file.Filename)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("open file err: %s", err))
			return
		}
		defer func(xmlFile *os.File) {
			err := xmlFile.Close()
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("close file err: %s", err))
			}
		}(xmlFile)

		var xtbData []*xtb.CSV
		gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
			r := csv.NewReader(in)
			r.Comma = ';'
			return r
		})
		if err := gocsv.UnmarshalFile(xmlFile, &xtbData); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("unmarshal file err: %s", err))
		}

		err = provider.Insert(ctx, xtbData)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upsert trades err: %s", err))
			return
		}

		err = os.Remove(file.Filename)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("remove file err: %s", err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"upload": "success",
		})
	}
}
