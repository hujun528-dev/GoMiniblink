package GoMiniblink

import (
	"os"
	"reflect"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

func toJsValue(mb IMiniblink, es jsExecState, value interface{}) jsValue {
	if value == nil {
		return mbApi.jsUndefined()
	}
	switch value.(type) {
	case int:
		return mbApi.jsInt(int32(value.(int)))
	case int8:
		return mbApi.jsInt(int32(value.(int8)))
	case int16:
		return mbApi.jsInt(int32(value.(int16)))
	case int32:
		return mbApi.jsInt(value.(int32))
	case int64:
		return mbApi.jsDouble(float64(value.(int64)))
	case uint:
		return mbApi.jsInt(int32(value.(uint)))
	case uint8:
		return mbApi.jsInt(int32(value.(uint8)))
	case uint16:
		return mbApi.jsInt(int32(value.(uint16)))
	case uint32:
		return mbApi.jsInt(int32(value.(uint32)))
	case uint64:
		return mbApi.jsDouble(float64(value.(uint64)))
	case float32:
		return mbApi.jsDouble(float64(value.(float32)))
	case float64:
		return mbApi.jsDouble(value.(float64))
	case bool:
		return mbApi.jsBoolean(value.(bool))
	case string:
		return mbApi.jsString(es, value.(string))
	case time.Time:
		return mbApi.jsDouble(float64(value.(time.Time).Unix()))
	default:
		break
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		length := rv.Len()
		arr := mbApi.jsEmptyArray(es)
		mbApi.jsSetLength(es, arr, uint32(length))
		for i := 0; i < length; i++ {
			v := toJsValue(mb, es, rv.Index(i).Interface())
			mbApi.jsSetAt(es, arr, uint32(i), v)
		}
		return arr
	case reflect.Map:
		obj := mbApi.jsEmptyObject(es)
		kv := rv.MapRange()
		for kv.Next() && kv.Key().Kind() == reflect.String {
			k := kv.Key().Interface().(string)
			v := toJsValue(mb, es, kv.Value().Interface())
			mbApi.jsSet(es, obj, k, v)
		}
		return obj
	case reflect.Struct:
		obj := mbApi.jsEmptyObject(es)
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Field(i).Type().Name()
			v := toJsValue(mb, es, rv.Field(i).Interface())
			mbApi.jsSet(es, obj, f, v)
		}
		return obj
	case reflect.Func:
		jsFn := jsData{}
		name, _ := syscall.UTF16FromString("function")
		for i := 0; i < len(name); i++ {
			jsFn.name[i] = name[i]
		}
		var call = func(fnes jsExecState, obj, args jsValue, count uint32) jsValue {
			arr := make([]reflect.Value, count)
			for i := uint32(0); i < count; i++ {
				jv := mbApi.jsArg(fnes, i)
				arr[i] = reflect.ValueOf(toGoValue(mb, fnes, jv))
			}
			rs := rv.Call(arr)
			if len(rs) > 0 {
				return toJsValue(mb, fnes, rs[0].Interface())
			}
			return 0
		}
		jsFn.callAsFunction = syscall.NewCallbackCDecl(call)
		jsFn.finalize = syscall.NewCallbackCDecl(func(ptr uintptr) {
			delete(keepRef, ptr)
		})
		keepRef[jsFn.callAsFunction] = call
		return mbApi.jsFunction(es, &jsFn)
	}
	panic("不支持的go类型：" + rv.Kind().String() + "(" + rv.Type().String() + ")")
}

func toGoValue(mb IMiniblink, es jsExecState, value jsValue) interface{} {
	switch mbApi.jsTypeOf(value) {
	case jsType_NULL, jsType_UNDEFINED:
		return nil
	case jsType_NUMBER:
		return mbApi.jsToDouble(es, value)
	case jsType_BOOLEAN:
		return mbApi.jsToBoolean(es, value)
	case jsType_STRING:
		return mbApi.jsToTempString(es, value)
	case jsType_ARRAY:
		length := mbApi.jsGetLength(es, value)
		ps := make([]interface{}, length)
		for i := 0; i < length; i++ {
			v := mbApi.jsGetAt(es, value, uint32(i))
			ps[i] = toGoValue(mb, es, v)
		}
		return ps
	case jsType_OBJECT:
		ps := make(map[string]interface{})
		keys := mbApi.jsGetKeys(es, value)
		for _, k := range keys {
			v := mbApi.jsGet(es, value, k)
			ps[k] = toGoValue(mb, es, v)
		}
		return ps
	case jsType_FUNCTION:
		name := "func" + strconv.FormatInt(time.Now().UnixNano(), 32)
		mbApi.jsSetGlobal(es, name, value)
		return JsFunc(func(param ...interface{}) interface{} {
			jses := mbApi.wkeGlobalExec(mb.GetHandle())
			ps := make([]jsValue, len(param))
			for i, v := range param {
				ps[i] = toJsValue(mb, jses, v)
			}
			fn := mbApi.jsGetGlobal(jses, name)
			rs := mbApi.jsCall(jses, fn, mbApi.jsUndefined(), ps)
			mbApi.jsSetGlobal(jses, name, mbApi.jsUndefined())
			return toGoValue(mb, jses, rs)
		})
	default:
		panic("不支持的js类型：" + strconv.Itoa(int(value)))
	}
}

var seed uint32 = 0

func seq() uint32 {
	seed++
	return seed
}

func toBool(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}

func toCallStr(str string) []byte {
	buf := []byte(str)
	rs := make([]byte, len(str)+1)
	for i, v := range buf {
		rs[i] = v
	}
	return rs
}

func wkePtrToUtf8(ptr uintptr) string {
	var seq []byte
	for {
		b := *((*byte)(unsafe.Pointer(ptr)))
		if b != 0 {
			seq = append(seq, b)
			ptr++
		} else {
			break
		}
	}
	return string(seq)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
