package main

func Opponent(of, white, black string) string {
	if of == white {
		return black
	} else if of == black {
		return white
	}
	Panic(of, "not white", white, "or black", black)
	return of
}
