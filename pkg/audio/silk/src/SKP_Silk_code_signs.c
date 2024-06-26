 

#include "SKP_Silk_main.h"

//#define SKP_enc_map(a)                ((a) > 0 ? 1 : 0)
//#define SKP_dec_map(a)                ((a) > 0 ? 1 : -1)
/* shifting avoids if-statement */
#define SKP_enc_map(a)                  ( SKP_RSHIFT( (a), 15 ) + 1 )
#define SKP_dec_map(a)                  ( SKP_LSHIFT( (a),  1 ) - 1 )

/* Encodes signs of excitation */
void SKP_Silk_encode_signs(
    SKP_Silk_range_coder_state      *sRC,               /* I/O  Range coder state                       */
    const SKP_int8                  q[],                /* I    Pulse signal                            */
    const SKP_int                   length,             /* I    Length of input                         */
    const SKP_int                   sigtype,            /* I    Signal type                             */
    const SKP_int                   QuantOffsetType,    /* I    Quantization offset type                */
    const SKP_int                   RateLevelIndex      /* I    Rate level index                        */
)
{
    SKP_int i;
    SKP_int inData;
    SKP_uint16 cdf[ 3 ];

    i = SKP_SMULBB( N_RATE_LEVELS - 1, SKP_LSHIFT( sigtype, 1 ) + QuantOffsetType ) + RateLevelIndex;
    cdf[ 0 ] = 0;
    cdf[ 1 ] = SKP_Silk_sign_CDF[ i ];
    cdf[ 2 ] = 65535;
    
    for( i = 0; i < length; i++ ) {
        if( q[ i ] != 0 ) {
            inData = SKP_enc_map( q[ i ] ); /* - = 0, + = 1 */
            SKP_Silk_range_encoder( sRC, inData, cdf );
        }
    }
}

/* Decodes signs of excitation */
void SKP_Silk_decode_signs(
    SKP_Silk_range_coder_state      *sRC,               /* I/O  Range coder state                           */
    SKP_int                         q[],                /* I/O  pulse signal                                */
    const SKP_int                   length,             /* I    length of output                            */
    const SKP_int                   sigtype,            /* I    Signal type                                 */
    const SKP_int                   QuantOffsetType,    /* I    Quantization offset type                    */
    const SKP_int                   RateLevelIndex      /* I    Rate Level Index                            */
)
{
    SKP_int i;
    SKP_int data;
    SKP_uint16 cdf[ 3 ];

    i = SKP_SMULBB( N_RATE_LEVELS - 1, SKP_LSHIFT( sigtype, 1 ) + QuantOffsetType ) + RateLevelIndex;
    cdf[ 0 ] = 0;
    cdf[ 1 ] = SKP_Silk_sign_CDF[ i ];
    cdf[ 2 ] = 65535;
    
    for( i = 0; i < length; i++ ) {
        if( q[ i ] > 0 ) {
            SKP_Silk_range_decoder( &data, sRC, cdf, 1 );
            /* attach sign */
            /* implementation with shift, subtraction, multiplication */
            q[ i ] *= SKP_dec_map( data );
        }
    }
}

