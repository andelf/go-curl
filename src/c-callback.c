#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"



typedef size_t (*WRITEFUNCTION)( char *ptr, size_t size, size_t nmemb, void *userdata);


size_t writefunction_template( char *ptr, size_t size, size_t nmemb, GoInterface userdata) {
    register void *func = (void *)0x51515151;
    return callWriteFunctionCallback(func, ptr, size*nmemb, userdata);

}



void *make_c_callback_function(void *go_func_pointer) {
    char *p = (char *)&writefunction_template;
    void *functionSpace = 0;
    WRITEFUNCTION my_callback = NULL;
    char *p2 = NULL;
    int i;

    functionSpace = mmap(NULL, 4096, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_SHARED | MAP_ANONYMOUS, 0, 0);
    memcpy((char *)functionSpace, p, 1024);

    printf ("in c, functionSpace %ul\n", functionSpace);

    for (i = 0; i< 50; i++) {
        if (p[i] == '\x51') {
            printf("!");
        }
        printf("%d: %08x -- %08x\n", i, p[i], ((char *)functionSpace)[i]);
        if (p[i] == '\xc3')
        {
            break;
        }
    }

    p2 = functionSpace;
    for (i = 0; i < 50; i++) {
        if (p2[i] == '\xea') {
            p2[i] = '\x11';
        }
        if (p2[i] == '\xc3') {
            break;
        }
    }

    /* p2 = functionSpace; */
    /* p2[0] = 0xb8; */
    /* p2[1] = 0x00; */
    /* p2[2] = 0x11; */
    /* p2[3] = 0xdd; */
    /* p2[4] = 0xee; */
    /* p2[5] = 0xc3; */

    my_callback = (WRITEFUNCTION)functionSpace;
    printf("my_callback=%x\n", my_callback);

    return my_callback;
}
