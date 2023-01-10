package main

import(
	"os"
	"fmt"
		"github.com/spf13/cobra"

		"github.com/pigfall/gosdk/internal/cert"
)

func main(){
	rootCmd := cobra.Command{
		Use: "cert",
	}

	rootCmd.AddCommand(
		cert.CaCommand(),
		cert.CSRCommand(),
		cert.RSACommand(),
	)

	if err := rootCmd.Execute();err != nil{
		fmt.Fprintf(os.Stderr, "failed: %v\n",err)
	}
}
