package common

type UserAgent struct {
	Data string
}

func ParseUserAgentFromString(data string) UserAgent {
	return UserAgent{
		Data: data,
	}
}

func (ua UserAgent) ToString() string {
	return ua.Data
}
