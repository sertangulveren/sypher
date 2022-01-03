package commander

import (
	"fmt"
	"github.com/sertangulveren/sypher"
	"github.com/sertangulveren/sypher/internal/shared"
	"github.com/sertangulveren/sypher/internal/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Generate new credential(s)
func Generate() {
	if _, err := os.Stat(shared.ContentDir); os.IsNotExist(err) {
		err := os.Mkdir(shared.ContentDir, os.ModePerm)
		utils.ExitWithMessage(err, shared.CannotCreateWorkingDirectory)
	}
	pieces := shared.CmdArgs()
	if len(pieces) == 0 {
		pieces = shared.DefaultConfigArgs
	}

	// create credential for each pieces
	for _, item := range pieces {
		s := sypher.Sypher{
			Name: item,
			Key:  utils.GenerateKey(),
		}
		if _, err := os.Stat(s.FileName()); err == nil {
			fmt.Println("Ignored: ", s.Name)
			continue
		}
		s.Write([]byte(shared.DefaultContent))
		s.WriteKey()
		fmt.Println("Created: ", s.Name)
	}
	fmt.Println(shared.Done)
}
// Print credential as plain
func Print() {
	s := sypher.Sypher{}
	if len(shared.CmdArgs()) > 0 {
		s.Name = shared.CmdArgs()[0]
	}

	s.Prepare()
	//fmt.Println(s.Data)
	for k, v := range s.Data {
		fmt.Printf("%s=%s\n", k, string(v))
	}
	fmt.Println(shared.Done)
}