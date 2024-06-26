 

#include "SKP_Silk_main_FIX.h"

#define QC  10
#define QS  14


#if EMBEDDED_ARM<6
/* Autocorrelations for a warped frequency axis */
void SKP_Silk_warped_autocorrelation_FIX(
          SKP_int32                 *corr,              /* O    Result [order + 1]                      */
          SKP_int                   *scale,             /* O    Scaling of the correlation vector       */
    const SKP_int16                 *input,             /* I    Input data to correlate                 */
    const SKP_int16                 warping_Q16,        /* I    Warping coefficient                     */
    const SKP_int                   length,             /* I    Length of input                         */
    const SKP_int                   order               /* I    Correlation order (even)                */
)
{
    SKP_int   n, i, lsh;
    SKP_int32 tmp1_QS, tmp2_QS;
    SKP_int32 state_QS[ MAX_SHAPE_LPC_ORDER + 1 ] = { 0 };
    SKP_int64 corr_QC[  MAX_SHAPE_LPC_ORDER + 1 ] = { 0 };

    /* Order must be even */
    SKP_assert( ( order & 1 ) == 0 );
    SKP_assert( 2 * QS - QC >= 0 );

    /* Loop over samples */
    for( n = 0; n < length; n++ ) {
        tmp1_QS = SKP_LSHIFT32( ( SKP_int32 )input[ n ], QS );
        /* Loop over allpass sections */
        for( i = 0; i < order; i += 2 ) {
            /* Output of allpass section */
            tmp2_QS = SKP_SMLAWB( state_QS[ i ], state_QS[ i + 1 ] - tmp1_QS, warping_Q16 );
            state_QS[ i ]  = tmp1_QS;
            corr_QC[  i ] += SKP_RSHIFT64( SKP_SMULL( tmp1_QS, state_QS[ 0 ] ), 2 * QS - QC );
            /* Output of allpass section */
            tmp1_QS = SKP_SMLAWB( state_QS[ i + 1 ], state_QS[ i + 2 ] - tmp2_QS, warping_Q16 );
            state_QS[ i + 1 ]  = tmp2_QS;
            corr_QC[  i + 1 ] += SKP_RSHIFT64( SKP_SMULL( tmp2_QS, state_QS[ 0 ] ), 2 * QS - QC );
        }
        state_QS[ order ] = tmp1_QS;
        corr_QC[  order ] += SKP_RSHIFT64( SKP_SMULL( tmp1_QS, state_QS[ 0 ] ), 2 * QS - QC );
    }

    lsh = SKP_Silk_CLZ64( corr_QC[ 0 ] ) - 35;
    lsh = SKP_LIMIT( lsh, -12 - QC, 30 - QC );
    *scale = -( QC + lsh ); 
    SKP_assert( *scale >= -30 && *scale <= 12 );
    if( lsh >= 0 ) {
        for( i = 0; i < order + 1; i++ ) {
            corr[ i ] = ( SKP_int32 )SKP_CHECK_FIT32( SKP_LSHIFT64( corr_QC[ i ], lsh ) );
        }
    } else {
        for( i = 0; i < order + 1; i++ ) {
            corr[ i ] = ( SKP_int32 )SKP_CHECK_FIT32( SKP_RSHIFT64( corr_QC[ i ], -lsh ) );
        }    
    }
    SKP_assert( corr_QC[ 0 ] >= 0 ); // If breaking, decrease QC
}
#endif

