 

#include "SKP_Silk_main.h"

/* Convert NLSF parameters to stable AR prediction filter coefficients */
void SKP_Silk_NLSF2A_stable(
    SKP_int16                       pAR_Q12[ MAX_LPC_ORDER ],   /* O    Stabilized AR coefs [LPC_order]     */ 
    const SKP_int                   pNLSF[ MAX_LPC_ORDER ],     /* I    NLSF vector         [LPC_order]     */
    const SKP_int                   LPC_order                   /* I    LPC/LSF order                       */
)
{
    SKP_int   i;
    SKP_int32 invGain_Q30;

    SKP_Silk_NLSF2A( pAR_Q12, pNLSF, LPC_order );

    /* Ensure stable LPCs */
    for( i = 0; i < MAX_LPC_STABILIZE_ITERATIONS; i++ ) {
        if( SKP_Silk_LPC_inverse_pred_gain( &invGain_Q30, pAR_Q12, LPC_order ) == 1 ) {
            SKP_Silk_bwexpander( pAR_Q12, LPC_order, 65536 - SKP_SMULBB( 10 + i, i ) );		/* 10_Q16 = 0.00015 */
        } else {
            break;
        }
    }

    /* Reached the last iteration */
    if( i == MAX_LPC_STABILIZE_ITERATIONS ) {
        SKP_assert( 0 );
        for( i = 0; i < LPC_order; i++ ) {
            pAR_Q12[ i ] = 0;
        }
    }
}
