package cert

import (
	"fmt"
	"strconv"
	"time"
	"crypto/x509/pkix"
	"crypto/x509"

	"github.com/spf13/cobra"

	"github.com/pigfall/gosdk/certs"
	pkgTime "github.com/pigfall/gosdk/time"
)

// Command for creating ca cert.
type CaCreateCommand struct{
	SubjectCommonName string
	ValidNotBefore string
	ValidDuration string
	OutputCertPath string
	OutputPrivKeyPath string

	validNotBefore time.Time
	validDuration time.Duration
}

type CaValidateCertCommand struct{
	CaPath string
	CertPath string
}

func CaCommand() *cobra.Command{
	cmd := &cobra.Command{
		Use: "ca",
	}

	caCreateCmd := &CaCreateCommand{}
	caCreateCobraCmd := &cobra.Command{
		Use:"create",
		RunE: func(cmd *cobra.Command, args []string) error{
				return caCreateCmd.Run(cmd,args)
		},
	}
	caCreateCobraCmd.Flags().StringVar(&caCreateCmd.SubjectCommonName,"subjectCommonName", "","Common Name")
	if err :=caCreateCobraCmd.MarkFlagRequired("subjectCommonName");err != nil{
		panic(err)
	}
	caCreateCobraCmd.Flags().StringVar(&caCreateCmd.ValidNotBefore, "validNotBefore", "now", "Valid Not Before")
	caCreateCobraCmd.Flags().StringVar(&caCreateCmd.ValidDuration, "validDuration", "1m", "Valid Duration")
	caCreateCobraCmd.Flags().StringVar(&caCreateCmd.OutputCertPath, "outputCertPath", "./ca.crt", "The path of the generated ca certificate")
	caCreateCobraCmd.Flags().StringVar(&caCreateCmd.OutputPrivKeyPath, "outputPrivKeyPath", "./ca.key", "The path of the generated ca private key")

	caValidateCertCmd := &CaValidateCertCommand{}
	caValidateCertCommand := &cobra.Command{
		Use: "validate",
		RunE: caValidateCertCmd.Run,
	}
	caValidateCertCommand.Flags().StringVar(&caValidateCertCmd.CaPath,"caPath","","ca cert path")
	caValidateCertCommand.Flags().StringVar(&caValidateCertCmd.CertPath,"certPath","","cert path")
	if err :=caValidateCertCommand.MarkFlagRequired("caPath");err != nil{
		panic(err)
	}
	if err :=caValidateCertCommand.MarkFlagRequired("certPath");err != nil{
		panic(err)
	}


	cmd.AddCommand(
			caCreateCobraCmd,
			caValidateCertCommand,
	)

	return cmd
}

func (c *CaCreateCommand) Run(cmd *cobra.Command,args []string)error{
	if err:=c.init();err != nil{
		return err
	}
	caCrtTpl := certs.NewX509CaCrtTpl(
		pkix.Name{
			CommonName: c.SubjectCommonName,
		},
		c.validNotBefore,
		c.validDuration,
		nil,
	)

	privKey,err := certs.RSAGenPrivateKey(certs.PrivateKeyBitSize_2048)
	if err != nil{
		return err
	}

	caCrt,err := certs.SignSelf(caCrtTpl,privKey)
	if err != nil{
		return err
	}
	if err :=certs.PemX509Save(c.OutputCertPath, caCrt.Raw);err != nil{
		return err
	}
	if err:=certs.PemSaveRSAPrivateKey(c.OutputPrivKeyPath, privKey);err != nil{
		return err
	}

	fmt.Printf("ca certificate generated at: %s\n", c.OutputCertPath)
	fmt.Printf("ca private key  generated at: %s\n", c.OutputPrivKeyPath)

	return nil
}

func (c *CaCreateCommand) init()error{
	// Parse validNotBefore
	if c.ValidNotBefore == "now"{
		c.validNotBefore = time.Now()
	}else{
		var err error
		if c.validNotBefore,err =pkgTime.ParseFromYYYY_MM_DD_HH_MM_SS(c.ValidNotBefore);err != nil{
			return err
		}
	}

	unit := string(c.ValidDuration[len(c.ValidDuration)-1])
	value,err := strconv.ParseInt(c.ValidDuration[:len(c.ValidDuration)-1],10,64)
	if err != nil{
		return err
	}
	switch unit{
	case "y":
		c.validDuration = time.Duration(value) * time.Hour*24*365
	case "d":
		c.validDuration = time.Duration(value) * time.Hour*24
	case "m":
		c.validDuration = time.Duration(value) * time.Minute
	case "s":
		c.validDuration = time.Duration(value) * time.Second
	}

	return nil
}

func (c *CaValidateCertCommand) Run(cmd *cobra.Command,args []string)error{
	caCrt,err := certs.PemX509FromFile(c.CaPath)
	if err != nil {
		return fmt.Errorf("parse ca cert from file `%s` error: %w", c.CaPath,err)
	}

	cert,err := certs.PemX509FromFile(c.CertPath)
	if err != nil{
		return fmt.Errorf("parse cert from file `%s` error: %w",c.CertPath,err )
	}

	caPool := x509.NewCertPool()
	caPool.AddCert(caCrt)
	_,err = cert.Verify(x509.VerifyOptions{
		Roots: caPool,
	})
	if err != nil{
		return fmt.Errorf("the cert `%s` is not signed by ca `%s`",c.CaPath,c.CertPath)
	}

	return nil
}

