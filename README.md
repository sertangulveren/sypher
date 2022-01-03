# sypher[⚠️Work in progress]

sypher provides you to store your credentials and secrets as encrypted in your repository.

## Usage

### Install the command line application:

```sh
go get github.com/sertangulveren/sypher
```

### Generate credential(s):

The ```gen``` command as below will create your credentials under the sypher folder in your working directory:

```sh
sypher gen
```

It will be generate your credentials as below:
```sh
sypher
├── master.enc
└── master.key
```

You can provide names to generator.

For example:
```sh
sypher gen development test production
```

The program will generate files as below:
```sh
sypher
├── development.enc
├── development.key
├── production.enc
├── production.key
├── test.enc
└── test.key
```

## Definitely ignore the key files with .gitignore:


You can ignore your key files manually with the ```.gitignore``` file or use the ```gitignore``` command. This command will generate or modify your ```.gitignore``` file.
 
```sh
sypher gitignore
```
⚠️ Before commit, be sure to check that the ignore operation is working properly:

## How to make changes on credentials:

Use ```edit``` command to make changes on your credentials.
In this case, sypher will launch an editor(vim by default) with your decrypted credentials.
When you save the changes and close the editor, sypher immediately reads your new credentials and writes it to encrypted credential file in your project. 

For example:
```sh
sypher edit production
```
**To use another editor like Visual Studio Code:**
```sh
EDITOR=code sypher edit production
```

## How to deploy key files:
Instead of using key files in your cloud or development environment, you should set the ```SYPHER_MASTER_KEY``` environment variable.

## Package Usage
In your program, all your need to do is to import sypher.

``` go
package main

import "github.com/sertangulveren/sypher"

func main() {
  // loads sypher/master.enc with sypher/master.key OR SYPHER_MASTER_KEY.
  sypher.Load()
  awsKey := sypher.Get("AWS_SECRET_KEY")
  //...
}
```

Example for production:

``` go
package main

import (
	"github.com/sertangulveren/sypher"
	"os"
)

func main() {
	// APP_ENV=production
	// SYPHER_MASTER_KEY=abcd...
	// production.key application has no production.key
	// loads sypher/production.enc with SYPHER_MASTER_KEY.
	sypher.Load(
	  sypher.Config{Name: os.Getenv("APP_ENV")},
	)

	awsKey := sypher.Get("AWS_SECRET_KEY")
	//...
}
```

### Todos:

* Embed on building.
* Get it ready to be used.
* Missing tests should be done.
* Code quality improvements.
* ...