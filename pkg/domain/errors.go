package domain

var (
	ErrInternalError = InternalError{"internal error"}

	ErrNotFound = NotFoundError{"not found"}

	ErrInvalidEntity = InvalidEntityError{"invalid entity"}

	ErrInvalidPasswordLength = InvalidLengthError{"invalid password length"}

	ErrAlreadyInUse = AlreadyInUseError{"already in use"}

	ErrInvalidPassword = InvalidPasswordError{"invalid password"}

	ErrInvalidAccessToken = InvalidAccessTokenError{"invalid access token"}
)

type InternalError struct {
	s string
}

func (e InternalError) Error() string {
	return e.s
}

type NotFoundError struct {
	s string
}

func (e NotFoundError) Error() string {
	return e.s
}

type InvalidEntityError struct {
	s string
}

func (e InvalidEntityError) Error() string {
	return e.s
}

type InvalidPasswordError struct {
	s string
}

func (e InvalidPasswordError) Error() string {
	return e.s
}

type InvalidLengthError struct {
	s string
}

func (e InvalidLengthError) Error() string {
	return e.s
}

type AlreadyInUseError struct {
	s string
}

func (e AlreadyInUseError) Error() string {
	return e.s
}

type InvalidAccessTokenError struct {
	s string
}

func (e InvalidAccessTokenError) Error() string {
	return e.s
}
