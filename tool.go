package sypher

import (
	"embed"
	"errors"
	"github.com/sertangulveren/sypher/internal/procs"
	"github.com/sertangulveren/sypher/internal/utils"
)

var cred *procs.Sypher

type Config struct {
	Name string
	Key  string
}

func RegisterFS(projectFs *embed.FS) {
	procs.FS = projectFs
}

// Load to make ready to use sypher
func Load(config ...Config) {
	if len(config) == 0 {
		cred = procs.NewSypher()
	} else {
		cfg := config[0]
		cred = &procs.Sypher{
			Name: cfg.Name,
			Key:  cfg.Key,
		}
	}

	cred.Prepare()
}

// Get provides the string value of key
func Get(configKey string) string {
	if !cred.Ready {
		utils.PanicWithError(errors.New("sypher is not ready"))
	}
	return string(cred.Data[configKey])
}
