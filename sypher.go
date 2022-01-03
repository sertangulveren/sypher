package sypher

import (
	"github.com/sertangulveren/sypher/internal/shared"
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
