package main

const (
	development = iota + 1
	testing
	production
	developmentWithoutTokens // DO NOT USE IN PRODUCTION
)

type envConfigurations struct {
	development              bool
	testing                  bool
	developmentWithoutTokens bool
	production               bool
	invalidConfiguration     bool
}

func Environments() *envConfigurations {
	// edita esto si deseas cambiar la configuracion
	env := development

	envConfigurations := &envConfigurations{
		development:              false,
		developmentWithoutTokens: false,
		testing:                  false,
		production:               false,
		invalidConfiguration:     false,
	}

	if env > 0 && env < 5 {
		switch env {
		case development:
			envConfigurations.development = true
			return envConfigurations
		case testing:
			envConfigurations.testing = true
			return envConfigurations
		case production:
			envConfigurations.production = true
			return envConfigurations
		case developmentWithoutTokens:
			envConfigurations.developmentWithoutTokens = true
			return envConfigurations
		}
	}

	envConfigurations.invalidConfiguration = true
	return envConfigurations

}
