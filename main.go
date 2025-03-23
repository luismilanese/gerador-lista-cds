package main

import (
	"log"
)

func main() {
	registros, err := LerCSV("registros.csv")
	if err != nil {
		log.Fatalf("erro ao ler o arquivo CSV: %v", err)
	}

	albuns, err := GerarListaDeAlbuns(registros)
	if err != nil {
		msg := err.Error()
		log.Fatal("erro ao gerar lista de albuns: " + msg)
	}

	OrdenarAlfabeticamente(albuns)

	err = ProcessarHtml("template.html", "index.html", albuns)
	if err != nil {
		msg := "erro ao processar html: %v"
		log.Fatalf(msg, err.Error())
	}
}
