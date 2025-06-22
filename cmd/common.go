package cmd

func appendSummary[T any](allSummaries *[]T, summary T, filterFunc func(T) bool) {
	if filterFunc(summary) {
		*allSummaries = append(*allSummaries, summary)
	}
}
