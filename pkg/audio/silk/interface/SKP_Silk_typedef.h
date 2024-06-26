#ifndef _SKP_SILK_API_TYPDEF_H_
#define _SKP_SILK_API_TYPDEF_H_

#ifndef SKP_USE_DOUBLE_PRECISION_FLOATS
#define SKP_USE_DOUBLE_PRECISION_FLOATS		0
#endif

#include <float.h>
#if defined( __GNUC__ )
#include <stdint.h>
#endif

#define SKP_int         int                     /* used for counters etc; at least 16 bits */
#ifdef __GNUC__
# define SKP_int64      int64_t
#else
# define SKP_int64      long long
#endif
#define SKP_int32       int
#define SKP_int16       short
#define SKP_int8        signed char

#define SKP_uint        unsigned int            /* used for counters etc; at least 16 bits */
#ifdef __GNUC__
# define SKP_uint64     uint64_t
#else
# define SKP_uint64     unsigned long long
#endif
#define SKP_uint32      unsigned int
#define SKP_uint16      unsigned short
#define SKP_uint8       unsigned char

#define SKP_int_ptr_size intptr_t

#if SKP_USE_DOUBLE_PRECISION_FLOATS
# define SKP_float      double
# define SKP_float_MAX  DBL_MAX
#else
# define SKP_float      float
# define SKP_float_MAX  FLT_MAX
#endif

#define SKP_INLINE      static __inline

#ifdef _WIN32
# define SKP_STR_CASEINSENSITIVE_COMPARE(x, y) _stricmp(x, y)
#else
# define SKP_STR_CASEINSENSITIVE_COMPARE(x, y) strcasecmp(x, y)
#endif 

#define	SKP_int64_MAX	((SKP_int64)0x7FFFFFFFFFFFFFFFLL)	/*  2^63 - 1  */
#define SKP_int64_MIN	((SKP_int64)0x8000000000000000LL)	/* -2^63	 */
#define	SKP_int32_MAX	0x7FFFFFFF							/*  2^31 - 1 =  2147483647*/
#define SKP_int32_MIN	((SKP_int32)0x80000000)				/* -2^31	 = -2147483648*/
#define	SKP_int16_MAX	0x7FFF								/*	2^15 - 1 =	32767*/
#define SKP_int16_MIN	((SKP_int16)0x8000)					/* -2^15	 = -32768*/
#define	SKP_int8_MAX	0x7F								/*	2^7 - 1  =  127*/
#define SKP_int8_MIN	((SKP_int8)0x80)					/* -2^7 	 = -128*/

#define SKP_uint32_MAX	0xFFFFFFFF	/* 2^32 - 1 = 4294967295 */
#define SKP_uint32_MIN	0x00000000
#define SKP_uint16_MAX	0xFFFF		/* 2^16 - 1 = 65535 */
#define SKP_uint16_MIN	0x0000
#define SKP_uint8_MAX	0xFF		/*  2^8 - 1 = 255 */
#define SKP_uint8_MIN	0x00

#define SKP_TRUE		1
#define SKP_FALSE		0

/* assertions */
#if (defined _WIN32 && !defined _WINCE && !defined(__GNUC__) && !defined(NO_ASSERTS))
# ifndef SKP_assert
#  include <crtdbg.h>      /* ASSERTE() */
#  define SKP_assert(COND)   _ASSERTE(COND)
# endif
#else
# define SKP_assert(COND)
#endif

#endif
