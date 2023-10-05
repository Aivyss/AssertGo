package errors

type MsgAndArgsErr struct {
	Err        error
	MsgAndArgs []interface{}
}

func (e MsgAndArgsErr) Error() string {
	return e.Err.Error()
}
