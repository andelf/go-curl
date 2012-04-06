package curl

/*
#include <curl/curl.h>
#include "compat.h"

*/
import "C"

// for GlobalInit(flag)
const (
	GLOBAL_SSL     = C.CURL_GLOBAL_SSL
	GLOBAL_WIN32   = C.CURL_GLOBAL_WIN32
	GLOBAL_ALL     = C.CURL_GLOBAL_ALL
	GLOBAL_NOTHING = C.CURL_GLOBAL_NOTHING
	GLOBAL_DEFAULT = C.CURL_GLOBAL_DEFAULT
)

// CURLcode
const (
	//        E_OK                       = C.CURLE_OK
	E_UNSUPPORTED_PROTOCOL     = C.CURLE_UNSUPPORTED_PROTOCOL
	E_FAILED_INIT              = C.CURLE_FAILED_INIT
	E_URL_MALFORMAT            = C.CURLE_URL_MALFORMAT
	E_COULDNT_RESOLVE_PROXY    = C.CURLE_COULDNT_RESOLVE_PROXY
	E_COULDNT_RESOLVE_HOST     = C.CURLE_COULDNT_RESOLVE_HOST
	E_COULDNT_CONNECT          = C.CURLE_COULDNT_CONNECT
	E_FTP_WEIRD_SERVER_REPLY   = C.CURLE_FTP_WEIRD_SERVER_REPLY
	E_REMOTE_ACCESS_DENIED     = C.CURLE_REMOTE_ACCESS_DENIED
	E_FTP_WEIRD_PASS_REPLY     = C.CURLE_FTP_WEIRD_PASS_REPLY
	E_FTP_WEIRD_PASV_REPLY     = C.CURLE_FTP_WEIRD_PASV_REPLY
	E_FTP_WEIRD_227_FORMAT     = C.CURLE_FTP_WEIRD_227_FORMAT
	E_FTP_CANT_GET_HOST        = C.CURLE_FTP_CANT_GET_HOST
	E_FTP_COULDNT_SET_TYPE     = C.CURLE_FTP_COULDNT_SET_TYPE
	E_PARTIAL_FILE             = C.CURLE_PARTIAL_FILE
	E_FTP_COULDNT_RETR_FILE    = C.CURLE_FTP_COULDNT_RETR_FILE
	E_QUOTE_ERROR              = C.CURLE_QUOTE_ERROR
	E_HTTP_RETURNED_ERROR      = C.CURLE_HTTP_RETURNED_ERROR
	E_WRITE_ERROR              = C.CURLE_WRITE_ERROR
	E_UPLOAD_FAILED            = C.CURLE_UPLOAD_FAILED
	E_READ_ERROR               = C.CURLE_READ_ERROR
	E_OUT_OF_MEMORY            = C.CURLE_OUT_OF_MEMORY
	E_OPERATION_TIMEDOUT       = C.CURLE_OPERATION_TIMEDOUT
	E_FTP_PORT_FAILED          = C.CURLE_FTP_PORT_FAILED
	E_FTP_COULDNT_USE_REST     = C.CURLE_FTP_COULDNT_USE_REST
	E_RANGE_ERROR              = C.CURLE_RANGE_ERROR
	E_HTTP_POST_ERROR          = C.CURLE_HTTP_POST_ERROR
	E_SSL_CONNECT_ERROR        = C.CURLE_SSL_CONNECT_ERROR
	E_BAD_DOWNLOAD_RESUME      = C.CURLE_BAD_DOWNLOAD_RESUME
	E_FILE_COULDNT_READ_FILE   = C.CURLE_FILE_COULDNT_READ_FILE
	E_LDAP_CANNOT_BIND         = C.CURLE_LDAP_CANNOT_BIND
	E_LDAP_SEARCH_FAILED       = C.CURLE_LDAP_SEARCH_FAILED
	E_FUNCTION_NOT_FOUND       = C.CURLE_FUNCTION_NOT_FOUND
	E_ABORTED_BY_CALLBACK      = C.CURLE_ABORTED_BY_CALLBACK
	E_BAD_FUNCTION_ARGUMENT    = C.CURLE_BAD_FUNCTION_ARGUMENT
	E_INTERFACE_FAILED         = C.CURLE_INTERFACE_FAILED
	E_TOO_MANY_REDIRECTS       = C.CURLE_TOO_MANY_REDIRECTS
	E_UNKNOWN_TELNET_OPTION    = C.CURLE_UNKNOWN_TELNET_OPTION
	E_TELNET_OPTION_SYNTAX     = C.CURLE_TELNET_OPTION_SYNTAX
	E_PEER_FAILED_VERIFICATION = C.CURLE_PEER_FAILED_VERIFICATION
	E_GOT_NOTHING              = C.CURLE_GOT_NOTHING
	E_SSL_ENGINE_NOTFOUND      = C.CURLE_SSL_ENGINE_NOTFOUND
	E_SSL_ENGINE_SETFAILED     = C.CURLE_SSL_ENGINE_SETFAILED
	E_SEND_ERROR               = C.CURLE_SEND_ERROR
	E_RECV_ERROR               = C.CURLE_RECV_ERROR
	E_SSL_CERTPROBLEM          = C.CURLE_SSL_CERTPROBLEM
	E_SSL_CIPHER               = C.CURLE_SSL_CIPHER
	E_SSL_CACERT               = C.CURLE_SSL_CACERT
	E_BAD_CONTENT_ENCODING     = C.CURLE_BAD_CONTENT_ENCODING
	E_LDAP_INVALID_URL         = C.CURLE_LDAP_INVALID_URL
	E_FILESIZE_EXCEEDED        = C.CURLE_FILESIZE_EXCEEDED
	E_USE_SSL_FAILED           = C.CURLE_USE_SSL_FAILED
	E_SEND_FAIL_REWIND         = C.CURLE_SEND_FAIL_REWIND
	E_SSL_ENGINE_INITFAILED    = C.CURLE_SSL_ENGINE_INITFAILED
	E_LOGIN_DENIED             = C.CURLE_LOGIN_DENIED
	E_TFTP_NOTFOUND            = C.CURLE_TFTP_NOTFOUND
	E_TFTP_PERM                = C.CURLE_TFTP_PERM
	E_REMOTE_DISK_FULL         = C.CURLE_REMOTE_DISK_FULL
	E_TFTP_ILLEGAL             = C.CURLE_TFTP_ILLEGAL
	E_TFTP_UNKNOWNID           = C.CURLE_TFTP_UNKNOWNID
	E_REMOTE_FILE_EXISTS       = C.CURLE_REMOTE_FILE_EXISTS
	E_TFTP_NOSUCHUSER          = C.CURLE_TFTP_NOSUCHUSER
	E_CONV_FAILED              = C.CURLE_CONV_FAILED
	E_CONV_REQD                = C.CURLE_CONV_REQD
	E_SSL_CACERT_BADFILE       = C.CURLE_SSL_CACERT_BADFILE
	E_REMOTE_FILE_NOT_FOUND    = C.CURLE_REMOTE_FILE_NOT_FOUND
	E_SSH                      = C.CURLE_SSH
	E_SSL_SHUTDOWN_FAILED      = C.CURLE_SSL_SHUTDOWN_FAILED
	E_AGAIN                    = C.CURLE_AGAIN
	E_SSL_CRL_BADFILE          = C.CURLE_SSL_CRL_BADFILE
	E_SSL_ISSUER_ERROR         = C.CURLE_SSL_ISSUER_ERROR
	E_FTP_PRET_FAILED          = C.CURLE_FTP_PRET_FAILED
	E_RTSP_CSEQ_ERROR          = C.CURLE_RTSP_CSEQ_ERROR
	E_RTSP_SESSION_ERROR       = C.CURLE_RTSP_SESSION_ERROR
	E_FTP_BAD_FILE_LIST        = C.CURLE_FTP_BAD_FILE_LIST
	E_CHUNK_FAILED             = C.CURLE_CHUNK_FAILED
)

// for easy.Setopt(flag, ...)
const (
	// This is the FILE * or void * the regular output should be written to.
	// NOTE: don't use it, use OPT_WRITEDATA instead
	OPT_FILE = C.CURLOPT_FILE
	// The full URL to get/put. string
	OPT_URL = C.CURLOPT_URL
	// Port number to connect to, if other than default. int
	OPT_PORT = C.CURLOPT_PORT
	// Name of proxy to use. string
	OPT_PROXY = C.CURLOPT_PROXY
	// "name:password" to use when fetching. string
	OPT_USERPWD = C.CURLOPT_USERPWD
	// "name:password" to use with proxy. string
	OPT_PROXYUSERPWD = C.CURLOPT_PROXYUSERPWD
	// Range to get, specified as an ASCII string. string
	OPT_RANGE = C.CURLOPT_RANGE
	// Specified file stream to upload from (use as input):
	// NOTE: don't use it, use OPT_READDATA instead
	OPT_INFILE = C.CURLOPT_INFILE
	// Buffer to receive error messages in, must be at least CURL_ERROR_SIZE
	// bytes big. If this is not used, error messages go to stderr instead:
	// WARN: not implemented yet
	OPT_ERRORBUFFER = C.CURLOPT_ERRORBUFFER
	// Function that will be called to store the output.
	// function prototype: func([]byte, interface{}) bool
	OPT_WRITEFUNCTION = C.CURLOPT_WRITEFUNCTION
	// Function that will be called to read the input.
	// function prototype: func([]byte, interface{}) int
	// TIP: return may be size of data or flags
	OPT_READFUNCTION = C.CURLOPT_READFUNCTION
	// Time-out the read operation after this amount of seconds. int
	OPT_TIMEOUT = C.CURLOPT_TIMEOUT
	// WARN: not tested yet. int
	OPT_INFILESIZE = C.CURLOPT_INFILESIZE
	// POST static input fields. string
	OPT_POSTFIELDS = C.CURLOPT_POSTFIELDS
	// Set the referrer page (needed by some CGIs). string
	OPT_REFERER = C.CURLOPT_REFERER
	// Set the FTP PORT string (interface name, named or numerical IP address)
	// Use i.e '-' to use default address.
	// WARN: not tested yet.
	OPT_FTPPORT = C.CURLOPT_FTPPORT
	// Set the User-Agent string (examined by some CGIs). string
	OPT_USERAGENT = C.CURLOPT_USERAGENT
	// Set the "low speed limit". int
	OPT_LOW_SPEED_LIMIT = C.CURLOPT_LOW_SPEED_LIMIT
	// Set the "low speed time". int
	OPT_LOW_SPEED_TIME = C.CURLOPT_LOW_SPEED_TIME
	// Set the continuation offset. int
	OPT_RESUME_FROM = C.CURLOPT_RESUME_FROM
	// Set cookie in request. string
	OPT_COOKIE = C.CURLOPT_COOKIE
	// This points to a linked list of headers. []string
	OPT_HTTPHEADER = C.CURLOPT_HTTPHEADER
	// This points to a linked list of post entries. curl.Form
	OPT_HTTPPOST = C.CURLOPT_HTTPPOST
	// name of the file keeping your private SSL-certificate. string
	OPT_SSLCERT = C.CURLOPT_SSLCERT
	// password for the SSL or SSH private key. string
	OPT_KEYPASSWD = C.CURLOPT_KEYPASSWD
	// send TYPE parameter? bool
	OPT_CRLF = C.CURLOPT_CRLF
	// send linked-list of QUOTE commands. []string
	OPT_QUOTE = C.CURLOPT_QUOTE
	// callback userdata to HEADERFUNCTION. interface{}
	OPT_WRITEHEADER = C.CURLOPT_WRITEHEADER
	// point to a file to read the initial cookies from, also enables "cookie awareness". string
	OPT_COOKIEFILE = C.CURLOPT_COOKIEFILE
	// What version to specifically try to use. see below flag
	OPT_SSLVERSION = C.CURLOPT_SSLVERSION
	// What kind of HTTP time condition to use, see below flag
	OPT_TIMECONDITION = C.CURLOPT_TIMECONDITION
	// Time to use with the above condition. Specified in number of seconds
	// since 1 Jan 1970. int
	OPT_TIMEVALUE = C.CURLOPT_TIMEVALUE
	/* Custom request, for customizing the get command like
	   HTTP: DELETE, TRACE and others
	   FTP: to use a different list command
	*/
	// WARN: not tested
	OPT_CUSTOMREQUEST = C.CURLOPT_CUSTOMREQUEST
	// WARN: not implemented yet
	OPT_STDERR = C.CURLOPT_STDERR
	// send linked-list of post-transfer QUOTE commands. []string
	// WARN: not tested
	OPT_POSTQUOTE = C.CURLOPT_POSTQUOTE
	// WARN: DEPRECATED, do not use!
	OPT_WRITEINFO = C.CURLOPT_WRITEINFO
	// talk a lot. bool
	OPT_VERBOSE = C.CURLOPT_VERBOSE
	// throw the header out too. bool
	OPT_HEADER = C.CURLOPT_HEADER
	// shut off the progress meter. bool
	OPT_NOPROGRESS = C.CURLOPT_NOPROGRESS
	// use HEAD to get http document. bool
	OPT_NOBODY = C.CURLOPT_NOBODY
	// no output on http error codes >= 300. bool
	OPT_FAILONERROR = C.CURLOPT_FAILONERROR
	// this is an upload. bool
	OPT_UPLOAD = C.CURLOPT_UPLOAD
	// HTTP POST method. bool
	OPT_POST = C.CURLOPT_POST
	// bare names when listing directories(FTP). bool
	OPT_DIRLISTONLY = C.CURLOPT_DIRLISTONLY
	// Append instead of overwrite on upload(FTP)!. bool
	OPT_APPEND = C.CURLOPT_APPEND
	// Specify whether to read the user+password from the .netrc or the URL.
	// see below flags
	OPT_NETRC = C.CURLOPT_NETRC
	// use Location: Luke!. bool
	OPT_FOLLOWLOCATION = C.CURLOPT_FOLLOWLOCATION
	// transfer data in text/ASCII format. bool
	OPT_TRANSFERTEXT = C.CURLOPT_TRANSFERTEXT
	// HTTP PUT. bool
	OPT_PUT = C.CURLOPT_PUT
	// Function that will be called instead of the internal progress display.
	// function prototype: func(float64, float64, float64, float64, interface{}) bool
	OPT_PROGRESSFUNCTION = C.CURLOPT_PROGRESSFUNCTION
	// Data passed to the progress callback. interface{}
	OPT_PROGRESSDATA = C.CURLOPT_PROGRESSDATA
	// We want the referrer field set automatically when following locations. bool
	OPT_AUTOREFERER = C.CURLOPT_AUTOREFERER
	// Port of the proxy. int
	// can be set in the proxy string as well with: "[host]:[port]"
	OPT_PROXYPORT = C.CURLOPT_PROXYPORT
	// size of the POST input data, if strlen() is not good to use. int
	OPT_POSTFIELDSIZE = C.CURLOPT_POSTFIELDSIZE
	// tunnel non-http operations through a HTTP proxy. bool
	OPT_HTTPPROXYTUNNEL = C.CURLOPT_HTTPPROXYTUNNEL
	// Set the interface string to use as outgoing network interface. string
	OPT_INTERFACE = C.CURLOPT_INTERFACE
	/* Set the krb4/5 security level, this also enables krb4/5 awareness.  This
	 * is a string, 'clear', 'safe', 'confidential' or 'private'.  If the string
	 * is set but doesn't match one of these, 'private' will be used.  */
	OPT_KRBLEVEL = C.CURLOPT_KRBLEVEL
	// Set if we should verify the peer in ssl handshake, set true to verify.
	OPT_SSL_VERIFYPEER = C.CURLOPT_SSL_VERIFYPEER
	// The CApath or CAfile used to validate the peer certificate
	// this option is used only if SSL_VERIFYPEER is true
	// WARN: not tested
	OPT_CAINFO = C.CURLOPT_CAINFO
	// Maximum number of http redirects to follow. int
	OPT_MAXREDIRS = C.CURLOPT_MAXREDIRS
	// Pass a long set to 1 to get the date of the requested document (if
	// possible)! Pass a zero to shut it off. bool
	OPT_FILETIME = C.CURLOPT_FILETIME
	// This points to a linked list of telnet options. []string
	OPT_TELNETOPTIONS = C.CURLOPT_TELNETOPTIONS
	// Max amount of cached alive connections. int
	OPT_MAXCONNECTS = C.CURLOPT_MAXCONNECTS
	// WARN: DEPRECATED, do not use!
	OPT_CLOSEPOLICY = C.CURLOPT_CLOSEPOLICY
	/* Set to explicitly use a new connection for the upcoming transfer.
	   Do not use this unless you're absolutely sure of this, as it makes the
	   operation slower and is less friendly for the network. */
	// bool
	OPT_FRESH_CONNECT = C.CURLOPT_FRESH_CONNECT
	/* Set to explicitly forbid the upcoming transfer's connection to be re-used
	   when done. Do not use this unless you're absolutely sure of this, as it
	   makes the operation slower and is less friendly for the network. */
	// bool
	OPT_FORBID_REUSE = C.CURLOPT_FORBID_REUSE
	/* Set to a file name that contains random data for libcurl to use to
	   seed the random engine when doing SSL connects. */
	// string
	OPT_RANDOM_FILE = C.CURLOPT_RANDOM_FILE
	// Set to the Entropy Gathering Daemon socket pathname. string
	OPT_EGDSOCKET = C.CURLOPT_EGDSOCKET
	/* Time-out connect operations after this amount of seconds, if connects
	   are OK within this time, then fine... This only aborts the connect
	   phase. [Only works on unix-style/SIGALRM operating systems] */
	// int
	OPT_CONNECTTIMEOUT = C.CURLOPT_CONNECTTIMEOUT
	// Function that will be called to store headers.
	// function prototype: func([]byte, interface{}) bool
	OPT_HEADERFUNCTION = C.CURLOPT_HEADERFUNCTION
	/* Set this to force the HTTP request to get back to GET. Only really usable
	   if POST, PUT or a custom request have been used first.*/
	// bool
	OPT_HTTPGET = C.CURLOPT_HTTPGET
	/* Set if we should verify the Common name from the peer certificate in ssl
	 * handshake, set 1 to check existence, 2 to ensure that it matches the
	 * provided hostname. */
	OPT_SSL_VERIFYHOST = C.CURLOPT_SSL_VERIFYHOST
	/* Specify which file name to write all known cookies in after completed
	   operation. Set file name to "-" (dash) to make it go to stdout. */
	// string
	OPT_COOKIEJAR = C.CURLOPT_COOKIEJAR
	// Specify which SSL ciphers to use
	// WARN: not tested
	OPT_SSL_CIPHER_LIST = C.CURLOPT_SSL_CIPHER_LIST
	// Specify which HTTP version to use! see below
	OPT_HTTP_VERSION = C.CURLOPT_HTTP_VERSION
	/* Specifically switch on or off the FTP engine's use of the EPSV command. By
	   default, that one will always be attempted before the more traditional
	   PASV command. */
	// bool
	OPT_FTP_USE_EPSV = C.CURLOPT_FTP_USE_EPSV
	// type of the file keeping your SSL-certificate ("DER", "PEM", "ENG"). string
	OPT_SSLCERTTYPE = C.CURLOPT_SSLCERTTYPE
	// name of the file keeping your private SSL-key. string
	OPT_SSLKEY = C.CURLOPT_SSLKEY
	//type of the file keeping your private SSL-key ("DER", "PEM", "ENG"). string
	OPT_SSLKEYTYPE = C.CURLOPT_SSLKEYTYPE
	// crypto engine for the SSL-sub system
	OPT_SSLENGINE = C.CURLOPT_SSLENGINE
	/* set the crypto engine for the SSL-sub system as default
	   the param has no meaning... */
	// WARN: not tested
	OPT_SSLENGINE_DEFAULT = C.CURLOPT_SSLENGINE_DEFAULT
	// WARN: DEPRECATED, do not use!
	OPT_DNS_USE_GLOBAL_CACHE = C.CURLOPT_DNS_USE_GLOBAL_CACHE
	// DNS cache timeout. int
	OPT_DNS_CACHE_TIMEOUT = C.CURLOPT_DNS_CACHE_TIMEOUT
	// send linked-list of pre-transfer QUOTE commands. []string
	OPT_PREQUOTE = C.CURLOPT_PREQUOTE
	// WARN: not implemented yet
	OPT_DEBUGFUNCTION = C.CURLOPT_DEBUGFUNCTION
	// WARN: not implemented yet
	OPT_DEBUGDATA = C.CURLOPT_DEBUGDATA
	// mark this as start of a cookie session. bool
	OPT_COOKIESESSION = C.CURLOPT_COOKIESESSION
	/* The CApath directory used to validate the peer certificate
	   this option is used only if SSL_VERIFYPEER is true */
	// string
	OPT_CAPATH = C.CURLOPT_CAPATH
	// Instruct libcurl to use a smaller receive buffer. int
	OPT_BUFFERSIZE = C.CURLOPT_BUFFERSIZE
	/* Instruct libcurl to not use any signal/alarm handlers, even when using
	   timeouts. This option is useful for multi-threaded applications.
	   See libcurl-the-guide for more background information. */
	// bool
	OPT_NOSIGNAL = C.CURLOPT_NOSIGNAL
	// Provide a CURLShare for mutexing non-ts data
	// WARN: not tested
	OPT_SHARE = C.CURLOPT_SHARE
	/* indicates type of proxy. accepted values are CURLPROXY_HTTP (default),
	   CURLPROXY_SOCKS4, CURLPROXY_SOCKS4A and CURLPROXY_SOCKS5. */
	// see below
	OPT_PROXYTYPE = C.CURLOPT_PROXYTYPE
	/* Set the Accept-Encoding string. Use this to tell a server you would like
	   the response to be compressed. Before 7.21.6, this was known as
	   CURLOPT_ENCODING */
	// string
	OPT_ACCEPT_ENCODING = C.CURLOPT_ACCEPT_ENCODING
	// Set pointer to private data
	// WARN: not tested
	OPT_PRIVATE = C.CURLOPT_PRIVATE
	// Set aliases for HTTP 200 in the HTTP Response header. []string
	OPT_HTTP200ALIASES = C.CURLOPT_HTTP200ALIASES
	/* Continue to send authentication (user+password) when following locations,
	   even when hostname changed. This can potentially send off the name
	   and password to whatever host the server decides. */
	// WARN: not tested
	OPT_UNRESTRICTED_AUTH = C.CURLOPT_UNRESTRICTED_AUTH
	/* Specifically switch on or off the FTP engine's use of the EPRT command (
	   it also disables the LPRT attempt). By default, those ones will always be
	   attempted before the good old traditional PORT command. */
	// bool
	OPT_FTP_USE_EPRT = C.CURLOPT_FTP_USE_EPRT
	/* Set this to a bitmask value to enable the particular authentications
	   methods you like. Use this in combination with CURLOPT_USERPWD.
	   Note that setting multiple bits may cause extra network round-trips. */
	// WARN: not tested
	OPT_HTTPAUTH = C.CURLOPT_HTTPAUTH
	// WARN: not implemented yet
	OPT_SSL_CTX_FUNCTION = C.CURLOPT_SSL_CTX_FUNCTION
	// WARN: not implemented yet
	OPT_SSL_CTX_DATA = C.CURLOPT_SSL_CTX_DATA
	// FTP Option that causes missing dirs to be created on the remote server. see below
	OPT_FTP_CREATE_MISSING_DIRS = C.CURLOPT_FTP_CREATE_MISSING_DIRS
	/* Set this to a bitmask value to enable the particular authentications
	   methods you like. Use this in combination with CURLOPT_PROXYUSERPWD.
	   Note that setting multiple bits may cause extra network round-trips. */
	// WARN: not tested
	OPT_PROXYAUTH = C.CURLOPT_PROXYAUTH
	/* FTP option that changes the timeout, in seconds, associated with
	   getting a response.  This is different from transfer timeout time and
	   essentially places a demand on the FTP server to acknowledge commands
	   in a timely manner. */
	// int
	OPT_FTP_RESPONSE_TIMEOUT    = C.CURLOPT_FTP_RESPONSE_TIMEOUT
	OPT_SERVER_RESPONSE_TIMEOUT = C.CURLOPT_SERVER_RESPONSE_TIMEOUT // alais of above
	/* Set this option to one of the CURL_IPRESOLVE_* defines (see below) to
	   tell libcurl to resolve names to those IP versions only. This only has
	   affect on systems with support for more than one, i.e IPv4 _and_ IPv6. */
	// see below
	OPT_IPRESOLVE = C.CURLOPT_IPRESOLVE
	/* Set this option to limit the size of a file that will be downloaded from
	   an HTTP or FTP server. int */
	OPT_MAXFILESIZE       = C.CURLOPT_MAXFILESIZE
	OPT_INFILESIZE_LARGE  = C.CURLOPT_INFILESIZE_LARGE
	OPT_RESUME_FROM_LARGE = C.CURLOPT_RESUME_FROM_LARGE
	OPT_MAXFILESIZE_LARGE = C.CURLOPT_MAXFILESIZE_LARGE
	/* Set this option to the file name of your .netrc file you want libcurl
	   to parse (using the CURLOPT_NETRC option). If not set, libcurl will do
	   a poor attempt to find the user's home directory and check for a .netrc
	   file in there. */
	// string
	OPT_NETRC_FILE = C.CURLOPT_NETRC_FILE
	/* Enable SSL/TLS for FTP, pick one of:
	   CURLFTPSSL_TRY     - try using SSL, proceed anyway otherwise
	   CURLFTPSSL_CONTROL - SSL for the control connection or fail
	   CURLFTPSSL_ALL     - SSL for all communication or fail
	*/
	OPT_USE_SSL             = C.CURLOPT_USE_SSL
	OPT_POSTFIELDSIZE_LARGE = C.CURLOPT_POSTFIELDSIZE_LARGE
	// Enable/disable the TCP Nagle algorithm. bool
	OPT_TCP_NODELAY   = C.CURLOPT_TCP_NODELAY
	OPT_FTPSSLAUTH    = C.CURLOPT_FTPSSLAUTH
	OPT_IOCTLFUNCTION = C.CURLOPT_IOCTLFUNCTION
	OPT_IOCTLDATA     = C.CURLOPT_IOCTLDATA
	OPT_FTP_ACCOUNT   = C.CURLOPT_FTP_ACCOUNT
	// feed cookies into cookie engine. string
	OPT_COOKIELIST = C.CURLOPT_COOKIELIST
	// ignore Content-Length. bool
	OPT_IGNORE_CONTENT_LENGTH      = C.CURLOPT_IGNORE_CONTENT_LENGTH
	OPT_FTP_SKIP_PASV_IP           = C.CURLOPT_FTP_SKIP_PASV_IP
	OPT_FTP_FILEMETHOD             = C.CURLOPT_FTP_FILEMETHOD
	OPT_LOCALPORT                  = C.CURLOPT_LOCALPORT
	OPT_LOCALPORTRANGE             = C.CURLOPT_LOCALPORTRANGE
	OPT_CONNECT_ONLY               = C.CURLOPT_CONNECT_ONLY
	OPT_CONV_FROM_NETWORK_FUNCTION = C.CURLOPT_CONV_FROM_NETWORK_FUNCTION
	OPT_CONV_TO_NETWORK_FUNCTION   = C.CURLOPT_CONV_TO_NETWORK_FUNCTION
	OPT_CONV_FROM_UTF8_FUNCTION    = C.CURLOPT_CONV_FROM_UTF8_FUNCTION
	OPT_MAX_SEND_SPEED_LARGE       = C.CURLOPT_MAX_SEND_SPEED_LARGE
	OPT_MAX_RECV_SPEED_LARGE       = C.CURLOPT_MAX_RECV_SPEED_LARGE
	OPT_FTP_ALTERNATIVE_TO_USER    = C.CURLOPT_FTP_ALTERNATIVE_TO_USER
	OPT_SOCKOPTFUNCTION            = C.CURLOPT_SOCKOPTFUNCTION
	OPT_SOCKOPTDATA                = C.CURLOPT_SOCKOPTDATA
	OPT_SSL_SESSIONID_CACHE        = C.CURLOPT_SSL_SESSIONID_CACHE
	OPT_SSH_AUTH_TYPES             = C.CURLOPT_SSH_AUTH_TYPES
	OPT_SSH_PUBLIC_KEYFILE         = C.CURLOPT_SSH_PUBLIC_KEYFILE
	OPT_SSH_PRIVATE_KEYFILE        = C.CURLOPT_SSH_PRIVATE_KEYFILE
	OPT_FTP_SSL_CCC                = C.CURLOPT_FTP_SSL_CCC
	OPT_TIMEOUT_MS                 = C.CURLOPT_TIMEOUT_MS
	OPT_CONNECTTIMEOUT_MS          = C.CURLOPT_CONNECTTIMEOUT_MS
	OPT_HTTP_TRANSFER_DECODING     = C.CURLOPT_HTTP_TRANSFER_DECODING
	OPT_HTTP_CONTENT_DECODING      = C.CURLOPT_HTTP_CONTENT_DECODING
	OPT_NEW_FILE_PERMS             = C.CURLOPT_NEW_FILE_PERMS
	OPT_NEW_DIRECTORY_PERMS        = C.CURLOPT_NEW_DIRECTORY_PERMS
	OPT_POSTREDIR                  = C.CURLOPT_POSTREDIR
	OPT_SSH_HOST_PUBLIC_KEY_MD5    = C.CURLOPT_SSH_HOST_PUBLIC_KEY_MD5
	OPT_OPENSOCKETFUNCTION         = C.CURLOPT_OPENSOCKETFUNCTION
	OPT_OPENSOCKETDATA             = C.CURLOPT_OPENSOCKETDATA
	OPT_COPYPOSTFIELDS             = C.CURLOPT_COPYPOSTFIELDS
	OPT_PROXY_TRANSFER_MODE        = C.CURLOPT_PROXY_TRANSFER_MODE
	OPT_SEEKFUNCTION               = C.CURLOPT_SEEKFUNCTION
	OPT_SEEKDATA                   = C.CURLOPT_SEEKDATA
	OPT_CRLFILE                    = C.CURLOPT_CRLFILE
	OPT_ISSUERCERT                 = C.CURLOPT_ISSUERCERT
	OPT_ADDRESS_SCOPE              = C.CURLOPT_ADDRESS_SCOPE
	OPT_CERTINFO                   = C.CURLOPT_CERTINFO
	OPT_USERNAME                   = C.CURLOPT_USERNAME
	OPT_PASSWORD                   = C.CURLOPT_PASSWORD
	OPT_PROXYUSERNAME              = C.CURLOPT_PROXYUSERNAME
	OPT_PROXYPASSWORD              = C.CURLOPT_PROXYPASSWORD
	OPT_NOPROXY                    = C.CURLOPT_NOPROXY
	OPT_TFTP_BLKSIZE               = C.CURLOPT_TFTP_BLKSIZE
	OPT_SOCKS5_GSSAPI_SERVICE      = C.CURLOPT_SOCKS5_GSSAPI_SERVICE
	OPT_SOCKS5_GSSAPI_NEC          = C.CURLOPT_SOCKS5_GSSAPI_NEC
	OPT_PROTOCOLS                  = C.CURLOPT_PROTOCOLS
	OPT_REDIR_PROTOCOLS            = C.CURLOPT_REDIR_PROTOCOLS
	OPT_SSH_KNOWNHOSTS             = C.CURLOPT_SSH_KNOWNHOSTS
	OPT_SSH_KEYFUNCTION            = C.CURLOPT_SSH_KEYFUNCTION
	OPT_SSH_KEYDATA                = C.CURLOPT_SSH_KEYDATA
	OPT_MAIL_FROM                  = C.CURLOPT_MAIL_FROM
	OPT_MAIL_RCPT                  = C.CURLOPT_MAIL_RCPT
	OPT_FTP_USE_PRET               = C.CURLOPT_FTP_USE_PRET
	OPT_RTSP_REQUEST               = C.CURLOPT_RTSP_REQUEST
	OPT_RTSP_SESSION_ID            = C.CURLOPT_RTSP_SESSION_ID
	OPT_RTSP_STREAM_URI            = C.CURLOPT_RTSP_STREAM_URI
	OPT_RTSP_TRANSPORT             = C.CURLOPT_RTSP_TRANSPORT
	OPT_RTSP_CLIENT_CSEQ           = C.CURLOPT_RTSP_CLIENT_CSEQ
	OPT_RTSP_SERVER_CSEQ           = C.CURLOPT_RTSP_SERVER_CSEQ
	OPT_INTERLEAVEDATA             = C.CURLOPT_INTERLEAVEDATA
	OPT_INTERLEAVEFUNCTION         = C.CURLOPT_INTERLEAVEFUNCTION
	OPT_WILDCARDMATCH              = C.CURLOPT_WILDCARDMATCH
	OPT_CHUNK_BGN_FUNCTION         = C.CURLOPT_CHUNK_BGN_FUNCTION
	OPT_CHUNK_END_FUNCTION         = C.CURLOPT_CHUNK_END_FUNCTION
	OPT_FNMATCH_FUNCTION           = C.CURLOPT_FNMATCH_FUNCTION
	OPT_CHUNK_DATA                 = C.CURLOPT_CHUNK_DATA
	OPT_FNMATCH_DATA               = C.CURLOPT_FNMATCH_DATA
	OPT_RESOLVE                    = C.CURLOPT_RESOLVE
	OPT_TLSAUTH_USERNAME           = C.CURLOPT_TLSAUTH_USERNAME
	OPT_TLSAUTH_PASSWORD           = C.CURLOPT_TLSAUTH_PASSWORD
	OPT_TLSAUTH_TYPE               = C.CURLOPT_TLSAUTH_TYPE
	OPT_TRANSFER_ENCODING          = C.CURLOPT_TRANSFER_ENCODING
	// unsupported
	//OPT_CLOSESOCKETFUNCTION        = C.CURLOPT_CLOSESOCKETFUNCTION
	//OPT_CLOSESOCKETDATA            = C.CURLOPT_CLOSESOCKETDATA
	// alias
	OPT_WRITEDATA  = C.CURLOPT_WRITEDATA
	OPT_READDATA   = C.CURLOPT_READDATA
	OPT_HEADERDATA = C.CURLOPT_HEADERDATA
	OPT_RTSPHEADER = C.CURLOPT_RTSPHEADER
)

// for easy.Getinfo(flag)
const (
	INFO_EFFECTIVE_URL           = C.CURLINFO_EFFECTIVE_URL
	INFO_RESPONSE_CODE           = C.CURLINFO_RESPONSE_CODE
	INFO_TOTAL_TIME              = C.CURLINFO_TOTAL_TIME
	INFO_NAMELOOKUP_TIME         = C.CURLINFO_NAMELOOKUP_TIME
	INFO_CONNECT_TIME            = C.CURLINFO_CONNECT_TIME
	INFO_PRETRANSFER_TIME        = C.CURLINFO_PRETRANSFER_TIME
	INFO_SIZE_UPLOAD             = C.CURLINFO_SIZE_UPLOAD
	INFO_SIZE_DOWNLOAD           = C.CURLINFO_SIZE_DOWNLOAD
	INFO_SPEED_DOWNLOAD          = C.CURLINFO_SPEED_DOWNLOAD
	INFO_SPEED_UPLOAD            = C.CURLINFO_SPEED_UPLOAD
	INFO_HEADER_SIZE             = C.CURLINFO_HEADER_SIZE
	INFO_REQUEST_SIZE            = C.CURLINFO_REQUEST_SIZE
	INFO_SSL_VERIFYRESULT        = C.CURLINFO_SSL_VERIFYRESULT
	INFO_FILETIME                = C.CURLINFO_FILETIME
	INFO_CONTENT_LENGTH_DOWNLOAD = C.CURLINFO_CONTENT_LENGTH_DOWNLOAD
	INFO_CONTENT_LENGTH_UPLOAD   = C.CURLINFO_CONTENT_LENGTH_UPLOAD
	INFO_STARTTRANSFER_TIME      = C.CURLINFO_STARTTRANSFER_TIME
	INFO_CONTENT_TYPE            = C.CURLINFO_CONTENT_TYPE
	INFO_REDIRECT_TIME           = C.CURLINFO_REDIRECT_TIME
	INFO_REDIRECT_COUNT          = C.CURLINFO_REDIRECT_COUNT
	INFO_PRIVATE                 = C.CURLINFO_PRIVATE
	INFO_HTTP_CONNECTCODE        = C.CURLINFO_HTTP_CONNECTCODE
	INFO_HTTPAUTH_AVAIL          = C.CURLINFO_HTTPAUTH_AVAIL
	INFO_PROXYAUTH_AVAIL         = C.CURLINFO_PROXYAUTH_AVAIL
	INFO_OS_ERRNO                = C.CURLINFO_OS_ERRNO
	INFO_NUM_CONNECTS            = C.CURLINFO_NUM_CONNECTS
	INFO_SSL_ENGINES             = C.CURLINFO_SSL_ENGINES
	INFO_COOKIELIST              = C.CURLINFO_COOKIELIST
	INFO_LASTSOCKET              = C.CURLINFO_LASTSOCKET
	INFO_FTP_ENTRY_PATH          = C.CURLINFO_FTP_ENTRY_PATH
	INFO_REDIRECT_URL            = C.CURLINFO_REDIRECT_URL
	INFO_PRIMARY_IP              = C.CURLINFO_PRIMARY_IP
	INFO_APPCONNECT_TIME         = C.CURLINFO_APPCONNECT_TIME
	INFO_CERTINFO                = C.CURLINFO_CERTINFO
	INFO_CONDITION_UNMET         = C.CURLINFO_CONDITION_UNMET
	INFO_RTSP_SESSION_ID         = C.CURLINFO_RTSP_SESSION_ID
	INFO_RTSP_CLIENT_CSEQ        = C.CURLINFO_RTSP_CLIENT_CSEQ
	INFO_RTSP_SERVER_CSEQ        = C.CURLINFO_RTSP_SERVER_CSEQ
	INFO_RTSP_CSEQ_RECV          = C.CURLINFO_RTSP_CSEQ_RECV
	INFO_PRIMARY_PORT            = C.CURLINFO_PRIMARY_PORT
	INFO_LOCAL_IP                = C.CURLINFO_LOCAL_IP
	INFO_LOCAL_PORT              = C.CURLINFO_LOCAL_PORT
)

// CURLMcode
const (
	M_CALL_MULTI_PERFORM = C.CURLM_CALL_MULTI_PERFORM
	//        M_OK                 = C.CURLM_OK
	M_BAD_HANDLE      = C.CURLM_BAD_HANDLE
	M_BAD_EASY_HANDLE = C.CURLM_BAD_EASY_HANDLE
	M_OUT_OF_MEMORY   = C.CURLM_OUT_OF_MEMORY
	M_INTERNAL_ERROR  = C.CURLM_INTERNAL_ERROR
	M_BAD_SOCKET      = C.CURLM_BAD_SOCKET
	M_UNKNOWN_OPTION  = C.CURLM_UNKNOWN_OPTION
)

// for multi.Setopt(flag, ...)
const (
	MOPT_SOCKETFUNCTION = C.CURLMOPT_SOCKETFUNCTION
	MOPT_SOCKETDATA     = C.CURLMOPT_SOCKETDATA
	MOPT_PIPELINING     = C.CURLMOPT_PIPELINING
	MOPT_TIMERFUNCTION  = C.CURLMOPT_TIMERFUNCTION
	MOPT_TIMERDATA      = C.CURLMOPT_TIMERDATA
	MOPT_MAXCONNECTS    = C.CURLMOPT_MAXCONNECTS
)

// CURLSHcode
const (
	//        SHE_OK         = C.CURLSHE_OK
	SHE_BAD_OPTION = C.CURLSHE_BAD_OPTION
	SHE_IN_USE     = C.CURLSHE_IN_USE
	SHE_INVALID    = C.CURLSHE_INVALID
	SHE_NOMEM      = C.CURLSHE_NOMEM
)

// for share.Setopt(flat, ...)
const (
	SHOPT_SHARE      = C.CURLSHOPT_SHARE
	SHOPT_UNSHARE    = C.CURLSHOPT_UNSHARE
	SHOPT_LOCKFUNC   = C.CURLSHOPT_LOCKFUNC
	SHOPT_UNLOCKFUNC = C.CURLSHOPT_UNLOCKFUNC
	SHOPT_USERDATA   = C.CURLSHOPT_USERDATA
)

// for share.Setopt(SHOPT_SHARE/SHOPT_UNSHARE, flag)
const (
	LOCK_DATA_SHARE       = C.CURL_LOCK_DATA_SHARE
	LOCK_DATA_COOKIE      = C.CURL_LOCK_DATA_COOKIE
	LOCK_DATA_DNS         = C.CURL_LOCK_DATA_DNS
	LOCK_DATA_SSL_SESSION = C.CURL_LOCK_DATA_SSL_SESSION
	LOCK_DATA_CONNECT     = C.CURL_LOCK_DATA_CONNECT
)

// for VersionInfo(flag)
const (
	VERSION_FIRST  = C.CURLVERSION_FIRST
	VERSION_SECOND = C.CURLVERSION_SECOND
	VERSION_THIRD  = C.CURLVERSION_THIRD
	VERSION_FOURTH = C.CURLVERSION_FOURTH
	VERSION_LAST   = C.CURLVERSION_LAST
	VERSION_NOW    = C.CURLVERSION_NOW
)

// for VersionInfo(...).Features mask flag
const (
	VERSION_IPV6         = C.CURL_VERSION_IPV6
	VERSION_KERBEROS4    = C.CURL_VERSION_KERBEROS4
	VERSION_SSL          = C.CURL_VERSION_SSL
	VERSION_LIBZ         = C.CURL_VERSION_LIBZ
	VERSION_NTLM         = C.CURL_VERSION_NTLM
	VERSION_GSSNEGOTIATE = C.CURL_VERSION_GSSNEGOTIATE
	VERSION_DEBUG        = C.CURL_VERSION_DEBUG
	VERSION_ASYNCHDNS    = C.CURL_VERSION_ASYNCHDNS
	VERSION_SPNEGO       = C.CURL_VERSION_SPNEGO
	VERSION_LARGEFILE    = C.CURL_VERSION_LARGEFILE
	VERSION_IDN          = C.CURL_VERSION_IDN
	VERSION_SSPI         = C.CURL_VERSION_SSPI
	VERSION_CONV         = C.CURL_VERSION_CONV
	VERSION_CURLDEBUG    = C.CURL_VERSION_CURLDEBUG
	VERSION_TLSAUTH_SRP  = C.CURL_VERSION_TLSAUTH_SRP
	VERSION_NTLM_WB      = C.CURL_VERSION_NTLM_WB
)

// for OPT_READFUNCTION, return a int flag
const (
	READFUNC_ABORT = C.CURL_READFUNC_ABORT
	READFUNC_PAUSE = C.CURL_READFUNC_PAUSE
)

// for easy.Setopt(OPT_HTTP_VERSION, flag)
const (
	HTTP_VERSION_NONE = C.CURL_HTTP_VERSION_NONE
	HTTP_VERSION_1_0  = C.CURL_HTTP_VERSION_1_0
	HTTP_VERSION_1_1  = C.CURL_HTTP_VERSION_1_1
)

// for easy.Setopt(OPT_PROXYTYPE, flag)
const (
	PROXY_HTTP     = C.CURLPROXY_HTTP     /* added in 7.10, new in 7.19.4 default is to use CONNECT HTTP/1.1 */
	PROXY_HTTP_1_0 = C.CURLPROXY_HTTP_1_0 /* added in 7.19.4, force to use CONNECT HTTP/1.0  */
	PROXY_SOCKS4   = C.CURLPROXY_SOCKS4   /* support added in 7.15.2, enum existed already in 7.10 */
	PROXY_SOCKS5   = C.CURLPROXY_SOCKS5   /* added in 7.10 */
	PROXY_SOCKS4A  = C.CURLPROXY_SOCKS4A  /* added in 7.18.0 */
	// Use the SOCKS5 protocol but pass along the host name rather than the IP address.
	PROXY_SOCKS5_HOSTNAME = C.CURLPROXY_SOCKS5_HOSTNAME
)

// for easy.Setopt(OPT_SSLVERSION, flag)
const (
	SSLVERSION_DEFAULT = C.CURL_SSLVERSION_DEFAULT
	SSLVERSION_TLSv1   = C.CURL_SSLVERSION_TLSv1
	SSLVERSION_SSLv2   = C.CURL_SSLVERSION_SSLv2
	SSLVERSION_SSLv3   = C.CURL_SSLVERSION_SSLv3
)

// for easy.Setopt(OPT_TIMECONDITION, flag)
const (
	TIMECOND_NONE         = C.CURL_TIMECOND_NONE
	TIMECOND_IFMODSINCE   = C.CURL_TIMECOND_IFMODSINCE
	TIMECOND_IFUNMODSINCE = C.CURL_TIMECOND_IFUNMODSINCE
	TIMECOND_LASTMOD      = C.CURL_TIMECOND_LASTMOD
)

// for easy.Setopt(OPT_NETRC, flag)
const (
	NETRC_IGNORED  = C.CURL_NETRC_IGNORED
	NETRC_OPTIONAL = C.CURL_NETRC_OPTIONAL
	NETRC_REQUIRED = C.CURL_NETRC_REQUIRED
)

// for easy.Setopt(OPT_FTP_CREATE_MISSING_DIRS, flag)
const (
	FTP_CREATE_DIR_NONE  = C.CURLFTP_CREATE_DIR_NONE
	FTP_CREATE_DIR       = C.CURLFTP_CREATE_DIR
	FTP_CREATE_DIR_RETRY = C.CURLFTP_CREATE_DIR_RETRY
)

// for easy.Setopt(OPT_IPRESOLVE, flag)
const (
	IPRESOLVE_WHATEVER = C.CURL_IPRESOLVE_WHATEVER
	IPRESOLVE_V4       = C.CURL_IPRESOLVE_V4
	IPRESOLVE_V6       = C.CURL_IPRESOLVE_V6
)

// for easy.Pause(flat)
const (
	PAUSE_RECV      = C.CURLPAUSE_RECV
	PAUSE_RECV_CONT = C.CURLPAUSE_RECV_CONT
	PAUSE_SEND      = C.CURLPAUSE_SEND
	PAUSE_SEND_CONT = C.CURLPAUSE_SEND_CONT
	PAUSE_ALL       = C.CURLPAUSE_ALL
	PAUSE_CONT      = C.CURLPAUSE_CONT
)
