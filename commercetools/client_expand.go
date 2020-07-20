package commercetools

import "net/url"

type RequestOption func(*url.Values)

func WithReferenceExpansion(references ...string) RequestOption {
	return func(v *url.Values) {
		for _, ref := range references {
			v.Add("expand", ref)
		}
	}
}
