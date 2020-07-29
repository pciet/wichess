package rules

func Normalize(s *Square) {
	s.Swaps = false
	s.Detonates = false
	s.Guards = false
	s.Fortified = false
	s.Locks = false
	s.Rallies = false
	s.Reveals = false
	s.Tense = false
	s.Fantasy = false
	s.Keep = false
	s.Protective = false
	s.Extricates = false
	s.Orders = false

	// the Normalize, MustEnd, and Ghost bools are left true
}
