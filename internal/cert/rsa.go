package cert

import(
	"github.com/spf13/cobra"

	"github.com/pigfall/gosdk/certs"
)

func RSACommand() *cobra.Command{
	cmd := &cobra.Command{
		Use:"rsa",
	}

	privateKeyCommand := &cobra.Command{
		Use:"private-key",
	}

	rsaPrivateKeyCreateCmd := &RSAPrivateKeyCreateCmd{}
	rsaPrivateKeyCreateCommand := &cobra.Command{
		Use:"create",
		RunE: rsaPrivateKeyCreateCmd.Run,
	}
	rsaPrivateKeyCreateCommand.Flags().IntVar(&rsaPrivateKeyCreateCmd.KeyBitSize, "keyBitSize", 2048, "private key bit size, 1024 | 2048 | 3072 | 4096 ")
	rsaPrivateKeyCreateCommand.Flags().StringVar(&rsaPrivateKeyCreateCmd.OutputPath, "outputPath", "priv.key", "the path of generated private key")

	privateKeyCommand.AddCommand(
			rsaPrivateKeyCreateCommand,
	)

	cmd.AddCommand(
		privateKeyCommand,
	)

	return cmd
}

type RSAPrivateKeyCreateCmd struct{
	KeyBitSize int

	OutputPath string
}


func (c *RSAPrivateKeyCreateCmd) Run(cmd *cobra.Command,args []string)error{
	pk,err := certs.RSAGenPrivateKey(certs.PrivateKeyBitSize(c.KeyBitSize))
	if err != nil{
		return err
	}

	return certs.PemSaveRSAPrivateKey(c.OutputPath, pk)
}
