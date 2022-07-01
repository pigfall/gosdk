package flags

import(
		stdflag "flag"
)

type ParamDescBase struct{
	Name string
	Usage string
}

type ParamString struct{
	ParamDescBase
	ValueAfterParsed string
	DefaultValue string
}

func (this *ParamString) SetFlag(){
	stdflag.StringVar(&this.ValueAfterParsed,this.Name,this.DefaultValue,this.Usage)
}

func NewParamString(name string,defaultValue string,usage string)*ParamString{
	return &ParamString{
		ParamDescBase:ParamDescBase{
			Name :name,
			Usage :usage,
		},
		DefaultValue:defaultValue,
	}
}

