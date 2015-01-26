package mogilefs

type ReadonlyError struct{}

func (e ReadonlyError) Error() string {
	return "Cannot create a new file because the connection is read-only."
}
