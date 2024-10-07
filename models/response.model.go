package models

type ThrowErr struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

type Failed struct {
	Status string `json:"status"`
  Code int `json:"code"`
	Message string `json:"message"`
  Error interface{} `json:"error"`
}

type Success struct {
	Status string `json:"status"`
  Code int `json:"code"`
	Message string `json:"message"`
  Data interface{} `json:"data"`
}
