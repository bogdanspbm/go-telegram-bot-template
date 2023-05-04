package bot

type storage struct {
	balance map[int64]int
}

func createStorage() *storage {
	stg := &storage{balance: make(map[int64]int)}
	return stg
}

func (stg *storage) getBalance(chatID int64) int {
	value, ok := stg.balance[chatID]

	if !ok {
		stg.balance[chatID] = 100
		return 100
	}

	return value
}

func (stg *storage) setBalance(chatID int64, balance int) {
	stg.balance[chatID] = balance
}

func (stg *storage) addBalance(chatID int64, delta int) {
	value, ok := stg.balance[chatID]

	if !ok {
		value = 0
	}

	stg.balance[chatID] = value + delta
}
