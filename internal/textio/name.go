package textio

import (
	"runtime"
	"strings"
)

func Logger(skip int) string {
	sb := strings.Builder{}
	sb.WriteString("[")
	sb.WriteString(FileName(skip + 1))
	sb.WriteString("]")
	return sb.String()
}

func FileName(skip int) string {
	_, file, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	i := strings.LastIndex(file, "/")
	if i < 0 {
		i = strings.LastIndex(file, "\\")
		if i < 0 {
			return file
		}
	}
	nm := file[i+1:]
	if len(nm) == 0 {
		return file
	}
	i = strings.LastIndex(nm, ".")
	if i <= 0 {
		return nm
	}
	return nm[:i]
}

func FuncName(skip int, lowerfirst bool) string {
	fn, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic("cannot get func name of caller")
	}
	f := runtime.FuncForPC(fn)
	if f == nil {
		panic("invalid func pc")
	}
	p := f.Name()
	i := strings.LastIndex(p, ".")
	if i < 0 {
		panic("func name of caller '" + p + " has no '.'")
	}
	p = p[i+1:]
	if len(p) <= 1 {
		panic("func name of caller '" + p + " too short'")
	}
	if lowerfirst {
		sb := strings.Builder{}
		sb.WriteString(strings.ToLower(p[:1]))
		sb.WriteString(p[1:])
		return sb.String()
	}
	return p
}
