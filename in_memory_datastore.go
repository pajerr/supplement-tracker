package main

//in_memory_player_store.go
func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{map[string]int{}}
}

type InMemoryDataStore struct {
	store map[string]int
}

func (i *InMemoryDataStore) StoreTakenDosage(name string, dosage int) {
	i.store[name] = dosage
}

/*
func (i *InMemoryDataStore) GetPlayerScore(name string) int {
	return i.store[name]
}
*/

func (i *InMemoryDataStore) GetSupplementDosage(name string) int {
	return 500
}
