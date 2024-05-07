package fallbacks

type FallbackResult struct {
	Location string
}

func (fb *FallbackResult) Url() string {
	return fb.Location
}

type Fallback interface {
	Search(terms []string) (*FallbackResult, error)
}