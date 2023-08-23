package domain

type Provider interface {
	LoadTrades() ([]GenericModel, error)
}

type GenericModel interface {
	ToDomainModel() (Trade, error)
}
