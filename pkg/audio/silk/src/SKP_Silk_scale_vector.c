 

#include "SKP_Silk_SigProc_FIX.h"

/* Multiply a vector by a constant */
void SKP_Silk_scale_vector32_Q26_lshift_18( 
    SKP_int32           *data1,                     /* (I/O): Q0/Q18        */
    SKP_int32           gain_Q26,                   /* (I):   Q26           */
    SKP_int             dataSize                    /* (I):   length        */
)
{
    SKP_int  i;

    for( i = 0; i < dataSize; i++ ) {
        data1[ i ] = (SKP_int32)SKP_CHECK_FIT32( SKP_RSHIFT64( SKP_SMULL( data1[ i ], gain_Q26 ), 8 ) );// OUTPUT: Q18
    }
}

