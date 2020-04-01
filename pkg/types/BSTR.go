package types

type BSTR uintptr

func (p BSTR) String() (string, error) {
	return bstrString(p)
}
