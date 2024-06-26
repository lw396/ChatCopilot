 

#include "SKP_Silk_SigProc_FIX.h"

/* Copy and multiply a vector by a constant */
void SKP_Silk_scale_copy_vector16( 
    SKP_int16           *data_out, 
    const SKP_int16     *data_in, 
    SKP_int32           gain_Q16,                   /* (I):   gain in Q16   */
    const SKP_int       dataSize                    /* (I):   length        */
)
{
    SKP_int  i;
    SKP_int32 tmp32;

    for( i = 0; i < dataSize; i++ ) {
        tmp32 = SKP_SMULWB( gain_Q16, data_in[ i ] );
        data_out[ i ] = (SKP_int16)SKP_CHECK_FIT16( tmp32 );
    }
}
