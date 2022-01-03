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
func (s *Sypher) Read() []byte {
	encData, err := os.ReadFile(s.FileName())
	utils.ExitWithMessage(err, shared.CannotReadEncryptedFile)


	bData := utils.DecodeBase64(encData)
	data := utils.Decrypt(s.Key, bData)
	s.Data = make(map[string][]byte)
	for _, line := range bytes.Split(data, []byte("\n")) {
		eqSignIndex := bytes.Index(line, []byte("="))
		if eqSignIndex == -1 {
			continue
		}
		s.Data[string(line[:eqSignIndex])] = line[eqSignIndex+1:]
	}
	return data
}

func (s *Sypher) Write(value []byte) {
	encrypted := utils.Encrypt(s.Key, value)
	base64Data := utils.EncodeBase64(encrypted)

	err := os.WriteFile(s.FileName(), base64Data, os.ModePerm)
	utils.PanicWithError(err)
}

func (s *Sypher) WriteKey() {
	err := os.WriteFile(s.KeyFileName(), []byte(s.Key), os.ModePerm)
	utils.PanicWithError(err)
}

