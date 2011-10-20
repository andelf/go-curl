#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"



// extern uintptr callWriteFunctionCallback(void* p0, schar* p1, uint64 p2, GoInterface p3);

typedef uintptr (*callWriteFunctionCallback_t)(void* p0, schar* p1, uint64 p2, GoInterface p3);


typedef size_t (*WRITEFUNCTION)( char *ptr, size_t size, size_t nmemb, void *userdata);

/*
size_t writefunction_template( char *ptr, size_t size, size_t nmemb, GoInterface userdata) {
    volatile register void *func = (void *)0x5151515151515151;
    void *p = (void *)&foo;
    void *q = NULL;
    asm volatile (
        "movq %1, %%rax;"
        "nop; nop;"
        "mov %%rax, %0;"
        :"=r"(q)
        :"g"(p)
        :"%rax"
        );
//    foo(1, 2);
    ((FOO)q)(100, (void *)func);
    return size * nmemb;

}
*/


size_t writefunction_template( char *ptr, size_t size, size_t nmemb, GoInterface userdata) {
    volatile register void *func = (void *)0x5151515151515151;
    void *p = (void *)&callWriteFunctionCallback;
    void *q = NULL;
        /* NOTE: ugly but works */
        /* this is a niubi way to do it~! */
    __asm__ __volatile__ (
        "movq %1, %%rax;"
        "mov %%rax, %0;"
        :"=r"(q)
        :"g"(p)
        :"%rax"
        );
        /* debug */
    if (size == 1) {
        return 0;
    }

    return ((callWriteFunctionCallback_t)q)(func, ptr, size*nmemb, userdata);
//    return (size_t)func;

}


 // :)
size_t writefunction_static_func( char *ptr, size_t size, size_t nmemb, void *userdata) {
        /* use static variable to save values */
    static void *func = NULL;
    static int called_flag = 0;
    size_t ret = 0;

    if (ptr == NULL) {
            /* set callback */
        func = userdata;
        called_flag = 1;
    } else {
        if (called_flag == 0) {
                /* not setted */
            return 0;
        } else {
            ret = callWriteFunctionCallback(func, ptr, size*nmemb, *((GoInterface *)userdata));
            called_flag += 1;
        }
    }
    return ret;
}

void *return_sample_callback(void *go_func_pointer)
{
    writefunction_static_func(NULL, 0, 0, go_func_pointer);
    return &writefunction_static_func;
}

void *make_c_callback_function(void *go_func_pointer) {
    char *p = (char *)&writefunction_template;
    void *functionSpace = 0;
    WRITEFUNCTION my_callback = NULL;
    char *p2 = NULL;
    int i, j;
    GoInterface gi;

        /* create a exec-able mem-space */
    functionSpace = mmap(NULL, 4096, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_SHARED | MAP_ANONYMOUS, 0, 0);
    memcpy((char *)functionSpace, p, 1024);

    p2 = (char *)functionSpace;

    for (i = 0; i< 50; i++) {
        if (p[i] == '\x51') {
            ((void **)(p2+i))[0] = go_func_pointer;
            break;
        }
    }

    my_callback = (WRITEFUNCTION)functionSpace;

    return my_callback;
}
