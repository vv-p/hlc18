package main

const (
	sexSliceDefaultSize = 6000 // half of total 10000 / 2 + 1000
)

type (
	IndexSex struct {
		f []*Account
		m []*Account
	}
)

func MakeIndexSex() *IndexSex {
	return &IndexSex{
		f: make([]*Account, 0, sexSliceDefaultSize),
		m: make([]*Account, 0, sexSliceDefaultSize),
	}
}

func (s *IndexSex) Add(a *Account) {
	if a.Sex == "f" {
		s.f = append(s.f, a)
	} else {
		s.m = append(s.m, a)
	}
}

func (s *IndexSex) Get(sex string, limit uint64) []*Account {
	if sex == "f" {
		return s.f[:limit]
	}
	return s.m[:limit]
}
