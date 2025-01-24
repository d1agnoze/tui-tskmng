package tskmng

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/d1agnoze/tui-tskmng/internal/parser"
	"github.com/d1agnoze/tui-tskmng/internal/ui"
	"github.com/spf13/cobra"
)

const appName = "tmg"

type config struct {
	cfg             string
	isDefaultConfig bool
	name            string
	tasks           []*parser.Task
}

var (
	App     config
	cfgFile string
)

func init() {
	App = new()

	cobra.OnInitialize(appInit)
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", fmt.Sprintf("config file (default is $HOME/%s/conf.%s)", App.name, App.name))
	root.AddCommand(versionCmd)
}

func new() config { return config{name: appName, isDefaultConfig: true} }

var root = &cobra.Command{
	Use:   appName,
	Short: "Tmg is an over engineered task tracker",
	Long:  `I don't know what to write here. and why are you even reading this? Built with Go so that's cool tho.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []parser.Task{}
		for _, t := range App.tasks {
			if t != nil {
				tasks = append(tasks, *t)
			}
		}
		err := ui.New(tasks).Run()
		cobra.CheckErr(err)
	},
}

func Execute() error {
	return root.Execute()
}

func appInit() {
	path, isDef, err := readConf()
	if err != nil {
		cobra.CheckErr(err)
	}

	tasks, err := parseConf(path)
	if err != nil {
		cobra.CheckErr(err)
	}

	App.cfg = path
	App.isDefaultConfig = isDef
	App.tasks = tasks
}

func readConf() (string, bool, error) {
	var appConfigDir string = ""
	isDefault := cfgFile == ""

	if !isDefault {
		appConfigDir = cfgFile
	} else {
		dir, err := os.UserConfigDir()
		if err != nil {
			return appConfigDir, isDefault, err
		}
		appConfigDir = filepath.Join(dir, appName, fmt.Sprintf("conf.%s", App.name))
	}

	if _, err := os.Stat(appConfigDir); err != nil {
		return appConfigDir, isDefault, err
	}

	return appConfigDir, isDefault, nil
}

func parseConf(path string) ([]*parser.Task, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	p, err := parser.New()
	if err != nil {
		return nil, err
	}

	tasks, err := p.Parse(string(raw))
	if err != nil {
		return nil, err
	}

	return tasks.Tasks, nil
}
