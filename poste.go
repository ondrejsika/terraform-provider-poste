package main

import "github.com/ondrejsika/poste-go"

func PosteApi(m interface{}) poste.PosteAPI {
	api := poste.New(m.(*Config).Origin, m.(*Config).Username,  m.(*Config).Password)
	return api
}
