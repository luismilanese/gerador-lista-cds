package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"sort"
	"text/template"
)

func LerCSV(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	registros, err := csvReader.ReadAll()
	if err != nil {
		msg := "erro ao ler os registros: " + err.Error()
		log.Fatal(msg)
		return nil, errors.New(msg)
	}

	return registros, nil
}

func GerarListaDeAlbuns(registros [][]string) ([]Album, error) {
	albums := []Album{}
	for i, registro := range registros {
		if i == 0 {
			continue
		}

		artista := Artista{}
		artista.Nome = registro[0]

		album := Album{
			Artista:     artista,
			Titulo:      registro[1],
			Formato:     registro[2],
			Observacoes: &registro[3],
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func ProcessarHtml(templateHtml, saidaHtml string, data []Album) error {
	tmpl, err := template.ParseFiles(templateHtml)
	if err != nil {
		msg := "erro ao processar template html: %v"
		log.Fatalf(msg, err.Error())
		return err
	}

	saida, err := os.Create(saidaHtml)
	if err != nil {
		msg := "erro ao criar arquivo saída: %v"
		log.Fatalf(msg, err.Error())
		return err
	}

	err = tmpl.Execute(saida, data)
	if err != nil {
		msg := "erro ao popular arquivo de saída: %v"
		log.Fatalf(msg, err.Error())
		return err
	}

	log.Print("Arquivo gerado: ", saidaHtml)
	return nil
}

func OrdenarAlfabeticamente(albuns []Album) {
	sort.Slice(albuns, func(i, j int) bool {
		return albuns[i].Artista.Nome < albuns[j].Artista.Nome
	})
}
