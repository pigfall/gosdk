package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"

	"github.com/pigfall/gosdk/internal/base64"
)

func main() {
	rootCmd := cobra.Command{
		Use:"ak",
		Short: "A cli tools hub",
	}

	rootCmd.AddCommand(
			base64.Base64Command(),
	)


	if err := rootCmd.Execute();err != nil{
		fmt.Fprintf(os.Stderr, "failed: %v",err)
		os.Exit(1)
	}
}
