package flags

import(
	"os"
	"testing"
	stdflag "flag"
)

func TestParse(t *testing.T){
	handleSetFlagErr := func(err error){
		if err != nil{
			t.Fatal(err)
		}
	}

	ResetForTesting(nil)

	paramName := NewParamString("name","","name")
	paramDefaultName := NewParamString("default_name","default_name","default_name")
	params := []Param{paramName,paramDefaultName}
	FlagParams(params...)

	handleSetFlagErr(stdflag.Set("name","foo"))
	handleSetFlagErr(ParseAndValidate(params))

	if paramName.ValueAfterParsed != "foo"{
		t.Fatal("unexpect")
	}

	if paramDefaultName.ValueAfterParsed != "default_name"{
		t.Fatal("unexpect")
	}
}

func TestParamValidator(t *testing.T){
	handleSetFlagErr := func(err error){
		if err != nil{
			t.Fatal(err)
		}
	}

	ResetForTesting(nil)
	
	paramName := NewParamString("name","","input user name",ParamStringNotEmpty())

	params := []Param{paramName}
	FlagParams(params...)

	handleSetFlagErr(stdflag.Set("name",""))

	err := ParseAndValidate(params)
	if err == nil{
		t.Fatal("unexptect")
	}
}


func ResetForTesting(usage func()) {
	stdflag.CommandLine = stdflag.NewFlagSet(os.Args[0], stdflag.ContinueOnError)
	stdflag.CommandLine.Usage = stdflag.Usage
	stdflag.Usage = usage
}
