#include <stdio.h>
#include <sys/mman.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"

/* for OPT_HEADERFUNCTION */
size_t header_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
    void *go_header_func = (void *)getCurlField((uintptr)ctx, "headerFunction");
    GoInterface *userdata = (GoInterface *)getCurlField((uintptr)ctx, "headerData");

    if (userdata == NULL) {
        return callWriteFunctionCallback(go_header_func, ptr, size*nmemb, nilInterface());
    }
    return callWriteFunctionCallback(go_header_func, ptr, size*nmemb, *userdata);
}

void *return_header_function() {
    return (void *)&header_function;
}


/* for OPT_WRITEFUNCTION */
size_t write_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
    void *go_write_func = (void *)getCurlField((uintptr)ctx, "writeFunction");
    GoInterface *userdata = (GoInterface *)getCurlField((uintptr)ctx, "writeData");

    if (userdata == NULL) {
        return callWriteFunctionCallback(go_write_func, ptr, size*nmemb, nilInterface());
    }
    return callWriteFunctionCallback(go_write_func, ptr, size*nmemb, *userdata);
}

void *return_write_function() {
    return (void *)&write_function;
}



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
