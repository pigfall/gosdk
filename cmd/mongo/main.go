package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pigfall/gosdk/internal/mongo"
)

func main() {
	cmd := &cobra.Command{
		Use:"mongo",
	}

	cmd.AddCommand(
		mongo.ObjectIDCommand(),
	)

	if err := cmd.Execute();err != nil{
		fmt.Fprintf(os.Stderr, "failed: %v",err)
		os.Exit(1)
	}
}
