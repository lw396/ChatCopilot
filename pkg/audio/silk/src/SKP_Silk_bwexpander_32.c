 

#include "SKP_Silk_SigProc_FIX.h"

/* Chirp (bandwidth expand) LP AR filter */
void SKP_Silk_bwexpander_32( 
    SKP_int32        *ar,      /* I/O    AR filter to be expanded (without leading 1)    */
    const SKP_int    d,        /* I    Length of ar                                      */
    SKP_int32        chirp_Q16 /* I    Chirp factor in Q16                               */
)
{
    SKP_int   i;
    SKP_int32 tmp_chirp_Q16;

    tmp_chirp_Q16 = chirp_Q16;
    for( i = 0; i < d - 1; i++ ) {
        ar[ i ]       = SKP_SMULWW( ar[ i ],   tmp_chirp_Q16 );
        tmp_chirp_Q16 = SKP_SMULWW( chirp_Q16, tmp_chirp_Q16 );
    }
    ar[ d - 1 ] = SKP_SMULWW( ar[ d - 1 ], tmp_chirp_Q16 );
}
