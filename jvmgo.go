package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ljvmgo -ljvm
#include "jni.h"
#include "jvmgo.h"
*/
import "C"
import "log"

// JNIEnv is java native interface
type JNIEnv struct{}

// NewJNIEnv returns JNIEnv pointer
func NewJNIEnv() *JNIEnv {
	ret, err := C.InitJVM()

	if err != nil {
		panic(err)
	}

	if ret < 0 {
		return nil
	}

	return &JNIEnv{}
}

// NewStringUTF java native interface
func (*JNIEnv) NewStringUTF(s string) (jstring C.jstring, err error) {
	jstring, err = C._NewStringUTF(C.CString(s))
	return
}

// FindClass java native interface
func (*JNIEnv) FindClass(className string) (jclass C.jclass, err error) {
	jclass, err = C._FindClass(C.CString(className))
	return
}

// GetStaticMethodID java native interface
func (*JNIEnv) GetStaticMethodID(clazz C.jclass, name string, sig string) (jmethodID C.jmethodID, err error) {
	jmethodID, err = C._GetStaticMethodID(clazz, C.CString(name), C.CString(sig))
	return
}

// CallStaticVoidMethod java native interface
func (*JNIEnv) CallStaticVoidMethod(jclass C.jclass, jmethodID C.jmethodID, arg C.jstring) {
	C._CallStaticVoidMethod(jclass, jmethodID, arg)
}

// GetStringUTFChars
func (*JNIEnv) GetStringUTFChars(jstring C.jstring, isCopy *C.jboolean) (ret *C.char, err error) {
	ret, err = C._GetStringUTFChars(jstring, isCopy)
	return
}

func (*JNIEnv) ReleaseStringUTFChars(jstring C.jstring, utf *C.char) (err error) {
	_, err = C._ReleaseStringUTFChars(jstring, utf)
	return
}

// DestroyJavaVM is call destroy JavaVM
func DestroyJavaVM() {
	_, err := C.DestroyJavaVM()
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	env := NewJNIEnv()
	if env == nil {
		return
	}
	defer DestroyJavaVM()

	arg, err := env.NewStringUTF("桃白白")
	if err != nil {
		log.Fatal(err)
	}

	jclass, err := env.FindClass("Hello")
	if err != nil {
		log.Fatal(err)
	}

	jmethodID, err := env.GetStaticMethodID(jclass, "say", "(Ljava/lang/String;)V")
	if err != nil {
		log.Fatal(err)
	}

	env.CallStaticVoidMethod(jclass, jmethodID, arg)

}
