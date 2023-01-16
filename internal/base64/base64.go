package base64

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Base64RawStdEncodingCommand struct{}

type Base64RawStdDecodingCommand struct{}

type Base64StdEncodingCommand struct{}

type Base64StdDecodingCommand struct{}


func Base64Command()*cobra.Command{
	cmd := &cobra.Command{
		Use: "base64",
	}

	rawStdCmd := &cobra.Command{
		Use:"rawStdEncoding",
	}

	rawStdEncodingCommand := &Base64RawStdEncodingCommand{}
	rawStdEncodingCmd := &cobra.Command{
		Use:"encode",
		RunE:rawStdEncodingCommand.Run,
	}

	rawStdDecodingCommand := &Base64RawStdDecodingCommand{}
	rawStdDecodingCmd := &cobra.Command{
		Use: "decode",
		RunE: rawStdDecodingCommand.Run,
	}
	rawStdCmd.AddCommand(
			rawStdEncodingCmd,
			rawStdDecodingCmd,
	)

	stdCmd := &cobra.Command{
		Use:"stdEncoding",
	}
	stdEncoddingCommand := &Base64StdEncodingCommand{}
	stdEncodingCmd := &cobra.Command{
		Use: "encode",
		RunE: stdEncoddingCommand.Run,
	}
	stdDecodingCommand := &Base64StdDecodingCommand{}
	stdDecodingCmd := &cobra.Command{
		Use:"decode",
		RunE: stdDecodingCommand.Run,
	}
	stdCmd.AddCommand(
			stdEncodingCmd,
			stdDecodingCmd,
	)




	cmd.AddCommand(
			rawStdCmd,
			stdCmd,
	)

	return cmd
}



func (c *Base64RawStdEncodingCommand) Run(cmd *cobra.Command,args []string)error{
	bytes,err := io.ReadAll(os.Stdin)
	if err != nil{
		return err
	}
	fmt.Println(b64.RawStdEncoding.EncodeToString(bytes))

	return nil
}

func (c *Base64RawStdDecodingCommand) Run(cmd *cobra.Command, args []string)error{
	bytes,err := io.ReadAll(os.Stdin)
	if err != nil{
		return err
	}
	
	decoded,err := b64.RawStdEncoding.DecodeString(string(bytes))
	if err != nil{
		return err
	}
	fmt.Println(string(decoded))

	return nil
}


func (c *Base64StdEncodingCommand) Run(cmd *cobra.Command,args []string)error{
	bytes,err := io.ReadAll(os.Stdin)
	if err != nil{
		return err
	}
	fmt.Println(b64.StdEncoding.EncodeToString(bytes))

	return nil
}

func (c *Base64StdDecodingCommand) Run(cmd *cobra.Command,args []string) error{
	bytes,err := io.ReadAll(os.Stdin)
	if err != nil{
		return err
	}

	decoded,err := b64.StdEncoding.DecodeString(string(bytes))
	if err != nil{
		return err
	}

	fmt.Println(decoded)

	return nil
}
