package cert

import (
	"fmt"
	"crypto/x509"
	"crypto/x509/pkix"
	crand "crypto/rand"

	"github.com/spf13/cobra"

	"github.com/pigfall/gosdk/certs"
)

// Certificate Signing Request
func  CSRCommand() *cobra.Command{
	cmd := &cobra.Command{
		Use: "csr",
	}

	csrCreateCmd := &CSRCreateCmd{}
	csrCreateCommand := &cobra.Command{
		Use:"create",
		RunE: csrCreateCmd.Run,
	}
	csrCreateCommand.Flags().StringVar(&csrCreateCmd.SubjectCommonName, "subjectCommonName", "","Common Name")
	if err:=csrCreateCommand.MarkFlagRequired("subjectCommonName");err != nil{
		panic(err)
	}
	csrCreateCommand.Flags().StringVar(&csrCreateCmd.PrivKeyPath, "privateKeyPath", "","RSA Private Key Path")
	if err:=csrCreateCommand.MarkFlagRequired("privateKeyPath");err != nil{
		panic(err)
	}
	csrCreateCommand.Flags().StringVar(&csrCreateCmd.OutputCSRPath, "outputCSRPath", "csr.pem","the path of generated csr")

	cmd.AddCommand(
			csrCreateCommand,
	)

	return cmd
}

type CSRCreateCmd struct{
	SubjectCommonName string
	PrivKeyPath string

	OutputCSRPath string
}

func (c *CSRCreateCmd) Run(cmd *cobra.Command, args []string) error{
	csrTpl := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: c.SubjectCommonName,
		},
	}

	privKey,err := certs.PemLoadRSAPrivateKey(c.PrivKeyPath)
	if err != nil{
		return fmt.Errorf("load private key from file `%s` error: %w", c.PrivKeyPath,err)
	}

	csrDERBytes ,err := x509.CreateCertificateRequest(crand.Reader, &csrTpl, privKey)
	if err != nil{
		return err
	}
	csr,err := x509.ParseCertificateRequest(csrDERBytes)
	if err != nil{
		return err
	}
	return certs.PemSaveCSRToFile(csr, c.OutputCSRPath)
}


