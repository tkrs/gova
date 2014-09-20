#ifndef JVMGO_H
#define JVMGO_H

#include "jni.h"

int InitJVM();
jstring _NewStringUTF(char *);
jclass _FindClass(char *);
jmethodID _GetStaticMethodID(jclass, const char *, const char *);
void _CallStaticVoidMethod(jclass, jmethodID, jstring);
const char *_GetStringUTFChars(jstring, jboolean *);
int _ReleaseStringUTFChars(jstring, const char *);
int DestroyJavaVM();

#endif
