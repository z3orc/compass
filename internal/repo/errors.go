package repo

type InvalidFlavourError struct{}

func (m *InvalidFlavourError) Error() string {
	return "the given flavour is not valid"
}
