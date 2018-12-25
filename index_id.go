package main

type (
	IndexId struct {
		accounts map[AccountId]*Account
	}
)

func MakeIndexId() *IndexId {
	return &IndexId{accounts: map[AccountId]*Account{}}
}

func (ii *IndexId) Add(account *Account) {
	ii.accounts[account.Id] = account
}

func (ii *IndexId) Get(id AccountId) *Account {
	return ii.accounts[id]
}

func (ii *IndexId) Len() int {
	return len(ii.accounts)
}
