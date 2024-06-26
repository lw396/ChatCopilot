 

#include "SKP_Silk_main_FIX.h"

/* Rate-Distortion calculations for multiple input data vectors */
void SKP_Silk_NLSF_VQ_rate_distortion_FIX(
    SKP_int32                       *pRD_Q20,           /* O    Rate-distortion values [psNLSF_CBS->nVectors*N] */
    const SKP_Silk_NLSF_CBS         *psNLSF_CBS,        /* I    NLSF codebook stage struct                      */
    const SKP_int                   *in_Q15,            /* I    Input vectors to be quantized                   */
    const SKP_int                   *w_Q6,              /* I    Weight vector                                   */
    const SKP_int32                 *rate_acc_Q5,       /* I    Accumulated rates from previous stage           */
    const SKP_int                   mu_Q15,             /* I    Weight between weighted error and rate          */
    const SKP_int                   N,                  /* I    Number of input vectors to be quantized         */
    const SKP_int                   LPC_order           /* I    LPC order                                       */
)
{
    SKP_int   i, n;
    SKP_int32 *pRD_vec_Q20;

    /* Compute weighted quantization errors for all input vectors over one codebook stage */
    SKP_Silk_NLSF_VQ_sum_error_FIX( pRD_Q20, in_Q15, w_Q6, psNLSF_CBS->CB_NLSF_Q15, 
        N, psNLSF_CBS->nVectors, LPC_order );

    /* Loop over input vectors */
    pRD_vec_Q20 = pRD_Q20;
    for( n = 0; n < N; n++ ) {
        /* Add rate cost to error for each codebook vector */
        for( i = 0; i < psNLSF_CBS->nVectors; i++ ) {
            SKP_assert( rate_acc_Q5[ n ] + psNLSF_CBS->Rates_Q5[ i ] >= 0 );
            SKP_assert( rate_acc_Q5[ n ] + psNLSF_CBS->Rates_Q5[ i ] <= SKP_int16_MAX );
            pRD_vec_Q20[ i ] = SKP_SMLABB( pRD_vec_Q20[ i ], rate_acc_Q5[ n ] + psNLSF_CBS->Rates_Q5[ i ], mu_Q15 );
            SKP_assert( pRD_vec_Q20[ i ] >= 0 );
        }
        pRD_vec_Q20 += psNLSF_CBS->nVectors;
    }
}
