#include <stdio.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"

/* for OPT_HEADERFUNCTION */
size_t header_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
    void *go_header_func = (void *)goGetCurlField((uintptr)ctx, "headerFunction");
    GoInterface *userdata = (GoInterface *)goGetCurlField((uintptr)ctx, "headerData");

    if (userdata == NULL) {
	return goCallWriteFunctionCallback(go_header_func, ptr, size*nmemb, goNilInterface());
    }
    return goCallWriteFunctionCallback(go_header_func, ptr, size*nmemb, *userdata);
}

void *return_header_function() {
    return (void *)&header_function;
}


/* for OPT_WRITEFUNCTION */
size_t write_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
    void *go_write_func = (void *)goGetCurlField((uintptr)ctx, "writeFunction");
    GoInterface *userdata = (GoInterface *)goGetCurlField((uintptr)ctx, "writeData");

    if (userdata == NULL) {
	return goCallWriteFunctionCallback(go_write_func, ptr, size*nmemb, goNilInterface());
    }
    return goCallWriteFunctionCallback(go_write_func, ptr, size*nmemb, *userdata);
}

void *return_write_function() {
    return (void *)&write_function;
}

/* for OPT_READFUNCTION */
size_t read_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
    void *go_read_func = (void *)goGetCurlField((uintptr)ctx, "readFunction");
    GoInterface *userdata = (GoInterface *)goGetCurlField((uintptr)ctx, "readData");

    if (userdata == NULL) {
	return goCallReadFunctionCallback(go_read_func, ptr, size*nmemb, goNilInterface());
    }
    return goCallReadFunctionCallback(go_read_func, ptr, size*nmemb, *userdata);
}

void *return_read_function() {
    return (void *)&read_function;
}


/* for OPT_PROGRESSFUNCTION */
int progress_function(void *ctx, double dltotal, double dlnow, double ultotal, double ulnow) {
    void *go_progress_func = (void *)goGetCurlField((uintptr)ctx, "progressFunction");
    GoInterface *clientp = (GoInterface *)goGetCurlField((uintptr)ctx, "progressData");

    if (clientp == NULL) {
	return goCallProgressCallback(go_progress_func, goNilInterface(),
				    dltotal, dlnow, ultotal, ulnow);
    }
    return goCallProgressCallback(go_progress_func, *clientp,
				dltotal, dlnow, ultotal, ulnow);
}

void *return_progress_function() {
    return (void *)progress_function;
}
