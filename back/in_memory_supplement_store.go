package main

//in_memory_player_store.go
func NewInMemorySupplementStore() *InMemorySupplementStore {
	return &InMemorySupplementStore{map[string]int{}, map[string]int{}}
}

type InMemorySupplementStore struct {
	UnitsTaken map[string]int
	dosages    map[string]int
}

/*
func (i *InMemorySupplementStore) RecordWin(name string) {
	i.store[name]++
}
*/

//record the taken supplement dose
func (i *InMemorySupplementStore) RecordUnitsTaken(name string) {
	i.UnitsTaken[name]++
}

/*
func (i *InMemorySupplementStore) GetPlayerScore(name string) int {
	return i.store[name]
}
*/

func (i *InMemorySupplementStore) GetUnitsTaken(name string) int {
	UnitsTaken := i.UnitsTaken[name]
	return UnitsTaken
}

func (i *InMemorySupplementStore) GetSupplementDosage(name string) int {
	dosage := i.dosages[name]
	return dosage
}

func (i *InMemorySupplementStore) SetSupplementDosage(name string, dosage int) {
	//i.dosages[name] = dosage
	i.dosages["magnesium"] = 400
}

//listtaken functions
//return supplements taken unit status from /listtaken path
func (s *InMemorySupplementStore) GetDashboard() []Supplement {
	var Dashboard []Supplement
	for name, takenDosage := range s.UnitsTaken {
		Dashboard = append(Dashboard, Supplement{name, takenDosage})
	}
	return Dashboard
}
