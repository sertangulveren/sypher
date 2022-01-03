package sypher

import (
	"errors"
	"github.com/sertangulveren/sypher/internal/utils"
)

var Cred *Sypher

func Load(name string, key string)  {
	Cred = &Sypher{
		Name: name,
		Key:  key,
	}
	Cred.Prepare()
}

func Get(configKey string) string {
	if !Cred.Ready {
		utils.PanicWithError(errors.New("sypher is not ready"))
	}
	return string(Cred.Data[configKey])
}