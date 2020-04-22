package GoMiniblink

import (
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"syscall"
	"unsafe"
)

type freeApiForWindows struct {
	_dll *windows.LazyDLL

	_wkeInitialize          *windows.LazyProc
	_wkeCreateWebView       *windows.LazyProc
	_wkeSetHandle           *windows.LazyProc
	_wkeOnPaintBitUpdated   *windows.LazyProc
	_wkeLoadURL             *windows.LazyProc
	_wkeResize              *windows.LazyProc
	_wkeNetOnResponse       *windows.LazyProc
	_wkeOnLoadUrlBegin      *windows.LazyProc
	_wkePaint               *windows.LazyProc
	_wkeGetWidth            *windows.LazyProc
	_wkeGetHeight           *windows.LazyProc
	_wkeFireMouseEvent      *windows.LazyProc
	_wkeFireMouseWheelEvent *windows.LazyProc
	_wkeGetCursorInfoType   *windows.LazyProc
	_wkeFireKeyUpEvent      *windows.LazyProc
	_wkeFireKeyDownEvent    *windows.LazyProc
	_wkeFireKeyPressEvent   *windows.LazyProc
	_wkeGetCaretRect        *windows.LazyProc
	_wkeSetFocus            *windows.LazyProc
	_wkeNetGetRequestMethod *windows.LazyProc
	_wkeNetSetData          *windows.LazyProc
	_wkeNetCancelRequest    *windows.LazyProc
	_wkeJsBindFunction      *windows.LazyProc
	_jsArgCount             *windows.LazyProc
	_jsArg                  *windows.LazyProc
	_jsTypeOf               *windows.LazyProc
	_jsToTempString         *windows.LazyProc
	_jsToDoubleString       *windows.LazyProc
	_jsToInt                *windows.LazyProc
	_jsToBoolean            *windows.LazyProc
	_jsGetLength            *windows.LazyProc
	_jsGetAt                *windows.LazyProc
	_jsGetKeys              *windows.LazyProc
	_jsGet                  *windows.LazyProc
	_jsSetGlobal            *windows.LazyProc
	_jsGetGlobal            *windows.LazyProc
	_wkeGlobalExec          *windows.LazyProc
	_jsCall                 *windows.LazyProc
	_jsUndefined            *windows.LazyProc
	_jsInt                  *windows.LazyProc
	_jsBoolean              *windows.LazyProc
	_jsDoubleString         *windows.LazyProc
	_jsString               *windows.LazyProc
	_jsEmptyArray           *windows.LazyProc
	_jsSetLength            *windows.LazyProc
	_jsSetAt                *windows.LazyProc
	_jsFunction             *windows.LazyProc
	_jsEmptyObject          *windows.LazyProc
	_jsSet                  *windows.LazyProc
	_wkeDestroyWebView      *windows.LazyProc
	_jsGetWebView           *windows.LazyProc
	_wkeKillFocus           *windows.LazyProc
	_jsEvalExW              *windows.LazyProc
}

func (_this *freeApiForWindows) init() *freeApiForWindows {
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	var lib *windows.LazyDLL
	if is64 {
		lib = windows.NewLazyDLL("miniblink_x64.dll")
	} else {
		lib = windows.NewLazyDLL("miniblink_x86.dll")
	}
	_this._jsEvalExW = lib.NewProc("jsEvalExW")
	_this._wkeKillFocus = lib.NewProc("wkeKillFocus")
	_this._jsToInt = lib.NewProc("jsToInt")
	_this._jsSet = lib.NewProc("jsSet")
	_this._jsEmptyObject = lib.NewProc("jsEmptyObject")
	_this._jsFunction = lib.NewProc("jsFunction")
	_this._jsSetAt = lib.NewProc("jsSetAt")
	_this._jsSetLength = lib.NewProc("jsSetLength")
	_this._jsEmptyArray = lib.NewProc("jsEmptyArray")
	_this._jsString = lib.NewProc("jsString")
	_this._jsDoubleString = lib.NewProc("jsDoubleString")
	_this._jsBoolean = lib.NewProc("jsBoolean")
	_this._jsInt = lib.NewProc("jsInt")
	_this._jsUndefined = lib.NewProc("jsUndefined")
	_this._jsCall = lib.NewProc("jsCall")
	_this._wkeGlobalExec = lib.NewProc("wkeGlobalExec")
	_this._jsGetGlobal = lib.NewProc("jsGetGlobal")
	_this._jsSetGlobal = lib.NewProc("jsSetGlobal")
	_this._jsGet = lib.NewProc("jsGet")
	_this._jsGetKeys = lib.NewProc("jsGetKeys")
	_this._jsGetAt = lib.NewProc("jsGetAt")
	_this._jsGetLength = lib.NewProc("jsGetLength")
	_this._jsToBoolean = lib.NewProc("jsToBoolean")
	_this._jsToDoubleString = lib.NewProc("jsToDoubleString")
	_this._jsToTempString = lib.NewProc("jsToTempString")
	_this._jsTypeOf = lib.NewProc("jsTypeOf")
	_this._jsArg = lib.NewProc("jsArg")
	_this._jsArgCount = lib.NewProc("jsArgCount")
	_this._wkeJsBindFunction = lib.NewProc("wkeJsBindFunction")
	_this._wkeNetCancelRequest = lib.NewProc("wkeNetCancelRequest")
	_this._wkeNetSetData = lib.NewProc("wkeNetSetData")
	_this._wkeNetGetRequestMethod = lib.NewProc("wkeNetGetRequestMethod")
	_this._wkeFireKeyPressEvent = lib.NewProc("wkeFireKeyPressEvent")
	_this._wkeFireKeyUpEvent = lib.NewProc("wkeFireKeyUpEvent")
	_this._wkeFireKeyDownEvent = lib.NewProc("wkeFireKeyDownEvent")
	_this._wkeGetCursorInfoType = lib.NewProc("wkeGetCursorInfoType")
	_this._wkeFireMouseWheelEvent = lib.NewProc("wkeFireMouseWheelEvent")
	_this._wkeFireMouseEvent = lib.NewProc("wkeFireMouseEvent")
	_this._wkeGetHeight = lib.NewProc("wkeGetHeight")
	_this._wkeGetWidth = lib.NewProc("wkeGetWidth")
	_this._wkePaint = lib.NewProc("wkePaint")
	_this._wkeOnLoadUrlBegin = lib.NewProc("wkeOnLoadUrlBegin")
	_this._wkeNetOnResponse = lib.NewProc("wkeNetOnResponse")
	_this._wkeLoadURL = lib.NewProc("wkeLoadURL")
	_this._wkeResize = lib.NewProc("wkeResize")
	_this._wkeOnPaintBitUpdated = lib.NewProc("wkeOnPaintBitUpdated")
	_this._wkeSetHandle = lib.NewProc("wkeSetHandle")
	_this._wkeCreateWebView = lib.NewProc("wkeCreateWebView")
	_this._wkeInitialize = lib.NewProc("wkeInitialize")
	_this._wkeGetCaretRect = lib.NewProc("wkeGetCaretRect2")
	_this._wkeSetFocus = lib.NewProc("wkeSetFocus")
	_this._wkeDestroyWebView = lib.NewProc("wkeDestroyWebView")
	_this._jsGetWebView = lib.NewProc("jsGetWebView")

	ret, _, err := _this._wkeInitialize.Call()
	if ret == 0 {
		fmt.Println("初始化失败", err)
	}
	return _this
}

func (_this *freeApiForWindows) wkeKillFocus(wke wkeHandle) {
	_this._wkeKillFocus.Call(uintptr(wke))
}

func (_this *freeApiForWindows) jsGetWebView(es jsExecState) wkeHandle {
	r, _, _ := _this._jsGetWebView.Call(uintptr(es))
	return wkeHandle(r)
}

func (_this *freeApiForWindows) wkeDestroyWebView(wke wkeHandle) {
	_this._wkeDestroyWebView.Call(uintptr(wke))
}

func (_this *freeApiForWindows) jsSet(es jsExecState, obj jsValue, name string, value jsValue) {
	ptr := []byte(name)
	_this._jsSet.Call(uintptr(es), uintptr(obj), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
}

func (_this *freeApiForWindows) jsEmptyObject(es jsExecState) jsValue {
	r, _, _ := _this._jsEmptyObject.Call(uintptr(es))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsFunction(es jsExecState, data *jsData) jsValue {
	r, _, _ := _this._jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(data)))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsSetAt(es jsExecState, array jsValue, index uint32, value jsValue) {
	_this._jsSetAt.Call(uintptr(es), uintptr(array), uintptr(index), uintptr(value))
}

func (_this *freeApiForWindows) jsSetLength(es jsExecState, array jsValue, length uint32) {
	_this._jsSetLength.Call(uintptr(es), uintptr(array), uintptr(length))
}

func (_this *freeApiForWindows) jsEmptyArray(es jsExecState) jsValue {
	r, _, _ := _this._jsEmptyArray.Call(uintptr(es))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsString(es jsExecState, value string) jsValue {
	ptr := toCallStr(value)
	r, _, _ := _this._jsString.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsDouble(value float64) jsValue {
	ptr := toCallStr(strconv.FormatFloat(value, 'f', 9, 64))
	r, _, _ := _this._jsDoubleString.Call(uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsBoolean(value bool) jsValue {
	r, _, _ := _this._jsBoolean.Call(uintptr(toBool(value)))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsInt(value int32) jsValue {
	r, _, _ := _this._jsInt.Call(uintptr(value))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsCall(es jsExecState, fn, thisObject jsValue, args []jsValue) jsValue {
	var ptr = uintptr(0)
	if len(args) > 0 {
		ptr = uintptr(unsafe.Pointer(&args[0]))
	}
	r, _, _ := _this._jsCall.Call(uintptr(es), uintptr(fn), uintptr(thisObject), ptr, uintptr(len(args)))
	return jsValue(r)
}

func (_this *freeApiForWindows) wkeGlobalExec(wke wkeHandle) jsExecState {
	r, _, _ := _this._wkeGlobalExec.Call(uintptr(wke))
	return jsExecState(r)
}

func (_this *freeApiForWindows) jsGetGlobal(es jsExecState, name string) jsValue {
	ptr := toCallStr(name)
	r, _, _ := _this._jsGetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsSetGlobal(es jsExecState, name string, value jsValue) {
	ptr := toCallStr(name)
	_this._jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
}

func (_this *freeApiForWindows) jsEvalExW(es jsExecState, js string, isInClosure bool) jsValue {
	ptr := toCallStr(js)
	rs, _, _ := _this._jsEvalExW.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(toBool(isInClosure)))
	return jsValue(rs)
}

func (_this *freeApiForWindows) jsGetKeys(es jsExecState, value jsValue) []string {
	//rs, _, _ := _this._jsGetKeys.Call(uintptr(es), uintptr(value))
	//keys := *((*jsKeys)(unsafe.Pointer(rs)))
	//items := make([]string, keys.length)
	//for i := 0; i < len(items); i++ {
	//	items[i] = wkePtrToUtf8(*((*uintptr)(unsafe.Pointer(keys.first))))
	//	keys.first += unsafe.Sizeof(uintptr(0))
	//}
	//return items

	//_this._jsGetKeys.Call(uintptr(es), uintptr(value))
	//return []string{"n1", "n2"}

	json := _this.jsGetGlobal(es, "Object")
	stringify := _this.jsGet(es, json, "keys")
	rs := _this.jsCall(es, stringify, _this.jsUndefined(), []jsValue{value})
	alen := _this.jsGetLength(es, rs)
	items := make([]string, alen)
	for i := 0; i < len(items); i++ {
		v := _this.jsGetAt(es, rs, uint32(i))
		str := _this.jsToTempString(es, v)
		items[i] = str
	}
	return items
}

func (_this *freeApiForWindows) jsGet(es jsExecState, value jsValue, name string) jsValue {
	ptr := toCallStr(name)
	r, _, _ := _this._jsGet.Call(uintptr(es), uintptr(value), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsGetAt(es jsExecState, value jsValue, index uint32) jsValue {
	r, _, _ := _this._jsGetAt.Call(uintptr(es), uintptr(value), uintptr(index))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsGetLength(es jsExecState, value jsValue) int {
	r, _, _ := _this._jsGetLength.Call(uintptr(es), uintptr(value))
	return int(r)
}

func (_this *freeApiForWindows) jsUndefined() jsValue {
	r, _, _ := _this._jsUndefined.Call()
	return jsValue(r)
}

func (_this *freeApiForWindows) jsToBoolean(es jsExecState, value jsValue) bool {
	r, _, _ := _this._jsToBoolean.Call(uintptr(es), uintptr(value))
	return byte(r) != 0
}

func (_this *freeApiForWindows) jsToDouble(es jsExecState, value jsValue) float64 {
	r, _, _ := _this._jsToDoubleString.Call(uintptr(es), uintptr(value))
	str := wkePtrToUtf8(r)
	n, _ := strconv.ParseFloat(str, 10)
	return n
}

func (_this *freeApiForWindows) jsToTempString(es jsExecState, value jsValue) string {
	r, _, _ := _this._jsToTempString.Call(uintptr(es), uintptr(value))
	return wkePtrToUtf8(r)
}

func (_this *freeApiForWindows) jsTypeOf(value jsValue) jsType {
	r, _, _ := _this._jsTypeOf.Call(uintptr(value))
	return jsType(r)
}

func (_this *freeApiForWindows) jsArg(es jsExecState, index uint32) jsValue {
	r, _, _ := _this._jsArg.Call(uintptr(es), uintptr(index))
	return jsValue(r)
}

func (_this *freeApiForWindows) jsArgCount(es jsExecState) uint32 {
	r, _, _ := _this._jsArgCount.Call(uintptr(es))
	return uint32(r)
}

func (_this *freeApiForWindows) wkeJsBindFunction(name string, fn wkeJsNativeFunction, param uintptr, argCount uint32) {
	ptr := toCallStr(name)
	_this._wkeJsBindFunction.Call(uintptr(unsafe.Pointer(&ptr[0])), syscall.NewCallbackCDecl(fn), param, uintptr(argCount))
}

func (_this *freeApiForWindows) wkeNetCancelRequest(job wkeNetJob) {
	_this._wkeNetCancelRequest.Call(uintptr(job))
}

func (_this *freeApiForWindows) wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr) {
	_this._wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *freeApiForWindows) wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr) {
	_this._wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *freeApiForWindows) wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType {
	r, _, _ := _this._wkeNetGetRequestMethod.Call(uintptr(job))
	return wkeRequestType(r)
}

func (_this *freeApiForWindows) wkeNetSetData(job wkeNetJob, buf []byte) {
	if len(buf) == 0 {
		buf = []byte{0}
	}
	_this._wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func (_this *freeApiForWindows) wkeGetCaretRect(wke wkeHandle) wkeRect {
	r, _, _ := _this._wkeGetCaretRect.Call(uintptr(wke))
	return *((*wkeRect)(unsafe.Pointer(r)))
}

func (_this *freeApiForWindows) wkeSetFocus(wke wkeHandle) {
	_this._wkeSetFocus.Call(uintptr(wke))
}

func (_this *freeApiForWindows) wkeFireKeyPressEvent(wke wkeHandle, code int, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *freeApiForWindows) wkeFireKeyDownEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *freeApiForWindows) wkeFireKeyUpEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *freeApiForWindows) wkeGetCursorInfoType(wke wkeHandle) wkeCursorType {
	r, _, _ := _this._wkeGetCursorInfoType.Call(uintptr(wke))
	return wkeCursorType(r)
}

func (_this *freeApiForWindows) wkeFireMouseWheelEvent(wke wkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := _this._wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}

func (_this *freeApiForWindows) wkeFireMouseEvent(wke wkeHandle, message, x, y, flags int32) bool {
	r, _, _ := _this._wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func (_this *freeApiForWindows) wkePaint(wke wkeHandle, bits []byte, pitch uint32) {
	_this._wkePaint.Call(uintptr(wke), uintptr(unsafe.Pointer(&bits[0])), uintptr(pitch))
}

func (_this *freeApiForWindows) wkeGetHeight(wke wkeHandle) uint32 {
	r, _, _ := _this._wkeGetHeight.Call(uintptr(wke))
	return uint32(r)
}

func (_this *freeApiForWindows) wkeGetWidth(wke wkeHandle) uint32 {
	r, _, _ := _this._wkeGetWidth.Call(uintptr(wke))
	return uint32(r)
}

func (_this *freeApiForWindows) wkeResize(wke wkeHandle, w, h uint32) {
	_this._wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func (_this *freeApiForWindows) wkeLoadURL(wke wkeHandle, url string) {
	ptr := toCallStr(url)
	_this._wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&ptr[0])))
}

func (_this *freeApiForWindows) wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr) {
	_this._wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *freeApiForWindows) wkeSetHandle(wke wkeHandle, handle uintptr) {
	_this._wkeSetHandle.Call(uintptr(wke), handle)
}

func (_this *freeApiForWindows) wkeCreateWebView() wkeHandle {
	r, _, _ := _this._wkeCreateWebView.Call()
	return wkeHandle(r)
}
