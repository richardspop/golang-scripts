package main

type rec struct {
	id uint64
}

func main() {
	recs := []*rec{
		&rec{
			id: 1,
		},
		&rec{
			id: 2,
		},
		&rec{
			id: 3,
		},
	}
	idmap := make(map[uint64]struct{},0)
	for _, re := range recs {
		idmap[re.id] = struct {}{}
	}
}
