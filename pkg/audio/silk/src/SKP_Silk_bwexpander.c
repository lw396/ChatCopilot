 

#include "SKP_Silk_SigProc_FIX.h"

/* Chirp (bandwidth expand) LP AR filter */
void SKP_Silk_bwexpander( 
    SKP_int16            *ar,        /* I/O  AR filter to be expanded (without leading 1)    */
    const SKP_int        d,          /* I    Length of ar                                    */
    SKP_int32            chirp_Q16   /* I    Chirp factor (typically in the range 0 to 1)    */
)
{
    SKP_int   i;
    SKP_int32 chirp_minus_one_Q16;

    chirp_minus_one_Q16 = chirp_Q16 - 65536;

    /* NB: Dont use SKP_SMULWB, instead of SKP_RSHIFT_ROUND( SKP_MUL() , 16 ), below. */
    /* Bias in SKP_SMULWB can lead to unstable filters                                */
    for( i = 0; i < d - 1; i++ ) {
        ar[ i ]    = (SKP_int16)SKP_RSHIFT_ROUND( SKP_MUL( chirp_Q16, ar[ i ]             ), 16 );
        chirp_Q16 +=            SKP_RSHIFT_ROUND( SKP_MUL( chirp_Q16, chirp_minus_one_Q16 ), 16 );
    }
    ar[ d - 1 ] = (SKP_int16)SKP_RSHIFT_ROUND( SKP_MUL( chirp_Q16, ar[ d - 1 ] ), 16 );
}
