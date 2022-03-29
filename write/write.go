package ft

import (
    "unsafe"
    "syscall"
)

func Write(fd int, p []byte, length int) (n int, err error) {

    var _p0 unsafe.Pointer

    _p0 = unsafe.Pointer(&p[0])

    r0, _, e1 := syscall.Syscall(syscall.SYS_WRITE, uintptr(fd), uintptr(_p0), uintptr(length))
    n = int(r0)
    if e1 != 0 {
        err = e1
    }
    return
}
