package flags

import(
		stdflag "flag"
)

func FlagParams(params ...Param){
	for _,param:= range params{
		param.SetFlag()
	}
}

func Parse(){
	stdflag.Parse()
}
