package main

import (
	"fmt"

	"github.com/GusTeixeira/alura-api-go-gin/models"
	"github.com/GusTeixeira/alura-api-go-gin/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
