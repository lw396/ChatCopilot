 

/*																		*
 * File Name:	SKP_Silk_resampler_private_copy.c                     *
 *																		*
 * Description: Copy.                                                   *
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 * All rights reserved.													*
 *                                                                      */

#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_resampler_private.h"

/* Copy */
void SKP_Silk_resampler_private_copy(
	void	                        *SS,		    /* I/O: Resampler state (unused)				*/
	SKP_int16						out[],		    /* O:	Output signal 							*/
	const SKP_int16					in[],		    /* I:	Input signal							*/
	SKP_int32					    inLen		    /* I:	Number of input samples					*/
)
{
    SKP_memcpy( out, in, inLen * sizeof( SKP_int16 ) );
}
