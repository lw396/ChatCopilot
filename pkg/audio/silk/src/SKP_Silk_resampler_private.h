 

/*																		*
 * File Name:	SKP_Silk_resampler_structs.h							*
 *																		*
 * Description: Structs for IIR/FIR resamplers							*
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 * All rights reserved.													*
 *																		*
 *                                                                      */

#ifndef SKP_Silk_RESAMPLER_H
#define SKP_Silk_RESAMPLER_H

#ifdef __cplusplus
extern "C" {
#endif

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_structs.h"
#include "SKP_Silk_resampler_rom.h"

/* Number of input samples to process in the inner loop */
#define RESAMPLER_MAX_BATCH_SIZE_IN             480

/* Description: Hybrid IIR/FIR polyphase implementation of resampling	*/
void SKP_Silk_resampler_private_IIR_FIR(
	void	                        *SS,		    /* I/O: Resampler state 						*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
);

/* Description: Hybrid IIR/FIR polyphase implementation of resampling	*/
void SKP_Silk_resampler_private_down_FIR(
	void	                        *SS,		    /* I/O: Resampler state 						*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
);

/* Copy */
void SKP_Silk_resampler_private_copy(
	void	                        *SS,		    /* I/O: Resampler state (unused)				*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
);

/* Upsample by a factor 2, high quality */
void SKP_Silk_resampler_private_up2_HQ_wrapper(
	void	                        *SS,		    /* I/O: Resampler state (unused)				*/
    SKP_int16                       *out,           /* O:   Output signal [ 2 * len ]               */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       len             /* I:   Number of input samples                 */
);

/* Upsample by a factor 2, high quality */
void SKP_Silk_resampler_private_up2_HQ(
	SKP_int32	                    *S,			    /* I/O: Resampler state [ 6 ]					*/
    SKP_int16                       *out,           /* O:   Output signal [ 2 * len ]               */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       len             /* I:   Number of input samples                 */
);

/* Upsample 4x, low quality */
void SKP_Silk_resampler_private_up4(
    SKP_int32                       *S,             /* I/O: State vector [ 2 ]                      */
    SKP_int16                       *out,           /* O:   Output signal [ 4 * len ]               */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       len             /* I:   Number of input samples                 */
);

/* Downsample 4x, low quality */
void SKP_Silk_resampler_private_down4(
    SKP_int32                       *S,             /* I/O: State vector [ 2 ]                      */
    SKP_int16                       *out,           /* O:   Output signal [ floor(len/2) ]          */
    const SKP_int16                 *in,            /* I:   Input signal [ len ]                    */
    SKP_int32                       inLen           /* I:   Number of input samples                 */
);

/* Second order AR filter */
void SKP_Silk_resampler_private_AR2(
	SKP_int32					    S[],		    /* I/O: State vector [ 2 ]			    	    */
	SKP_int32					    out_Q8[],		/* O:	Output signal				    	    */
	const SKP_int16				    in[],			/* I:	Input signal				    	    */
	const SKP_int16				    A_Q14[],		/* I:	AR coefficients, Q14 	                */
	SKP_int32				        len				/* I:	Signal length				        	*/
);

/* Fourth order ARMA filter */
void SKP_Silk_resampler_private_ARMA4(
	SKP_int32					    S[],		    /* I/O: State vector [ 4 ]			    	    */
	SKP_int16					    out[],		    /* O:	Output signal				    	    */
	const SKP_int16				    in[],			/* I:	Input signal				    	    */
	const SKP_int16				    Coef[],		    /* I:	ARMA coefficients [ 7 ]                 */
	SKP_int32				        len				/* I:	Signal length				        	*/
);


#ifdef __cplusplus
}
#endif
#endif // SKP_Silk_RESAMPLER_H

