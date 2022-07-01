package flags

import(
	"testing"
	stdflag "flag"
)

func TestParse(t *testing.T){
	handleSetFlagErr := func(err error){
		if err != nil{
			t.Fatal(err)
		}
	}

	paramName := NewParamString("name","","name")
	paramDefaultName := NewParamString("default_name","default_name","default_name")
	params := []Param{paramName,paramDefaultName}
	FlagParams(params...)

	handleSetFlagErr(stdflag.Set("name","foo"))
	Parse()

	if paramName.ValueAfterParsed != "foo"{
		t.Fatal("unexpect")
	}

	if paramDefaultName.ValueAfterParsed != "default_name"{
		t.Fatal("unexpect")
	}
}
