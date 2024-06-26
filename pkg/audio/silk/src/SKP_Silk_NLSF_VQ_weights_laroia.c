 

#include "SKP_Silk_SigProc_FIX.h"

/* 
R. Laroia, N. Phamdo and N. Farvardin, "Robust and Efficient Quantization of Speech LSP
Parameters Using Structured Vector Quantization", Proc. IEEE Int. Conf. Acoust., Speech,
Signal Processing, pp. 641-644, 1991.
*/

#define Q_OUT                       6
#define MIN_NDELTA                  3

/* Laroia low complexity NLSF weights */
void SKP_Silk_NLSF_VQ_weights_laroia(
    SKP_int             *pNLSFW_Q6,         /* O: Pointer to input vector weights           [D x 1]     */
    const SKP_int       *pNLSF_Q15,         /* I: Pointer to input vector                   [D x 1]     */ 
    const SKP_int       D                   /* I: Input vector dimension (even)                         */
)
{
    SKP_int   k;
    SKP_int32 tmp1_int, tmp2_int;
    
    /* Check that we are guaranteed to end up within the required range */
    SKP_assert( D > 0 );
    SKP_assert( ( D & 1 ) == 0 );
    
    /* First value */
    tmp1_int = SKP_max_int( pNLSF_Q15[ 0 ], MIN_NDELTA );
    tmp1_int = SKP_DIV32_16( 1 << ( 15 + Q_OUT ), tmp1_int );
    tmp2_int = SKP_max_int( pNLSF_Q15[ 1 ] - pNLSF_Q15[ 0 ], MIN_NDELTA );
    tmp2_int = SKP_DIV32_16( 1 << ( 15 + Q_OUT ), tmp2_int );
    pNLSFW_Q6[ 0 ] = (SKP_int)SKP_min_int( tmp1_int + tmp2_int, SKP_int16_MAX );
    SKP_assert( pNLSFW_Q6[ 0 ] > 0 );
    
    /* Main loop */
    for( k = 1; k < D - 1; k += 2 ) {
        tmp1_int = SKP_max_int( pNLSF_Q15[ k + 1 ] - pNLSF_Q15[ k ], MIN_NDELTA );
        tmp1_int = SKP_DIV32_16( 1 << ( 15 + Q_OUT ), tmp1_int );
        pNLSFW_Q6[ k ] = (SKP_int)SKP_min_int( tmp1_int + tmp2_int, SKP_int16_MAX );
        SKP_assert( pNLSFW_Q6[ k ] > 0 );

        tmp2_int = SKP_max_int( pNLSF_Q15[ k + 2 ] - pNLSF_Q15[ k + 1 ], MIN_NDELTA );
        tmp2_int = SKP_DIV32_16( 1 << ( 15 + Q_OUT ), tmp2_int );
        pNLSFW_Q6[ k + 1 ] = (SKP_int)SKP_min_int( tmp1_int + tmp2_int, SKP_int16_MAX );
        SKP_assert( pNLSFW_Q6[ k + 1 ] > 0 );
    }
    
    /* Last value */
    tmp1_int = SKP_max_int( ( 1 << 15 ) - pNLSF_Q15[ D - 1 ], MIN_NDELTA );
    tmp1_int = SKP_DIV32_16( 1 << ( 15 + Q_OUT ), tmp1_int );
    pNLSFW_Q6[ D - 1 ] = (SKP_int)SKP_min_int( tmp1_int + tmp2_int, SKP_int16_MAX );
    SKP_assert( pNLSFW_Q6[ D - 1 ] > 0 );
}
