package valuer

type Valuer struct {
	defaultValue interface{}
	candidates   []interface{}
}

func Value(defaultValue ...interface{}) Valuer {
	v := Valuer{}
	if len(defaultValue) > 0 {
		return v.Default(defaultValue[0])
	}
	return v
}

func (v Valuer) Default(value interface{}) Valuer {
	v.defaultValue = value
	return v
}

func (v Valuer) Try(value ...interface{}) Valuer {
	v.candidates = append(v.candidates, value...)
	return v
}

func (v Valuer) String() string {
	for _, value := range v.candidates {
		if s := InterfaceToString(value); s != "" {
			return s
		}
	}
	if v.defaultValue != nil {
		return InterfaceToString(v.defaultValue)
	}
	return ""
}

func (v Valuer) StringSlice() []string {
	for _, value := range v.candidates {
		if s := InterfaceToStringSlice(value); len(s) > 0 {
			return s
		}
	}
	if v.defaultValue != nil {
		return InterfaceToStringSlice(v.defaultValue)
	}
	return []string{}
}

func (v Valuer) Int() int {
	for _, value := range v.candidates {
		if i := InterfaceToInt(value); i != 0 {
			return i
		}
	}
	if v.defaultValue != nil {
		return InterfaceToInt(v.defaultValue)
	}
	return 0
}

func (v Valuer) Int64() int64 {
	for _, value := range v.candidates {
		if i := InterfaceToInt64(value); i != 0 {
			return i
		}
	}
	if v.defaultValue != nil {
		return InterfaceToInt64(v.defaultValue)
	}
	return 0
}

func (v Valuer) Float64() float64 {
	for _, value := range v.candidates {
		if i := InterfaceToFloat64(value); i != 0 {
			return i
		}
	}
	if v.defaultValue != nil {
		return InterfaceToFloat64(v.defaultValue)
	}
	return 0
}

func (v Valuer) Bool() bool {
	for _, value := range v.candidates {
		if b := InterfaceToBool(value); b {
			return b
		}
	}
	return false
}
