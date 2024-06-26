/*                                                                      *
 * SKP_Silk_int16_array_maxabs.c                                      *
 *                                                                      *
 * Function that returns the maximum absolut value of                   *
 * the input vector                                                     *
 *                                                                      *
 * Copyright 2006 (c), Skype Limited                                    *
 * Date: 060221                                                         *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"

/* Function that returns the maximum absolut value of the input vector */
#if (EMBEDDED_ARM<4)  
SKP_int16 SKP_Silk_int16_array_maxabs(    /* O    Maximum absolute value, max: 2^15-1   */
    const SKP_int16        *vec,            /* I    Input vector  [len]                   */
    const SKP_int32        len              /* I    Length of input vector                */
)                    
{
    SKP_int32 max = 0, i, lvl = 0, ind;
	if( len == 0 ) return 0;

    ind = len - 1;
    max = SKP_SMULBB( vec[ ind ], vec[ ind ] );
    for( i = len - 2; i >= 0; i-- ) {
        lvl = SKP_SMULBB( vec[ i ], vec[ i ] );
        if( lvl > max ) {
            max = lvl;
            ind = i;
        }
    }

    /* Do not return 32768, as it will not fit in an int16 so may lead to problems later on */
    if( max >= 1073676289 ) { // (2^15-1)^2 = 1073676289
        return( SKP_int16_MAX );
    } else {
        if( vec[ ind ] < 0 ) {
            return( -vec[ ind ] );
        } else {
            return(  vec[ ind ] );
        }
    }
}
#endif
