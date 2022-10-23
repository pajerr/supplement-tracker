package main

//in_memory_player_store.go
func NewInMemorySupplementStore() *InMemorySupplementStore {
	return &InMemorySupplementStore{map[string]int{}, map[string]int{}}
}

type InMemorySupplementStore struct {
	takenSupplements map[string]int
	dosages          map[string]int
}

/*
func (i *InMemorySupplementStore) RecordWin(name string) {
	i.store[name]++
}
*/

//record the taken supplement dose
func (i *InMemorySupplementStore) RecordTakenSupplement(name string) {
	i.takenSupplements[name]++
}

/*
func (i *InMemorySupplementStore) GetPlayerScore(name string) int {
	return i.store[name]
}
*/

func (i *InMemorySupplementStore) GetTakenSupplement(name string) int {
	takenSupplementdosages := i.takenSupplements[name]
	return takenSupplementdosages
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
func (s *InMemorySupplementStore) GetAllSupplementsStatus() []Supplement {
	var supplementsStatus []Supplement
	for name, takenDosage := range s.takenSupplements {
		supplementsStatus = append(supplementsStatus, Supplement{name, takenDosage})
	}
	return supplementsStatus
}
