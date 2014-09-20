#include "jvmgo.h"

JNIEnv *env;
JavaVM *jvm;

int InitJVM() {
    int ret;
    JavaVMInitArgs vm_args;
    vm_args.version = JNI_VERSION_1_8;
    ret = JNI_GetDefaultJavaVMInitArgs(&vm_args);
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
