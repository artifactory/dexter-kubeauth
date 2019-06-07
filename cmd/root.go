package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	BANNER = `    .___               __                
  __| _/____ ___  ____/  |_  ___________ 
 / __ |/ __ \\  \/  /\   __\/ __ \_  __ \
/ /_/ \  ___/ >    <  |  | \  ___/|  | \/
\____ |\___  >__/\_ \ |__|  \___  >__|   
     \/    \/      \/           \/       
| |  / /
| | / /
|  \ \ UBEAUTH
| | \ \
| |  \ \

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
