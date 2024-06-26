 

/*																		*
 * SKP_Silk_resampler_private_AR2. c                                  *
 *																		*
 * Second order AR filter with single delay elements                	*
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

#if (EMBEDDED_ARM<5)  
/* Second order AR filter with single delay elements */
void SKP_Silk_resampler_private_AR2(
	SKP_int32					    S[],		    /* I/O: State vector [ 2 ]			    	    */
	SKP_int32					    out_Q8[],		/* O:	Output signal				    	    */
	const SKP_int16				    in[],			/* I:	Input signal				    	    */
	const SKP_int16				    A_Q14[],		/* I:	AR coefficients, Q14 	                */
	SKP_int32				        len				/* I:	Signal length				        	*/
)
{
	SKP_int32	k;
	SKP_int32	out32;

	for( k = 0; k < len; k++ ) {
		out32       = SKP_ADD_LSHIFT32( S[ 0 ], (SKP_int32)in[ k ], 8 );
		out_Q8[ k ] = out32;
		out32       = SKP_LSHIFT( out32, 2 );
		S[ 0 ]      = SKP_SMLAWB( S[ 1 ], out32, A_Q14[ 0 ] );
		S[ 1 ]      = SKP_SMULWB( out32, A_Q14[ 1 ] );
	}
}
#endif
