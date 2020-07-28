package test

type (
	// CaseJSON is asserted to the type for the test after calling ParseAllCases
	CaseJSON interface{}

	CaseJSONParser func(*CaseJSON, []byte) error
)

func ParseAllCases(parse CaseJSONParser, tag string) []CaseJSON {
	raws := LoadAllCases(tag)
	out := make([]CaseJSON, len(raws))
	for i, r := range raws {
		err := parse(&(out[i]), r)
		if err != nil {
			panic(err.Error())
		}
	}

	return out
}
