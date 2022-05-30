package apps

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	"os"
	"runtime"
	"strings"

	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs
	cmd         *cobra.Command
}

type Option func(*App)

type RunFunc func(basename string) error

func NewApp(name string, basename string, opts ...Option) *App {
	app := &App{
		basename: basename,
		name:     name,
	}

	// option  写入
	for _, o := range opts {
		o(app)
	}

	app.buildCommand()

	return app
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           formatBaseName(a.basename),
		Short:         a.basename,
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	cmd.Flags().SortFlags = true

	cliflag.InitFlags(cmd.Flags())

	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}

	var namedFlagSets cliflag.NamedFlagSets

	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()

		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))

	a.cmd = &cmd
}

func formatBaseName(basename string) string {
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

// WithDescription is used to set the description of the application.
func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v \n", color.RedString("Error:"), err)
		os.Exit(0)
	}
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}
