 

#include "SKP_Silk_main_FIX.h"

/* Add noise to matrix diagonal */
void SKP_Silk_regularize_correlations_FIX(
    SKP_int32                       *XX,                /* I/O  Correlation matrices                        */
    SKP_int32                       *xx,                /* I/O  Correlation values                          */
    SKP_int32                       noise,              /* I    Noise to add                                */
    SKP_int                         D                   /* I    Dimension of XX                             */
)
{
    SKP_int i;
    for( i = 0; i < D; i++ ) {
        matrix_ptr( &XX[ 0 ], i, i, D ) = SKP_ADD32( matrix_ptr( &XX[ 0 ], i, i, D ), noise );
    }
    xx[ 0 ] += noise;
}
