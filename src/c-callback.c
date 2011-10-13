#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"



// extern uintptr callWriteFunctionCallback(void* p0, schar* p1, uint64 p2, GoInterface p3);

typedef uintptr (*callWriteFunctionCallback_t)(void* p0, schar* p1, uint64 p2, GoInterface p3);


typedef size_t (*WRITEFUNCTION)( char *ptr, size_t size, size_t nmemb, void *userdata);


size_t writefunction_template( char *ptr, size_t size, size_t nmemb, GoInterface userdata) {
    register void *func = (void *)0x5151515151515151;
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
    return ((callWriteFunctionCallback_t)q)(func, ptr, size*nmemb, userdata);
//    return (size_t)func;

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

    printf ("in c, functionSpace %ul\n", functionSpace);

    p2 = (char *)functionSpace;

    for (i = 0; i< 50; i++) {
        if (p[i] == '\x51') {
            ((void **)(p2+i))[0] = go_func_pointer;
/*
            printf ("debug~ size of go_func_pointer%d\n",sizeof(go_func_pointer));
            for (j = 0; j < sizeof(go_func_pointer); j++) {
                p2[i] = (char *)(&go_func_pointer)[i];
                    //&((void **)(p + i)) = go_func_pointer;
                    } */
            printf ("modify mem ok\n");
            break;


        }
    }

    for (i = 0; i < 50; ++i)
    {
        printf("%d: %08x -- %08x\n", i, p[i], ((char *)functionSpace)[i]);
        if (p[i] == '\xc3')
        {
            break;
        }

    }


    my_callback = (WRITEFUNCTION)functionSpace;
    printf("my_callback=%x\n", my_callback);
//    printf("~~~~~~~~~~~~~~~~wo~ca~lei~!~ %d\n", my_callback(NULL, 100,  100, NULL));

    return my_callback;
}
