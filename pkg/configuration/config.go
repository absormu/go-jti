package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration
type Configuration struct {
	ListenPort       string `default:":9670" split_words:"true"`
	RootURL          string `default:"/go-jti" split_words:"true"`
	OriginHost       string `default:"go-jti" split_words:"true"`
	Timeout          int64  `default:"60000" split_words:"true"`
	Addr             string `default:"localhost" split_words:"true"`
	MariaDBAddr      string `default:"localhost" split_words:"true"`
	MariaDBPort      string `default:"3306" split_words:"true"`
	MariaDBUser      string `default:"root" split_words:"true"`
	MariaDBPassword  string `default:"" split_words:"true"`
	MariaDBDatabase  string `default:"db_gojti" split_words:"true"`
	LimitQuery       int64  `default:"10" split_words:"true"`
	ClientSecret     string `default:"PlatformSecretAbsormu" split_words:"true"`
	TokenLifeTime    int64  `default:"86400" split_words:"true"`
	LimitGenerate    int    `default:"25" split_words:"true"`
	MyKeySecret      string `default:"BackendGoPlatformSecretAbsormu09" split_words:"true"`
	MyIvSecret       string `default:"SecretAbsormu212" split_words:"true"`
	AuthClientKey    string `default:"989807958377-3drskv0hhpdqk5avdi68vsvon8jeq39q.apps.googleusercontent.com" split_words:"true"`
	AuthClientSecret string `default:"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" split_words:"true"`
	AuthCallbackURL  string `default:"http://localhost:9670/auth/google/callback" split_words:"true"`
}

// Config .
var Config Configuration

// LoadConfig .
func LoadConfig() {
	if err := envconfig.Process("GOJTI", &Config); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
