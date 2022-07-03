package flags

import(
	"fmt"
		stdflag "flag"
)

type ParamDescBase struct{
	Name string
	Usage string
	NonEmpty bool
	Validators []func(v interface{})error
}

func (this ParamDescBase) validate(v interface{})error{
	for _,validator := range this.Validators{
		err := validator(v)
		if err != nil{
			return fmt.Errorf("value of flag %s is <%v>,invalid: %w",this.Name,v,err)
		}
	}

	return nil
}

// { ParamString

type ParamString struct{
	ParamDescBase
	ValueAfterParsed string
	DefaultValue string
}

// {{ impl Param interface
func (this *ParamString) SetFlag(){
	stdflag.StringVar(&this.ValueAfterParsed,this.Name,this.DefaultValue,this.Usage)
}

func (this *ParamString) Validate()error{
	return this.validate(this.ValueAfterParsed)
}
// }}

type ParamStringOption func(param *ParamString)

func ParamStringNotEmpty()ParamStringOption{
	return ParamStringValidator(
			func(param string)error{
				if len(param) == 0{
					return fmt.Errorf("string is empty")
				}
				return nil
			},
	)
}

func ParamStringValidator(check func(param string)error)ParamStringOption{
	return func(param *ParamString){
		if param.Validators  == nil{
			param.Validators = make([]func(v interface{})error,0,1)
		}
		param.Validators = append(param.Validators,func(paramV interface{})error{
			return check(paramV.(string))
		})
	}
}


func NewParamString(name string,defaultValue string,usage string,opts ...ParamStringOption)*ParamString{
	ret := &ParamString{
		ParamDescBase:ParamDescBase{
			Name :name,
			Usage :usage,
		},
		DefaultValue:defaultValue,
	}
	for _,opt := range opts{
		opt(ret)
	}

	return ret
}

// }


// { ParamInt
type ParamInt struct {
	ParamDescBase
	ValueAfterParsed int
	DefaultValue int
}

type ParamIntOption func(*ParamInt)

func ParamIntValidator(check func(param int)error)ParamIntOption{
	return func(param *ParamInt){
		if param.Validators  == nil{
			param.Validators = make([]func(v interface{})error,0,1)
		}
		param.Validators = append(param.Validators,func(paramV interface{})error{
			return check(paramV.(int))
		})
	}
}

func ParamIntNotZero()ParamIntOption{
	return ParamIntValidator(
		func(p int)error{
			if p == 0{
				return fmt.Errorf("int is zero")
			}
			return nil
		},
	)
}

func NewParamInt(name string,dftValue int, usage string,opts ...ParamIntOption)*ParamInt{
	ret:= &ParamInt{
		ParamDescBase:ParamDescBase{
			Name :name,
			Usage :usage,
		},
		DefaultValue:dftValue,
	}

	for _,opt := range opts{
		opt(ret)
	}

	return ret
}


func (this *ParamInt) SetFlag(){
	stdflag.IntVar(&this.ValueAfterParsed,this.Name,this.DefaultValue,this.Usage)
}

func (this *ParamInt) Validate()error{
	return this.validate(this.ValueAfterParsed)
}
// }

