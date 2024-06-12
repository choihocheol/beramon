package store

var GlobalState GlobalStateType

func init() {
	GlobalState = GlobalStateType{
		CL: CLType{
			Status: true,
		},
		EL: ELType{
			Status: true,
		},
	}
}
