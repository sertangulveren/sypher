package sypher

import (
	"bytes"
	"github.com/sertangulveren/sypher/internal/shared"
	"github.com/sertangulveren/sypher/internal/utils"
	"os"
	"path/filepath"
)

type Sypher struct {
	Name string
	Key  string
	Data map[string][]byte
	Ready bool
}

func (s *Sypher) RootFilePath() string {
	return filepath.Join(shared.WorkingDir(), shared.ContentDir, s.Name)
}

func (s *Sypher) FileName() string {
	return s.RootFilePath() + ".enc"
}

func (s *Sypher) KeyFileName() string {
	return s.RootFilePath() + ".key"
}

func (s *Sypher) ReadKeyFile() {
	keyData, err := os.ReadFile(s.KeyFileName())
	utils.ExitWithMessage(err, shared.CannotReadKeyFile)
	s.Key = string(keyData)
}

func (s *Sypher) WriteKey() {
	err := os.WriteFile(s.KeyFileName(), []byte(s.Key), os.ModePerm)
	utils.PanicWithError(err)
}

