package flags

import(
		stdflag "flag"
)

func FlagParams(params ...Param){
	for _,param:= range params{
		param.SetFlag()
	}
}

func ParseAndValidate(params []Param)error{
	stdflag.Parse()

	for _,p := range params {
		err := p.Validate()
		if err != nil{
			return err
		}
	}

	return nil
}
