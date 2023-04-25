package skiplist

type SkiplistOption func(opts *skiplistOption)

type skiplistOption struct {
	maxLevel    int
	probability float64
}

func WithMaxLevel(maxLevel int) SkiplistOption {
	return func(opts *skiplistOption) {
		opts.maxLevel = maxLevel
	}
}

func WithProbability(probability float64) SkiplistOption {
	return func(opts *skiplistOption) {
		opts.probability = probability
	}
}
