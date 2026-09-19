package main

type input struct {
	moduleName   string
	wantMakefile bool

	env struct {
		loader envLoader
		parser envParser
	}
}

type envLoader int

const (
	noLoader envLoader = iota
	dotenvx
	godotenv
)

type envParser int

const (
	noParser envParser = iota
	caarlosEnv
)
