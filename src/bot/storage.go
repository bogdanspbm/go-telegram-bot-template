package bot

type storage struct {
	balance map[int64]int
	bet     map[int64]int
	mutex   chan struct{}
}

func (stg *storage) lock() {
	stg.mutex <- struct {
	}{}
}

func (stg *storage) unlock() {
	<-stg.mutex
}

func createStorage() *storage {
	stg := &storage{balance: make(map[int64]int), bet: make(map[int64]int), mutex: make(chan struct{}, 1)}
	return stg
}

func (stg *storage) getBalance(chatID int64) int {
	stg.lock()
	defer stg.unlock()

	value, ok := stg.balance[chatID]

	if !ok {
		stg.balance[chatID] = 100
		return 100
	}

	return value
}

func (stg *storage) setBalance(chatID int64, balance int) {
	stg.lock()
	defer stg.unlock()

	stg.balance[chatID] = balance
}

func (stg *storage) addBalance(chatID int64, delta int) {
	stg.lock()
	defer stg.unlock()

	value, ok := stg.balance[chatID]

	if !ok {
		value = 0
	}

	if value+delta < 0 {
		return
	}

	stg.balance[chatID] = value + delta
}

func (stg *storage) getBet(chatID int64) int {
	stg.lock()
	defer stg.unlock()

	value, ok := stg.bet[chatID]

	if !ok {
		stg.bet[chatID] = 1
		return 1
	}

	return value
}

func (stg *storage) setBet(chatID int64, bet int) {
	stg.lock()
	defer stg.unlock()

	stg.bet[chatID] = bet
}
