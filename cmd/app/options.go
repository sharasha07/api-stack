package main

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
