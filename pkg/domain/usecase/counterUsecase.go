package usecase

type CounterUseCase interface {
	CountWord(text string) (map[string]int, error)
}
