package exception

type AdminCreating struct {
}

func (e *AdminCreating) Error() string {
	return "creating admin failed"
}
