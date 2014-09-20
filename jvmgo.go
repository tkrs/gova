package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ljvm

#include<jni.h>
#include <stdio.h>
#include <stdlib.h>

JNIEnv *env;
JavaVM *jvm;

int InitJVM() {
    int ret;
    JavaVMInitArgs vm_args;
    vm_args.version = JNI_VERSION_1_8;
    JNI_GetDefaultJavaVMInitArgs(&vm_args);
    if (ret != JNI_OK) {
        printf("vm init args error:%d\n", ret);
        return -1;
    }

    JavaVMOption options[2];
    options[0].optionString = "-Djava.compiler=NONE";
    options[1].optionString = "-Djava.class.path=.";
    //options[2].optionString = "-verbose:jni";
    //options[3].optionString = "-Djava.library.path=C:\\Program Files\\Java\\jdk1.8.0_05\\lib";

    vm_args.nOptions = 1;
    vm_args.options = options;

    ret = JNI_CreateJavaVM(&jvm, (void **)&env, &vm_args);
    if (ret != JNI_OK) {
        printf("create vm error:%d\n", ret);
        return -1;
    }
    return 0;
}

jstring _NewStringUTF(char *bytes) {
    jstring ret = (*env)->NewStringUTF(env, bytes);
    if (ret == 0) {
        fprintf(stderr, "NewStringUTF error\n");
    }
    return ret;
}

jclass _FindClass(char *bytes) {
    jclass ret = (*env)->FindClass(env, bytes);
    if (ret == 0) {
        fprintf(stderr, "NewStringUTF error\n");
    }
    return ret;
}

jmethodID _GetStaticMethodID(jclass clazz, const char *name, const char *sig) {
    jmethodID ret = (*env)->GetStaticMethodID(env, clazz, name, sig);
    if (ret == 0) {
        fprintf(stderr, "GetStaticMethodID error\n");
        return 0;
    }
    return ret;
}

void _CallStaticVoidMethod(jclass clazz, jmethodID methodID, jstring arg) {
    (*env)->CallStaticVoidMethod(env, clazz, methodID, arg);
}

const char *_GetStringUTFChars(jstring string, jboolean *isCopy) {
    const char *nativeString = (*env)->GetStringUTFChars(env, string, isCopy);
    return nativeString;
}

int _ReleaseStringUTFChars(jstring string, const char *utf) {
    (*env)->ReleaseStringUTFChars(env, string, utf);
    return 0;
}

int DestroyJavaVM() {
    (*jvm)->DestroyJavaVM(jvm);
    return 0;
}
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
