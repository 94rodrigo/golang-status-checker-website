package controllers

import (
	"encoding/json"
	"fmt"
	"main/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var sitesList []models.Website
var jsonResp string
var idReq int = 0
var urlReq string = ""

func VerificaSites(c *gin.Context) {
	urlReq = c.Request.FormValue("urlReqIpt")
	if urlReq == "" {
		ShowIndexPage(c)
		return
	}
	reqInformation := getReqInformation(urlReq)
	if reqInformation.Error == "" {
		addToReqList(reqInformation)
	}
	jsonConvert, _ := json.Marshal(&reqInformation)
	jsonResp = string(jsonConvert)
	urlReq = ""
	ShowIndexPage(c)
}

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusMovedPermanently, "index.html", gin.H{
		"sitesList": sitesList,
		"jsonResp":  jsonResp,
	})
}

func addToReqList(website models.Website) {
	if website.Error == "" {
		sitesList = append(sitesList, website)
	}
}

func getReqInformation(address string) models.Website {
	idReq += 1
	website := models.Website{Id: idReq, UrlReq: address, UrlSimple: removeUrlPrefix(address)}
	validate := validator.New()
	err := validate.Struct(website)

	if err != nil {
		website.Error = "The requested URL is not in the correct format (example: http://www.domain.com)"
	} else {
		statusUrl, err := http.Get(address)
		if err != nil {
			website.ReqStatus = 404
			website.ReqStatusDesc = "404 Not Found"
		} else {
			website.ReqStatus = statusUrl.StatusCode
			website.ReqStatusDesc = statusUrl.Status
			defer statusUrl.Body.Close()
		}
	}

	return website
}

func removeUrlPrefix(address string) string {
	var simpleUrl string
	if strings.Contains(address, "https://") {
		simpleUrl = simpleUrl + strings.Replace(address, "https://", "", 1)
	} else if strings.Contains(address, "http://") {
		simpleUrl = simpleUrl + strings.Replace(address, "http://", "", 1)
	}
	return simpleUrl
}

func VerificaSitePorId(c *gin.Context) {
	var foundWebsite models.Website
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	for _, v := range sitesList {
		if v.Id == id {
			foundWebsite = v
			break
		}
	}
	newReq, err := http.Get(foundWebsite.UrlReq)
	if err == nil {
		foundWebsite.ReqStatus = newReq.StatusCode
		foundWebsite.ReqStatusDesc = newReq.Status
	} else {
		foundWebsite.ReqStatus = 404
		foundWebsite.ReqStatusDesc = "404 Not Found"
	}
	c.JSON(http.StatusOK, foundWebsite)
}

func DeleteSiteFromList(c *gin.Context) {
	fmt.Println("aaaaaaaaaa")
	var foundWebsite models.Website
	var index int = 0
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	for _, v := range sitesList {
		if v.Id == id {
			foundWebsite = v
			fmt.Println(index)
			break
		}
		index += 1
	}
	sitesList = remove(sitesList, index)
	c.JSON(http.StatusOK, foundWebsite)
}

func remove(s []models.Website, i int) []models.Website {
	var newSlice []models.Website
	newSlice = append(newSlice, s[:i]...)
	newSlice = append(newSlice, s[i+1:]...)
	//return append(s[:i], s[i+1:]...)
	return newSlice
}
