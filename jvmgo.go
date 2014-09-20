package main
/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ljvm

#include<jni.h>
#include <stdio.h>
#include <stdlib.h>

JNIEnv *env;
JavaVM *jvm;

int CallDestroyJavaVM();
int InitJVM();

int InitJVM() {
    int ret;
    JavaVMInitArgs vm_args;
    vm_args.version = JNI_VERSION_1_8;
    JNI_GetDefaultJavaVMInitArgs(&vm_args);
    if (ret != JNI_OK) {
        printf("vm init args error:%d\n", ret);
        return ret;
    }

    JavaVMOption options[3];
    options[0].optionString = "-Djava.compiler=NONE";
    //options[1].optionString = "-Djava.library.path=C:\\Program Files\\Java\\jdk1.8.0_05\\lib";
    options[1].optionString = "-Djava.class.path=.";
    options[2].optionString = "-verbose:jni";

    vm_args.nOptions = 1;
    vm_args.options = options;

    ret = JNI_CreateJavaVM(&jvm, (void **)&env, &vm_args);
    if (ret != JNI_OK) {
        printf("create vm error:%d\n", ret);
        return ret;
    }

}

int Hello(char *name) {
    jstring jname = (*env)->NewStringUTF(env, name);

    jclass cls = (*env)->FindClass(env, "Hello");
    if (cls == 0) {
        fprintf(stderr, "FindClass error\n");
        return -1;
    }
    jmethodID mid = (*env)->GetStaticMethodID(env, cls, "say", "(Ljava/lang/String;)V");
     if (mid == 0) {
        fprintf(stderr, "GetStaticMethodID error\n");
        return -1;
    }
    (*env)->CallStaticVoidMethod(env, cls, mid, jname);
}

int CallDestroyJavaVM() {
    (*jvm)->DestroyJavaVM(jvm);
    return 0;
}
*/
import "C"

type JNI struct{}

func NewJavaVM() *JNI {
	ret, err := C.InitJVM()
	if err != nil {
		return nil
	}
	println(ret)
	return new(JNI)
}

func (*JNI) Hello(name *C.char) {
	C.Hello(name)
}

func (j *JNI) DestroyJavaVM() {
	ret, err := C.CallDestroyJavaVM()
	if err != nil {
		return
	}
	println(ret)
}

func main() {
	jvm := NewJavaVM()
	defer jvm.DestroyJavaVM()
	jvm.Hello(C.CString("桃白白"))
}
