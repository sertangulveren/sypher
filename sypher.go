package sypher

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/sertangulveren/sypher/internal/shared"
	"github.com/sertangulveren/sypher/internal/utils"
	"os"
	"path/filepath"
)

//go:embed sypher/*.enc
var encryeptedCredentials embed.FS

type Sypher struct {
	Name  string
	Key   string
	Data  map[string][]byte
	Ready bool
}

func NewSypher() *Sypher {
	return &Sypher{Name: shared.DefaultName}
}

func (s *Sypher) RootFilePath() string {
	return filepath.Join(shared.WorkingDir(), shared.ContentDir, s.Name)
}

func (s *Sypher) FileName() string {
	return s.RootFilePath() + ".enc"
}

func (s *Sypher) EmbedPath() string {
	return "sypher/" + s.Name + ".enc"
}

func (s *Sypher) KeyFileName() string {
	return s.RootFilePath() + ".key"
}

func (s *Sypher) ReadKeyFile() {
	keyData, err := os.ReadFile(s.KeyFileName())
	utils.ExitWithMessage(err, shared.CannotReadKeyFile)
	s.Key = string(keyData)
}

func (s *Sypher) readEncryptedContent() []byte  {
	encData, err := os.ReadFile(s.FileName())
	if err == nil {
		fmt.Println("sypher loaded the content from encrypted file")
		return encData
	}

	encData, err = encryeptedCredentials.ReadFile(s.EmbedPath())
	if err == nil {
		fmt.Println("sypher loaded the content from embedded file")
		return encData
	}
	utils.ExitWithMessage(err, shared.CannotReadEncryptedFile)
	return nil
}
func (s *Sypher) Read() []byte {
	encData := s.readEncryptedContent()

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

func (s *Sypher) Prepare() {
	if s.Name == "" {
		s.Name = shared.DefaultName
	}
	if s.Key == "" {
		if s.Key = os.Getenv("SYPHER_MASTER_KEY"); s.Key == "" {
			s.ReadKeyFile()
		}
	}

	// read data and set sypher data
	s.Read()
	s.Ready = true
}
