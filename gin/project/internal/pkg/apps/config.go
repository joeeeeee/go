package apps

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)


var cfgFile string
const configFlagName = "config"


func addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configFlagName))

	viper.AutomaticEnv()

	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "-"))

	viper.AddConfigPath("./configs")

	viper.SetConfigName(basename)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : failed to read configuration file (%s)", err)
		os.Exit(1)
	}
}