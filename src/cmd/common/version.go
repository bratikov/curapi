package common

import (
	"currency/internal/version"
	"io"
	"os"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
)

var versionTemplate = `Currency server
Version:      {{.Version}}
Built:        {{.BuildDate}}
Go version:   {{.GoVersion}}
OS/Arch:      {{.Os}}/{{.Arch}}
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetPrint(os.Stdout); err != nil {
			panic("Cant generate version")
		}
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func GetPrint(wr io.Writer) error {
	tmpl, err := template.New("").Parse(versionTemplate)
	if err != nil {
		return err
	}

	v := struct {
		Version   string
		Codename  string
		GoVersion string
		BuildDate string
		Os        string
		Arch      string
	}{
		Version:   version.Version,
		GoVersion: runtime.Version(),
		BuildDate: version.Date,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}

	return tmpl.Execute(wr, v)
}
