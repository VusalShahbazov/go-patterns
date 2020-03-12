package Singleton

type singleton struct {
	Number int
}

func (s *singleton) SetNumber(n int)  {
	s.Number = n
}
var instance *singleton
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{
			Number:1,
		}
	}
	return instance
}
