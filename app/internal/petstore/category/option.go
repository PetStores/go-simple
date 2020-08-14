package category

type CategoryOptions struct {
	queryParams map[string]interface{}
}

type CategoryOption func(*CategoryOptions)

func WithID(id int64) CategoryOption {
	return func(opts *CategoryOptions) {
		opts.queryParams["id"] = id
	}
}

func WithName(name string) CategoryOption {
	return func(opts *CategoryOptions) {
		opts.queryParams["name"] = name
	}
}
