package createconfig

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/bketelsen/captainhook/types"
	"github.com/robmerrell/comandante"
)

var filename string

func NewCommand() *comandante.Command {
	return comandante.NewCommand("createconfig", "create a command configuration template", func() error {

		return createCommand()
	})
}

func createCommand() error {

	o := types.Orchestration{}
	a1 := []string{"-l", "-a"}
	s1 := types.Script{Command: "ls", Args: a1}

	s2 := types.Script{Command: "ps"}

	scripts := []types.Script{s1, s2}
	o.Scripts = scripts

	fmt.Printf("Some Config would be spit out here and it would be named %s\n", filename)

	output, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(output)
	if err != nil {
		return err
	}
	return nil
}

func GetFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&filename, "filename", "sample.json", "File to write")
}
