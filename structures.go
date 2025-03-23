package main

type Artista struct {
	Nome string
}

type Album struct {
	Artista     Artista
	Titulo      string
	Formato     string
	Observacoes *string
}
