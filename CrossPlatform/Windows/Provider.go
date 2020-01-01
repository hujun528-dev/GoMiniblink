package Windows

import (
	MB "GoMiniblink"
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"os"
	"syscall"
	"time"
	"unsafe"
)

type Provider struct {
	hInstance  win32.HINSTANCE
	className  string
	main       string
	handleWnds map[win32.HWND]baseWindow
	nameWnds   map[string]baseWindow
	defOwner   win32.HWND
	defIcon    win32.HICON
	msClick    *mouseClickWorker
}

func (_this *Provider) Init() *Provider {
	_this.handleWnds = make(map[win32.HWND]baseWindow)
	_this.nameWnds = make(map[string]baseWindow)
	_this.className = Utils.NewUUID()
	_this.hInstance = win32.GetModuleHandle(nil)
	_this.msClick = new(mouseClickWorker).init()
	return _this
}

func (_this *Provider) GetScreen() MB.Screen {
	var s = MB.Screen{
		Full: MB.Rect{
			Wdith:  int(win32.GetSystemMetrics(win32.SM_CXSCREEN)),
			Height: int(win32.GetSystemMetrics(win32.SM_CYSCREEN)),
		},
		WorkArea: MB.Rect{
			Wdith:  int(win32.GetSystemMetrics(win32.SM_CXFULLSCREEN)),
			Height: int(win32.GetSystemMetrics(win32.SM_CYFULLSCREEN)),
		},
	}
	return s
}

func (_this *Provider) SetIcon(file string) {
	h := win32.LoadImage(_this.hInstance, sto16(file), win32.IMAGE_ICON, 0, 0, win32.LR_LOADFROMFILE)
	_this.defIcon = win32.HICON(h)
}

func (_this *Provider) registerWndClass() {
	var class = win32.WNDCLASSEX{
		Style:         win32.CS_HREDRAW | win32.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallback(_this.defaultMsgProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
		HIcon:         _this.defIcon,
		HIconSm:       _this.defIcon,
	}
	class.CbSize = uint32(unsafe.Sizeof(class))
	win32.RegisterClassEx(&class)
	_this.defOwner = win32.CreateWindowEx(0,
		sto16(_this.className), sto16(""),
		win32.WS_BORDER, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
}

func (_this *Provider) add(wnd baseWindow) {
	_this.nameWnds[wnd.name()] = wnd
}

func (_this *Provider) remove(hWnd win32.HWND, isExit bool) {
	if w, ok := _this.handleWnds[hWnd]; ok {
		delete(_this.nameWnds, w.name())
		delete(_this.handleWnds, hWnd)
		if isExit && w.name() == _this.main {
			_this.Exit(0)
		}
	}
}

func (_this *Provider) defaultMsgProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	isdlg := false
	if msg == win32.WM_CREATE {
		cp := *((*win32.CREATESTRUCT)(unsafe.Pointer(lParam)))
		if cp.CreateParams != 0 {
			id := *((*string)(unsafe.Pointer(cp.CreateParams)))
			if w, ok := _this.nameWnds[id]; ok {
				isdlg = w.isDialog()
				_this.handleWnds[hWnd] = w
				w.fireWndCreate(hWnd)
			}
		}
	} else if msg == win32.WM_INITDIALOG && lParam != 0 {
		id := *((*string)(unsafe.Pointer(lParam)))
		if w, ok := _this.nameWnds[id]; ok {
			isdlg = true
			_this.handleWnds[hWnd] = w
			w.fireWndCreate(hWnd)
		}
	} else if w, ok := _this.handleWnds[hWnd]; ok {
		isdlg = w.isDialog()
		ret := w.fireWndProc(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
		ret = _this.sendMouseClick(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
	}
	if isdlg && msg != win32.WM_CLOSE {
		return 0
	} else {
		return win32.DefWindowProc(hWnd, msg, wParam, lParam)
	}
}

func (_this *Provider) sendMouseClick(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_LBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, MB.MouseButtons_Left, int(lParam))
	case win32.WM_LBUTTONUP:
		if _this.msClick.mouseUp(hWnd, MB.MouseButtons_Left, int(lParam)) {
			return 1
		}
	case win32.WM_RBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, MB.MouseButtons_Right, int(lParam))
	case win32.WM_RBUTTONUP:
		if _this.msClick.mouseUp(hWnd, MB.MouseButtons_Right, int(lParam)) {
			return 1
		}
	case win32.WM_MBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, MB.MouseButtons_Middle, int(lParam))
	case win32.WM_MBUTTONUP:
		if _this.msClick.mouseUp(hWnd, MB.MouseButtons_Middle, int(lParam)) {
			return 1
		}
	}
	return 0
}

func (_this *Provider) Exit(code int) {
	win32.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form CrossPlatform.IForm, show func()) {
	frm, ok := form.(*winForm)
	if ok == false {
		panic("类型不正确")
	}
	_this.registerWndClass()
	_this.main = frm.name()
	show()
	var message win32.MSG
	for {
		if win32.GetMessage(&message, 0, 0, 0) {
			if win32.IsDialogMessage(message.HWnd, &message) == false {
				win32.TranslateMessage(&message)
				win32.DispatchMessage(&message)
			}
		} else {
			break
		}
	}
	os.Exit(0)
}

type mouseClickWorker struct {
	down_hWnd  win32.HWND
	down_key   MB.MouseButtons
	down_pos   int
	click_hWnd win32.HWND
	click_key  MB.MouseButtons
	click_pos  int
	time       int64
	isDouble   bool
}

func (_this *mouseClickWorker) init() *mouseClickWorker {
	go _this.fire()
	return _this
}

func (_this *mouseClickWorker) fire() {
	for {
		time.Sleep(time.Millisecond)
		if _this.click_hWnd != 0 && _this.time <= time.Now().UnixNano() {
			x, y := int(win32.LOWORD(int32(_this.click_pos))), int(win32.HIWORD(int32(_this.click_pos)))
			e := MB.MouseEvArgs{
				Buttons:  _this.click_key,
				X:        x,
				Y:        y,
				Delta:    0,
				IsDouble: _this.isDouble,
				Time:     time.Now(),
			}
			win32.PostMessage(_this.click_hWnd, uint32(win32.WM_COMMAND), uintptr(cmd_mouse_click), uintptr(unsafe.Pointer(&e)))
			_this.click_hWnd = 0
			_this.click_key = 0
			_this.click_pos = 0
		}
	}
}

func (_this *mouseClickWorker) mouseDown(hWnd win32.HWND, key MB.MouseButtons, pos int) {
	_this.down_hWnd = hWnd
	_this.down_key = key
	_this.down_pos = pos
}

func (_this *mouseClickWorker) mouseUp(hWnd win32.HWND, key MB.MouseButtons, pos int) bool {
	if _this.down_hWnd == hWnd {
		if _this.down_key == key && _this.down_pos == pos {
			_this.click_key = key
			_this.click_pos = pos
			if _this.click_hWnd != hWnd {
				_this.click_hWnd = hWnd
				_this.isDouble = false
				_this.time = time.Now().Add(time.Millisecond * 200).UnixNano()
				return true
			} else if _this.time > time.Now().UnixNano() {
				_this.isDouble = true
				return true
			}
		}
	}
	_this.click_hWnd = 0
	return false
}
