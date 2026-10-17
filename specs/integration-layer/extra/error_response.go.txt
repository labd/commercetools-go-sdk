package integrationlayer

// Not generated: copied in by the generate-integrationlayer task. The
// management API answers an error as `{ "error", "error_description" }`, which
// the generator gives no Error method, since it has no `message`.

func (obj ErrorResponse) Error() string {
	if obj.Message == nil {
		return "unknown ErrorResponse: failed to parse error response"
	}
	if obj.ErrorDescription != nil && *obj.ErrorDescription != "" {
		return *obj.Message + ": " + *obj.ErrorDescription
	}
	return *obj.Message
}
