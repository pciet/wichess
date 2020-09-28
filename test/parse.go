package test

type (
	// caseJSON is asserted to the type for the test after calling ParseAllCases
	caseJSON interface{}

	caseJSONParser func(*caseJSON, []byte) error
)

func parseAllCases(parse caseJSONParser, tag string) []caseJSON {
	raws := loadAllCases(tag)
	out := make([]caseJSON, len(raws))
	for i, r := range raws {
		err := parse(&(out[i]), r)
		if err != nil {
			panic(err.Error())
		}
	}

	return out
}
