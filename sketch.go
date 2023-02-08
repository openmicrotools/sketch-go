package sketch

// Unwrap will handle returns from calls to functions which can return a value or an error
// if Unwrap encounters and error it will panic and provide the error's contents
func Unwrap[T any](val T, err error) T {

	if err != nil {
		panic(err.Error())
	}
	return val
}

// CustUnwrap will return a function that will handle returns from calls to functions which can return a value or an error
// if CustUnwrap's returned function encounters an error it will return your privide msg
func CustUnwrap[T any](msg string) func(val T, err error) T {
	return func(val T, err error) T {
		if err != nil {
			panic(msg)
		}
		return val
	}
}
