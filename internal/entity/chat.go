// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

//Models!

type CreatePairRequest struct {
	Value interface{} `json:"value" example:"karaduman" `
}

type GetPairRequest struct {
	Key string
}

type Pair struct {
	Sender  string      `json:"key" `
	Message interface{} `json:"value" `
}
