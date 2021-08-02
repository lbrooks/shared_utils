package doublekey

type StringString struct {
	Data map[string]map[string]struct{}
	initialCapacity int
}

func NewStringString() *StringString {
	return &StringString{
		Data: make(map[string]map[string]struct{}),
		initialCapacity: -1,
	}
}

func NewStringStringOfSize(size int) *StringString {
	return &StringString{
		Data: make(map[string]map[string]struct{}, size),
		initialCapacity: size,
	}
}

func (s *StringString) createSubMap(key string){
	if s.initialCapacity < 0 {
		s.Data[key] = make(map[string]struct{})
	} else {
		s.Data[key] = make(map[string]struct{}, s.initialCapacity)
	}
}

func (s *StringString) Add(key1, key2 string) {
	if v1, ok := s.Data[key1]; ok {
		v1[key2] = struct{}{}
	} else {
		s.createSubMap(key1)
		s.Data[key1][key2] = struct{}{}
	}
}

func (s *StringString) AddSlice(key1 string, values []string) {
	if _, ok := s.Data[key1]; !ok {
		s.createSubMap(key1)
	}
	for _, v := range values {
		s.Data[key1][v] = struct{}{}
	}
}

func (s *StringString) Remove(key1, key2 string) bool {
	if v1, o1 := s.Data[key1]; o1 {
		if _, o2 := s.Data[key1]; o2 {
			delete(v1, key2)
			if len(v1) == 0 {
				delete(s.Data, key1)
			}
			return true
		}
	}
	return false
}