 

#include "SKP_Silk_main.h"

/* Interpolate two vectors */
void SKP_Silk_interpolate(
    SKP_int                         xi[ MAX_LPC_ORDER ],    /* O    interpolated vector                     */
    const SKP_int                   x0[ MAX_LPC_ORDER ],    /* I    first vector                            */
    const SKP_int                   x1[ MAX_LPC_ORDER ],    /* I    second vector                           */
    const SKP_int                   ifact_Q2,               /* I    interp. factor, weight on 2nd vector    */
    const SKP_int                   d                       /* I    number of parameters                    */
)
{
    SKP_int i;

    SKP_assert( ifact_Q2 >= 0 );
    SKP_assert( ifact_Q2 <= ( 1 << 2 ) );

    for( i = 0; i < d; i++ ) {
        xi[ i ] = ( SKP_int )( ( SKP_int32 )x0[ i ] + SKP_RSHIFT( SKP_MUL( ( SKP_int32 )x1[ i ] - ( SKP_int32 )x0[ i ], ifact_Q2 ), 2 ) );
    }
}
