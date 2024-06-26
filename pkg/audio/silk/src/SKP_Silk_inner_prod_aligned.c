 

/*                                                                                *
 * SKP_Silk_inner_prod_aligned.c                                                *
 *                                                                                *
 *                                                                          	   *
 * Copyright 2008-2010 (c), Skype Limited                                              *
 * Date: 080601                                                                   *
 *                                                                                */
#include "SKP_Silk_SigProc_FIX.h"

/* sum= for(i=0;i<len;i++)inVec1[i]*inVec2[i];      ---        inner product    */
/* Note for ARM asm:                                                            */
/*        * inVec1 and inVec2 should be at least 2 byte aligned.    (Or defined as short/int16) */
/*        * len should be positive 16bit integer.                               */
/*        * only when len>6, memory access can be reduced by half.              */

#if (EMBEDDED_ARM<5) 
SKP_int32 SKP_Silk_inner_prod_aligned(
    const SKP_int16* const inVec1,  /*    I input vector 1    */
    const SKP_int16* const inVec2,  /*    I input vector 2    */
    const SKP_int             len   /*    I vector lengths    */
)
{
    SKP_int   i; 
    SKP_int32 sum = 0;
    for( i = 0; i < len; i++ ) {
        sum = SKP_SMLABB( sum, inVec1[ i ], inVec2[ i ] );
    }
    return sum;
}
#endif

#if (EMBEDDED_ARM<5) 
SKP_int64 SKP_Silk_inner_prod16_aligned_64(
    const SKP_int16 *inVec1,        /*    I input vector 1    */ 
    const SKP_int16 *inVec2,        /*    I input vector 2    */
    const SKP_int   len             /*    I vector lengths    */
)
{
    SKP_int   i; 
    SKP_int64 sum = 0;
    for( i = 0; i < len; i++ ) {
        sum = SKP_SMLALBB( sum, inVec1[ i ], inVec2[ i ] );
    }
    return sum;
}
#endif
