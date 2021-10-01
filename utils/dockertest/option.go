package dockertest

import "github.com/spf13/viper"

type SetupOption func(o *setupOptions)

type setupOptions struct {
	DockerHost     string
	DockerImage    string
	DockerImageTag string

	DBUsername string
	DBPassword string
	DBName     string
}

func parseOptions(opts ...SetupOption) *setupOptions {
	viper.AutomaticEnv()
	viper.SetDefault("DOCKERTEST_HOST", "docker")
	viper.SetDefault("DOCKERTEST_IMAGE", "postgres")
	viper.SetDefault("DOCKERTEST_IMAGE_TAG", "12.4")

	options := &setupOptions{
		DockerHost:     viper.GetString("DOCKERTEST_HOST"),
		DockerImage:    viper.GetString("DOCKERTEST_IMAGE"),
		DockerImageTag: viper.GetString("DOCKERTEST_IMAGE_TAG"),

		DBUsername: "postgres",
		DBPassword: "postgres",
		DBName:     "database",
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithDBPassword(password string) SetupOption {
	return func(o *setupOptions) {
		o.DBPassword = password
	}
}

func WithDBName(name string) SetupOption {
	return func(o *setupOptions) {
		o.DBName = name
	}
}
