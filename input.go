package main

type input struct {
	moduleName string

	env struct {
		loader        envLoader
		parser        envParser
		secretManager envSecretManager
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

type envSecretManager int

const (
	noSecretManager envSecretManager = iota
	dotenvxSecrets
	awsSecretsManager
	googleCloudSecretManager
	azureKeyVault
)
