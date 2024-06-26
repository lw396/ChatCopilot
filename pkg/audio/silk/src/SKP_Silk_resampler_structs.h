 

/*																		*
 * File Name:	SKP_Silk_resampler_structs.h							*
 *																		*
 * Description: Structs for IIR/FIR resamplers							*
 *                                                                      *
 * Copyright 2010 (c), Skype Limited                                    *
 * All rights reserved.													*
 *																		*
 *                                                                      */

#ifndef SKP_Silk_RESAMPLER_STRUCTS_H
#define SKP_Silk_RESAMPLER_STRUCTS_H

#ifdef __cplusplus
extern "C" {
#endif

/* Flag to enable support for input/output sampling rates above 48 kHz. Turn off for embedded devices */
#define RESAMPLER_SUPPORT_ABOVE_48KHZ                   1

#define SKP_Silk_RESAMPLER_MAX_FIR_ORDER                 16
#define SKP_Silk_RESAMPLER_MAX_IIR_ORDER                 6


typedef struct _SKP_Silk_resampler_state_struct{
	SKP_int32       sIIR[ SKP_Silk_RESAMPLER_MAX_IIR_ORDER ];        /* this must be the first element of this struct */
	SKP_int32       sFIR[ SKP_Silk_RESAMPLER_MAX_FIR_ORDER ];
	SKP_int32       sDown2[ 2 ];
	void            (*resampler_function)( void *, SKP_int16 *, const SKP_int16 *, SKP_int32 );
	void            (*up2_function)(  SKP_int32 *, SKP_int16 *, const SKP_int16 *, SKP_int32 );
    SKP_int32       batchSize;
	SKP_int32       invRatio_Q16;
	SKP_int32       FIR_Fracs;
    SKP_int32       input2x;
	const SKP_int16	*Coefs;
#if RESAMPLER_SUPPORT_ABOVE_48KHZ
	SKP_int32       sDownPre[ 2 ];
	SKP_int32       sUpPost[ 2 ];
	void            (*down_pre_function)( SKP_int32 *, SKP_int16 *, const SKP_int16 *, SKP_int32 );
	void            (*up_post_function)(  SKP_int32 *, SKP_int16 *, const SKP_int16 *, SKP_int32 );
	SKP_int32       batchSizePrePost;
	SKP_int32       ratio_Q16;
	SKP_int32       nPreDownsamplers;
	SKP_int32       nPostUpsamplers;
#endif
	SKP_int32 magic_number;
} SKP_Silk_resampler_state_struct;

#ifdef __cplusplus
}
#endif
#endif /* SKP_Silk_RESAMPLER_STRUCTS_H */

