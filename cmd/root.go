package cmd

import (
	"bufio"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The internal variable cfgFile holds the config file name
var cfgFile string

// LogLevel defines the verbosity of the logging. The following levels are defined:
//  0 = Panic: Only log panic messages (e.g. the universe is blowing up)
//  1 = Fatal: The application can't continue after this error
//  2 = Error: This is an error, but we can continue
//  3 = Warn:  This is a warning, something went wrong, but we can ignore it
//  4 = Info:  Just for you to know that it happens
//  5 = Debug: Lots of internal status messages, allows tracing the application
var LogLevel int

// LogFile is the name of the file to log into. By default, the application logs
// to stdout
var LogFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ppapt",
	Short: "Backend component of the Pen And paper Tools project",
	Long:  `This is the backend and command line component of the Pen And paper Tools project`,
	PersistentPreRun: func(ccmd *cobra.Command, args []string) {
		if LogFile == "" {
			log.SetOutput(os.Stdout)
		} else {
			f, err := os.Create(LogFile)
			if err != nil {
				fmt.Println("Could not create logfile '" + LogFile + "'")
				os.Exit(10)
			}
			w := bufio.NewWriter(f)
			log.SetOutput(w)
		}
		switch LogLevel {
		case 0:
			log.SetLevel(log.PanicLevel)
		case 1:
			log.SetLevel(log.FatalLevel)
		case 2:
			log.SetLevel(log.ErrorLevel)
		case 3:
			log.SetLevel(log.WarnLevel)
		case 4:
			log.SetLevel(log.InfoLevel)
		case 5:
			log.SetLevel(log.DebugLevel)
		default:
			log.SetLevel(log.DebugLevel)
		}
		log.WithFields(log.Fields{
			"LogFile":  LogFile,
			"LogLevel": LogLevel,
		}).Debug("Logging configured")
	},
}

// Start the application => do nothing at the moment, you need to specify "server"
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize the global variable handling
func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "configuration file (default is $HOME/ppapt.yaml)")
	RootCmd.PersistentFlags().IntVarP(&LogLevel, "loglevel", "l", 5, "log level (defaults to 4 (Info))")
	RootCmd.PersistentFlags().StringVarP(&LogFile, "logfile", "L", "", "logfile (defaults to stdout)")

	viper.BindPFlag("loglevel", serverCmd.Flags().Lookup("loglevel"))
	viper.BindPFlag("logfile", serverCmd.Flags().Lookup("logfile"))

	viper.SetDefault("loglevel", 5)
	viper.SetDefault("logfile", "")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		log.Debug("Read config from " + cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		log.Debug("Read config from home directory")
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("ppapt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Error("Can't read config" + err.Error())
		os.Exit(1)
	}
	log.Debug("initConfig finished")
}
