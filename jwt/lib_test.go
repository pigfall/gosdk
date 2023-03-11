package jwt

import (
	"testing"
	"time"

	"github.com/pigfall/gosdk/certs"
)


func TestSignAndValidate(t *testing.T){
	type Claim struct{
		Name string
	}
	privKey,err := certs.RSAGenPrivateKey(certs.PrivateKeyBitSize_2048)
	if err != nil{
		t.Fatal(err)
	}

	now := time.Now()
	token,_,err := SignWithRSA(privKey, now , time.Minute ,&Claim{Name:"pigfall"})
	if err != nil{
		t.Fatal(err)
	}

	signedData,err := ValidateWithRSA(&privKey.PublicKey,token)
	if err != nil{
		panic(err)
	}
	t.Log(signedData.(map[string]interface{})["Name"])

}
