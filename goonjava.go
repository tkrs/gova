package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -ljvm
// #include<jni.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	env := new(C.JNIEnv)
	jvm := new(C.JavaVM)
	vmArgs := new(C.struct_JavaVMInitArgs)
	vmArgs.version = C.JNI_VERSION_1_6

	ret, err := C.JNI_GetDefaultJavaVMInitArgs(unsafe.Pointer(vmArgs))
	if err != nil {
		return
	}
	println(ret)

	options := make([]C.struct_JavaVMOption, 3)
	options[0].optionString = C.CString("-Djava.compiler=NONE")
	options[1].optionString = C.CString("-Djava.library.path=C:\\Program Files\\Java\\jdk1.8.0_05\\lib")
	options[2].optionString = C.CString("verbose:jni")

	vmArgs.nOptions = 1
	vmArgs.options = &options[0]

	ret, err = C.JNI_CreateJavaVM(
		&jvm,
		(*unsafe.Pointer)(unsafe.Pointer(env)),
		unsafe.Pointer(vmArgs))

	if err != nil {
		return
	}
	println(ret)

	//	defer jvm.DestroyJavaVM()

	defer C.free(unsafe.Pointer(env))
	defer C.free(unsafe.Pointer(jvm))
}
