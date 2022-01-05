package sypher

import (
	"embed"
	"errors"
	"github.com/sertangulveren/sypher/internal/utils"
)

var Cred *Sypher

var fs *embed.FS

type Config struct {
	Name string
	Key  string
}

func RegisterFS(projectFs *embed.FS) {
	fs = projectFs
}

// Load to make ready to use sypher
func Load(config ...Config) {
	if len(config) == 0 {
		Cred = newSypher()
	} else {
		cfg := config[0]
		Cred = &Sypher{
			Name: cfg.Name,
			Key:  cfg.Key,
		}
	}

	Cred.Prepare()
}

// Get provides the string value of key
func Get(configKey string) string {
	if !Cred.Ready {
		utils.PanicWithError(errors.New("sypher is not ready"))
	}
	return string(Cred.Data[configKey])
}
