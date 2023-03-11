package models

type Website struct {
	Id            	int    `json:"id"`
	UrlReq        	string `json:"urlReq" validate:"url"`
	UrlSimple  		string `json:"urlReqSimple"`
	ReqStatus     	int    `json:"reqStatus"`
	ReqStatusDesc 	string `json:"reqStatusDesc"`
	Error         	string `json:"error"`
}
