package common

// 过滤函数, 入参是structName, 返回值是false表示不过滤该表, 返回值是true表示过滤该表
type FiltHook func(string) bool

func (h *FiltHook) Fill() {
	if *h == nil {
		*h = func(string) bool {
			return false
		}
	}
}

func (h *FiltHook) LoadStruct(structNames ...string) {
	allowMap := make(map[string]struct{}, len(structNames))

	for i := len(structNames) - 1; i > -1; i-- {
		allowMap[structNames[i]] = struct{}{}
	}

	*h = func(structNames string) bool {
		_, ok := allowMap[structNames]
		return ok
	}
}

func (h *FiltHook) FiltStruct(structNames ...string) {
	refuseMap := make(map[string]struct{}, len(structNames))

	for i := len(structNames) - 1; i > -1; i-- {
		refuseMap[structNames[i]] = struct{}{}
	}

	*h = func(structNames string) bool {
		_, ok := refuseMap[structNames]
		return !ok
	}
}