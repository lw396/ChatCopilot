 

#include "SKP_Silk_main_FIX.h"

/* Residual energy: nrg = wxx - 2 * wXx * c + c' * wXX * c */
SKP_int32 SKP_Silk_residual_energy16_covar_FIX(
    const SKP_int16                 *c,                 /* I    Prediction vector                           */
    const SKP_int32                 *wXX,               /* I    Correlation matrix                          */
    const SKP_int32                 *wXx,               /* I    Correlation vector                          */
    SKP_int32                       wxx,                /* I    Signal energy                               */
    SKP_int                         D,                  /* I    Dimension                                   */
    SKP_int                         cQ                  /* I    Q value for c vector 0 - 15                 */
)
{
    SKP_int   i, j, lshifts, Qxtra;
    SKP_int32 c_max, w_max, tmp, tmp2, nrg;
    SKP_int   cn[ MAX_MATRIX_SIZE ]; 
    const SKP_int32 *pRow;

    /* Safety checks */
    SKP_assert( D >=  0 );
    SKP_assert( D <= 16 );
    SKP_assert( cQ >  0 );
    SKP_assert( cQ < 16 );

    lshifts = 16 - cQ;
    Qxtra = lshifts;

    c_max = 0;
    for( i = 0; i < D; i++ ) {
        c_max = SKP_max_32( c_max, SKP_abs( ( SKP_int32 )c[ i ] ) );
    }
    Qxtra = SKP_min_int( Qxtra, SKP_Silk_CLZ32( c_max ) - 17 );

    w_max = SKP_max_32( wXX[ 0 ], wXX[ D * D - 1 ] );
    Qxtra = SKP_min_int( Qxtra, SKP_Silk_CLZ32( SKP_MUL( D, SKP_RSHIFT( SKP_SMULWB( w_max, c_max ), 4 ) ) ) - 5 );
    Qxtra = SKP_max_int( Qxtra, 0 );
    for( i = 0; i < D; i++ ) {
        cn[ i ] = SKP_LSHIFT( ( SKP_int )c[ i ], Qxtra );
        SKP_assert( SKP_abs(cn[i]) <= ( SKP_int16_MAX + 1 ) ); /* Check that SKP_SMLAWB can be used */
    }
    lshifts -= Qxtra;

    /* Compute wxx - 2 * wXx * c */
    tmp = 0;
    for( i = 0; i < D; i++ ) {
        tmp = SKP_SMLAWB( tmp, wXx[ i ], cn[ i ] );
    }
    nrg = SKP_RSHIFT( wxx, 1 + lshifts ) - tmp;                         /* Q: -lshifts - 1 */

    /* Add c' * wXX * c, assuming wXX is symmetric */
    tmp2 = 0;
    for( i = 0; i < D; i++ ) {
        tmp = 0;
        pRow = &wXX[ i * D ];
        for( j = i + 1; j < D; j++ ) {
            tmp = SKP_SMLAWB( tmp, pRow[ j ], cn[ j ] );
        }
        tmp  = SKP_SMLAWB( tmp,  SKP_RSHIFT( pRow[ i ], 1 ), cn[ i ] );
        tmp2 = SKP_SMLAWB( tmp2, tmp,                        cn[ i ] );
    }
    nrg = SKP_ADD_LSHIFT32( nrg, tmp2, lshifts );                       /* Q: -lshifts - 1 */

    /* Keep one bit free always, because we add them for LSF interpolation */
    if( nrg < 1 ) {
        nrg = 1;
    } else if( nrg > SKP_RSHIFT( SKP_int32_MAX, lshifts + 2 ) ) {
        nrg = SKP_int32_MAX >> 1;
    } else {
        nrg = SKP_LSHIFT( nrg, lshifts + 1 );                           /* Q0 */
    }
    return nrg;

}
