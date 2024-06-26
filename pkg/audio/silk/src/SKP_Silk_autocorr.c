
#include "SKP_Silk_SigProc_FIX.h"

/* Compute autocorrelation */
void SKP_Silk_autocorr( 
    SKP_int32        *results,                   /* O    Result (length correlationCount)            */
    SKP_int          *scale,                     /* O    Scaling of the correlation vector           */
    const SKP_int16  *inputData,                 /* I    Input data to correlate                     */
    const SKP_int    inputDataSize,              /* I    Length of input                             */
    const SKP_int    correlationCount            /* I    Number of correlation taps to compute       */
)
{
    SKP_int   i, lz, nRightShifts, corrCount;
    SKP_int64 corr64;

    corrCount = SKP_min_int( inputDataSize, correlationCount );

    /* compute energy (zero-lag correlation) */
    corr64 = SKP_Silk_inner_prod16_aligned_64( inputData, inputData, inputDataSize );

    /* deal with all-zero input data */
    corr64 += 1;

    /* number of leading zeros */
    lz = SKP_Silk_CLZ64( corr64 );

    /* scaling: number of right shifts applied to correlations */
    nRightShifts = 35 - lz;
    *scale = nRightShifts;

    if( nRightShifts <= 0 ) {
        results[ 0 ] = SKP_LSHIFT( (SKP_int32)SKP_CHECK_FIT32( corr64 ), -nRightShifts );

        /* compute remaining correlations based on int32 inner product */
          for( i = 1; i < corrCount; i++ ) {
            results[ i ] = SKP_LSHIFT( SKP_Silk_inner_prod_aligned( inputData, inputData + i, inputDataSize - i ), -nRightShifts );
        }
    } else {
        results[ 0 ] = (SKP_int32)SKP_CHECK_FIT32( SKP_RSHIFT64( corr64, nRightShifts ) );

        /* compute remaining correlations based on int64 inner product */
          for( i = 1; i < corrCount; i++ ) {
            results[ i ] =  (SKP_int32)SKP_CHECK_FIT32( SKP_RSHIFT64( SKP_Silk_inner_prod16_aligned_64( inputData, inputData + i, inputDataSize - i ), nRightShifts ) );
        }
    }
}
