package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration
type Configuration struct {
	ListenPort         string `default:":9670" split_words:"true"`
	RootURL            string `default:"/go-jti" split_words:"true"`
	OriginHost         string `default:"go-jti" split_words:"true"`
	Timeout            int64  `default:"60000" split_words:"true"`
	Addr               string `default:"localhost" split_words:"true"`
	MariaDBAddr        string `default:"localhost" split_words:"true"`
	MariaDBPort        string `default:"3306" split_words:"true"`
	MariaDBUser        string `default:"root" split_words:"true"`
	MariaDBPassword    string `default:"" split_words:"true"`
	MariaDBDatabase    string `default:"belajar-crud" split_words:"true"`
	LimitQuery         int    `default:"25" split_words:"true"`
	ClientSecret       string `default:"PlatformSecretAbsormu" split_words:"true"`
	TokenLifeTime      int64  `default:"86400" split_words:"true"`
	LimitGenerate      int    `default:"25" split_words:"true"`
	MyKeySecret        string `default:"BackendGoPlatformSecretAbsormu09" split_words:"true"`
	MyIvSecret         string `default:"SecretAbsormu212" split_words:"true"`
	IsLocal            bool   `default:"false" split_words:"true"`
	MariaDBAddrDev     string `default:"aws.connect.psdb.cloud" split_words:"true"`
	MariaDBUserDev     string `default:"n6pp26ba35kcp15s8mmi" split_words:"true"`
	MariaDBPasswordDev string `default:"pscale_pw_wl5R5wLpvKCG1NipXgMzKtoTw73jzOsddvviiMj2LSW" split_words:"true"`
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
