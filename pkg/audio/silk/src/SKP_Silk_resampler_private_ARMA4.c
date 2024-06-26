 

/*																		*
 * SKP_Silk_resampler_private_ARMA4.c                                 *
 *																		*
 * Fourth order ARMA filter, applies 64x gain                           *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

/* Fourth order ARMA filter                                             */
/* Internally operates as two biquad filters in sequence.               */

/* Coeffients are stored in a packed format:                                                        */
/*    { B1_Q14[1], B2_Q14[1], -A1_Q14[1], -A1_Q14[2], -A2_Q14[1], -A2_Q14[2], gain_Q16 }            */
/* where it is assumed that B*_Q14[0], B*_Q14[2], A*_Q14[0] are all 16384                           */
#if (EMBEDDED_ARM<5) 
void SKP_Silk_resampler_private_ARMA4(
	SKP_int32					    S[],		    /* I/O: State vector [ 4 ]			    	    */
	SKP_int16					    out[],		    /* O:	Output signal				    	    */
	const SKP_int16				    in[],			/* I:	Input signal				    	    */
	const SKP_int16				    Coef[],		    /* I:	ARMA coefficients [ 7 ]                 */
	SKP_int32				        len				/* I:	Signal length				        	*/
)
{
	SKP_int32 k;
	SKP_int32 in_Q8, out1_Q8, out2_Q8, X;

	for( k = 0; k < len; k++ ) {
        in_Q8  = SKP_LSHIFT32( (SKP_int32)in[ k ], 8 );

        /* Outputs of first and second biquad */
        out1_Q8 = SKP_ADD_LSHIFT32( in_Q8,   S[ 0 ], 2 );
        out2_Q8 = SKP_ADD_LSHIFT32( out1_Q8, S[ 2 ], 2 );

        /* Update states, which are stored in Q6. Coefficients are in Q14 here */
        X      = SKP_SMLAWB( S[ 1 ], in_Q8,   Coef[ 0 ] );
        S[ 0 ] = SKP_SMLAWB( X,      out1_Q8, Coef[ 2 ] );

        X      = SKP_SMLAWB( S[ 3 ], out1_Q8, Coef[ 1 ] );
        S[ 2 ] = SKP_SMLAWB( X,      out2_Q8, Coef[ 4 ] );

        S[ 1 ] = SKP_SMLAWB( SKP_RSHIFT32( in_Q8,   2 ), out1_Q8, Coef[ 3 ] );
        S[ 3 ] = SKP_SMLAWB( SKP_RSHIFT32( out1_Q8, 2 ), out2_Q8, Coef[ 5 ] );

        /* Apply gain and store to output. The coefficient is in Q16 */
        out[ k ] = (SKP_int16)SKP_SAT16( SKP_RSHIFT32( SKP_SMLAWB( 128, out2_Q8, Coef[ 6 ] ), 8 ) );
	}
}
#endif

