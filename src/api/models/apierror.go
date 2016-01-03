package models


type ApiError struct{
	Code string 	`json:"code"`
	Message string  `json:"message"`
}