package flags

type Param interface{
	SetFlag()
	Validate()error
}
