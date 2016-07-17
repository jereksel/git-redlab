package main

type mockedIO struct {
}

func (io mockedIO) ScanString(a *string) (n int, err error) {
	*a = "asd"
	return 0, nil
}

func (io mockedIO) ScanInt(a *int) (n int, err error) {
	*a = 1
	return 0, nil
}

//We don't have to print anything - we're testing input, not output
func (io mockedIO) Print(a ...interface{}) (n int, err error) {
	return 0, nil
}

func (io mockedIO) Printf(format string, a ...interface{}) (n int, err error) {
	return 0, nil
}
