package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/EconomistDigitalSolutions/fcreds/awssecrets"
	"github.com/alecthomas/kingpin"
)

var (
	app                  = kingpin.New("awssecrets", "A command line tool to get AWS secrets.")
	region               = app.Flag("region", "Configure the AWS region").Required().Short('r').String()
	profile              = app.Flag("profile", "Configure the AWS profile").Short('p').String()
	cmdExecute           = app.Command("exec", "Execute a command with all secrets loaded as environment variables.")
	cmdExecuteSecretName = awssecrets.SecretList(cmdExecute.Flag("secret-name", "Secret name").Required().Short('n'))
	cmdExecuteCommand    = cmdExecute.Arg("command", "The command to execute.").Required().Strings()
)

func main() {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case cmdExecute.FullCommand():
		args := []string(*cmdExecuteCommand)
		commandPath, err := exec.LookPath(args[0])
		if err != nil {
			log.Fatal(err)
		}

		svc := awssecrets.GetSecretManager(profile, region)

		awssecrets.SetSecretEnv(svc, *cmdExecuteSecretName)

		err = syscall.Exec(commandPath, args, os.Environ())
		if err != nil {
			log.Fatal(err)
		}
	}
}
