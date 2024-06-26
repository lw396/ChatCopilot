 

/*                                                                      *
 * SKP_Silk_sum_sqr_shift.c                                           *
 *                                                                      *
 * compute number of bits to right shift the sum of squares of a vector *
 * of int16s to make it fit in an int32                                 *
 *                                                                      *
 * Copyright 2006-2008 (c), Skype Limited                               *
 *                                                                      */
#include "SKP_Silk_SigProc_FIX.h"
#if (EMBEDDED_ARM<5) 
/* Compute number of bits to right shift the sum of squares of a vector */
/* of int16s to make it fit in an int32                                 */
void SKP_Silk_sum_sqr_shift(
    SKP_int32            *energy,            /* O    Energy of x, after shifting to the right            */
    SKP_int              *shift,             /* O    Number of bits right shift applied to energy        */
    const SKP_int16      *x,                 /* I    Input vector                                        */
    SKP_int              len                 /* I    Length of input vector                              */
)
{
    SKP_int   i, shft;
    SKP_int32 in32, nrg_tmp, nrg;

    if( (SKP_int32)( (SKP_int_ptr_size)x & 2 ) != 0 ) {
        /* Input is not 4-byte aligned */
        nrg = SKP_SMULBB( x[ 0 ], x[ 0 ] );
        i = 1;
    } else {
        nrg = 0;
        i   = 0;
    }
    shft = 0;
    len--;
    while( i < len ) {
        /* Load two values at once */
        in32 = *( (SKP_int32 *)&x[ i ] );
        nrg = SKP_SMLABB_ovflw( nrg, in32, in32 );
        nrg = SKP_SMLATT_ovflw( nrg, in32, in32 );
        i += 2;
        if( nrg < 0 ) {
            /* Scale down */
            nrg = (SKP_int32)SKP_RSHIFT_uint( (SKP_uint32)nrg, 2 );
            shft = 2;
            break;
        }
    }
    for( ; i < len; i += 2 ) {
        /* Load two values at once */
        in32 = *( (SKP_int32 *)&x[ i ] );
        nrg_tmp = SKP_SMULBB( in32, in32 );
        nrg_tmp = SKP_SMLATT_ovflw( nrg_tmp, in32, in32 );
        nrg = (SKP_int32)SKP_ADD_RSHIFT_uint( nrg, (SKP_uint32)nrg_tmp, shft );
        if( nrg < 0 ) {
            /* Scale down */
            nrg = (SKP_int32)SKP_RSHIFT_uint( (SKP_uint32)nrg, 2 );
            shft += 2;
        }
    }
    if( i == len ) {
        /* One sample left to process */
        nrg_tmp = SKP_SMULBB( x[ i ], x[ i ] );
        nrg = (SKP_int32)SKP_ADD_RSHIFT_uint( nrg, nrg_tmp, shft );
    }

    /* Make sure to have at least one extra leading zero (two leading zeros in total) */
    if( nrg & 0xC0000000 ) {
        nrg = SKP_RSHIFT_uint( (SKP_uint32)nrg, 2 );
        shft += 2;
    }

    /* Output arguments */
    *shift  = shft;
    *energy = nrg;
}

#endif
