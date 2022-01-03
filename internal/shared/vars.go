package shared

import (
	"fmt"
	"github.com/sertangulveren/sypher/internal/utils"
	"os"
)

const (
	DefaultContent               = "AWS_ACCESS_KEY=123\nAWS_ACCESS_SECRET_KEY=456\n"
	DefaultName                  = "master"
	DefaultEditor                = "vim"
	ContentDir                   = "sypher"
	CannotCreateWorkingDirectory = "cannot create working directory"
	CannotReadKeyFile            = "cannot read key file"
	CannotReadEncryptedFile      = "cannot read encrypted file"
	CannotCreateTempFile         = "cannot create temp file"
	CannotWriteToTempFile        = "cannot write to temp file"
	ChangesSavedSuccessfully     = "Changes saved successfully"
	Done                         = "Done!"
)

const HelpContent = `sypher is a tool that provides you to save your credentials encrypted.

Usage:
sypher COMMAND [ARGS]

COMMANDS:
gen			Generate new credential(s).
edit		Edit credential
print		Print credential as PLAIN(All values will be printed).
gitignore	Modify or create .gitignore to ignore key files.
-help		Help.

EXAMPLES:
To generate a master credential:

	$ sypher gen

Credential generated by this command:
	sypher/master.enc
	sypher/master.key

To generate seperated credentials:

	$ sypher gen development test production

Files generated by this command:
	sypher/development.enc
	sypher/development.key
	sypher/test.enc
	sypher/test.key
	sypher/production.enc
	sypher/production.key

To edit configuration:

	$ sypher edit production

Configurations will be opened in an editor. Save your changes and close the editor.

It opens the vim editor by default. You can change this with the EDITOR variable.

Example:

	$ EDITOR=code sypher edit production

To print credentials as plain:

	$ sypher print development

Do not use this command outside of safe environments. For example: Github Actions
Decrypted content will be printed as below:
	AWS_ACCESS_KEY=123
	AWS_ACCESS_SECRET_KEY=456`

const GitIgnoreTemplate = "sypher/*.key\n"

var DefaultConfigArgs = []string{DefaultName}

func CmdArgs() []string {
	return os.Args[2:]
}

func WorkingDir() string {
	wd, err := os.Getwd()
	utils.PanicWithError(err)
	return wd
}

func GetEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}
	return DefaultEditor
}

func Help() {
	fmt.Println(HelpContent)
	os.Exit(1)
}