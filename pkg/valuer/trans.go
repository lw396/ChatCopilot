package valuer

import (
	"fmt"
	"strconv"
	"strings"
)

func InterfaceToBool(v interface{}) bool {
	switch v := v.(type) {
	case int64:
		return v > 0
	case int32:
		return v > 0
	case int:
		return v > 0
	case uint:
		return v > 0
	case uint32:
		return v > 0
	case uint64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	case string:
		b, _ := strconv.ParseBool(v)
		return b
	case bool:
		return v
	default:
		return false
	}
}

func InterfaceToInt(v interface{}) int {
	switch v := v.(type) {
	case int64:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
	case uint:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		i, _ := strconv.ParseInt(v, 10, 64)
		return int(i)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func InterfaceToUint64(v interface{}) uint64 {
	switch v := v.(type) {
	case int64:
		return uint64(v)
	case int32:
		return uint64(v)
	case int:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case string:
		i, _ := strconv.ParseUint(v, 10, 64)
		return i
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func InterfaceToInt64(v interface{}) int64 {
	switch v := v.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case int:
		return int64(v)
	case uint:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		i, _ := strconv.ParseInt(v, 10, 64)
		return i
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func InterfaceToFloat64(v interface{}) float64 {
	switch v := v.(type) {
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func InterfaceToStringSlice(v interface{}) []string {
	switch v := v.(type) {
	case []string:
		return v
	default:
		s := InterfaceToString(v)
		if s != "" {
			return strings.Split(s, ",")
		}
		return []string{}
	}
}

func InterfaceToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
