package data

type Store struct {
	M map[string]int
}

type KeyValue struct {
	K Key
	V Val
}

type Key struct {
	String string
}

type Val struct {
	Value int
}


