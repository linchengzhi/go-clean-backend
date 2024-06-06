package cerror

// system
var (
	ErrSystem     = NewError(10001, "System error")
	ErrBusy       = NewError(10002, "The system is busy, please try again later")
	ErrFrequently = NewError(10003, "The operation is too frequent, please try again later")
)

// account
var (
	ErrNotAccount = NewError(11001, "The account or password is incorrect")
	ErrPassword   = NewError(11002, "The account or password is incorrect")
	ErrEmailExist = NewError(11003, "The email was exist")
	ErrLogout     = NewError(11004, "Not logged in")
)
