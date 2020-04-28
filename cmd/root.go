package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (`| |        | |                    | | | |    
| | ___   _| |__   ___  __ _ _   _| |_| |__  
| |/ / | | | '_ \ / _ \/ _` | | | | __| '_ \ 
|   <| |_| | |_) |  __/ (_| | |_| | |_| | | |
|_|\_\\__,_|_.__/ \___|\__,_|\__,_|\__|_| |_|

`
)

var (
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "kubeauth",
		Short: "A OpenId connect authentication helper for Kubernetes",
		Long: fmt.Sprintf(`%s
kubeauth is a authentication helper for Kubernetes (based on dexter) that does the heavy
lifting for SSO (Single Sign On) for Kubernetes.`, BANNER),
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// Execute executes the root command.
func Execute() {
	if Verbose {
		log.SetLevel(log.DebugLevel)
	}

	rootCmd.Execute()
}
