package medium

type HttpRequestMedium struct {
}

func NewHttpRequestMedium() HttpRequestMedium {
	return HttpRequestMedium{}
}

func (m *HttpRequestMedium) Send() string {
	return "Demo response"
}
