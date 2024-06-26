 

#include "SKP_Silk_main_FIX.h"

#if (!defined(__mips__)) && (EMBEDDED_ARM < 6)

/* Compute weighted quantization errors for an LPC_order element input vector, over one codebook stage */
void SKP_Silk_NLSF_VQ_sum_error_FIX(
    SKP_int32                       *err_Q20,           /* O    Weighted quantization errors  [N*K]         */
    const SKP_int                   *in_Q15,            /* I    Input vectors to be quantized [N*LPC_order] */
    const SKP_int                   *w_Q6,              /* I    Weighting vectors             [N*LPC_order] */
    const SKP_int16                 *pCB_Q15,           /* I    Codebook vectors              [K*LPC_order] */
    const SKP_int                   N,                  /* I    Number of input vectors                     */
    const SKP_int                   K,                  /* I    Number of codebook vectors                  */
    const SKP_int                   LPC_order           /* I    Number of LPCs                              */
)
{
    SKP_int         i, n, m;
    SKP_int32       diff_Q15, sum_error, Wtmp_Q6;
    SKP_int32       Wcpy_Q6[ MAX_LPC_ORDER / 2 ];
    const SKP_int16 *cb_vec_Q15;

    SKP_assert( LPC_order <= 16 );
    SKP_assert( ( LPC_order & 1 ) == 0 );

    /* Copy to local stack and pack two weights per int32 */
    for( m = 0; m < SKP_RSHIFT( LPC_order, 1 ); m++ ) {
        Wcpy_Q6[ m ] = w_Q6[ 2 * m ] | SKP_LSHIFT( ( SKP_int32 )w_Q6[ 2 * m + 1 ], 16 );
    }

    /* Loop over input vectors */
    for( n = 0; n < N; n++ ) {
        /* Loop over codebook */
        cb_vec_Q15 = pCB_Q15;
        for( i = 0; i < K; i++ ) {
            sum_error = 0;
            for( m = 0; m < LPC_order; m += 2 ) {
                /* Get two weights packed in an int32 */
                Wtmp_Q6 = Wcpy_Q6[ SKP_RSHIFT( m, 1 ) ];

                /* Compute weighted squared quantization error for index m */
                diff_Q15 = in_Q15[ m ] - *cb_vec_Q15++; // range: [ -32767 : 32767 ]
                sum_error = SKP_SMLAWB( sum_error, SKP_SMULBB( diff_Q15, diff_Q15 ), Wtmp_Q6 );

                /* Compute weighted squared quantization error for index m + 1 */
                diff_Q15 = in_Q15[m + 1] - *cb_vec_Q15++; // range: [ -32767 : 32767 ]
                sum_error = SKP_SMLAWT( sum_error, SKP_SMULBB( diff_Q15, diff_Q15 ), Wtmp_Q6 );
            }
            SKP_assert( sum_error >= 0 );
            err_Q20[ i ] = sum_error;
        }
        err_Q20 += K;
        in_Q15 += LPC_order;
    }
}

#endif

