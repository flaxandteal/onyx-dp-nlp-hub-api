package params

type CategoryParams struct {
	Query string `url:"query,omitempty"`
}

func GetCategoryParams(q string) *CategoryParams {
	return &CategoryParams{
		Query: q,
	}
}
