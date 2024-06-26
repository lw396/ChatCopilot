 

/*                                                                      *
 * File Name:    SKP_Silk_resample_rom.h                              *
 *                                                                      *
 * Description: Header file for FIR resampling of                       *
 *                32 and 44 kHz input                                   *
 *                                                                      *
 * Copyright 2007 (c), Skype Limited                                    *
 * All rights reserved.                                                 *
 *                                                                      *
 * Date: 070807                                                         *
 *                                                                      */

#ifndef _SKP_SILK_FIX_RESAMPLER_ROM_H_
#define _SKP_SILK_FIX_RESAMPLER_ROM_H_

#ifdef  __cplusplus
extern "C"
{
#endif

#include "SKP_Silk_typedef.h"
#include "SKP_Silk_resampler_structs.h"

#define RESAMPLER_DOWN_ORDER_FIR                12
#define RESAMPLER_ORDER_FIR_144                 6


/* Tables for 2x downsampler. Values above 32767 intentionally wrap to a negative value. */
extern const SKP_int16 SKP_Silk_resampler_down2_0;
extern const SKP_int16 SKP_Silk_resampler_down2_1;

/* Tables for 2x upsampler, low quality. Values above 32767 intentionally wrap to a negative value. */
extern const SKP_int16 SKP_Silk_resampler_up2_lq_0;
extern const SKP_int16 SKP_Silk_resampler_up2_lq_1;

/* Tables for 2x upsampler, high quality. Values above 32767 intentionally wrap to a negative value. */
extern const SKP_int16 SKP_Silk_resampler_up2_hq_0[ 2 ];
extern const SKP_int16 SKP_Silk_resampler_up2_hq_1[ 2 ];
extern const SKP_int16 SKP_Silk_resampler_up2_hq_notch[ 4 ];

/* Tables with IIR and FIR coefficients for fractional downsamplers */
extern const SKP_int16 SKP_Silk_Resampler_3_4_COEFS[ 2 + 3 * RESAMPLER_DOWN_ORDER_FIR / 2 ];
extern const SKP_int16 SKP_Silk_Resampler_2_3_COEFS[ 2 + 2 * RESAMPLER_DOWN_ORDER_FIR / 2 ];
extern const SKP_int16 SKP_Silk_Resampler_1_2_COEFS[ 2 +     RESAMPLER_DOWN_ORDER_FIR / 2 ];
extern const SKP_int16 SKP_Silk_Resampler_3_8_COEFS[ 2 + 3 * RESAMPLER_DOWN_ORDER_FIR / 2 ];
extern const SKP_int16 SKP_Silk_Resampler_1_3_COEFS[ 2 +     RESAMPLER_DOWN_ORDER_FIR / 2 ];
extern const SKP_int16 SKP_Silk_Resampler_2_3_COEFS_LQ[ 2 + 2 * 2 ];
extern const SKP_int16 SKP_Silk_Resampler_1_3_COEFS_LQ[ 2 + 3 ];

/* Tables with coefficients for 4th order ARMA filter */
extern const SKP_int16 SKP_Silk_Resampler_320_441_ARMA4_COEFS[ 7 ];
extern const SKP_int16 SKP_Silk_Resampler_240_441_ARMA4_COEFS[ 7 ];
extern const SKP_int16 SKP_Silk_Resampler_160_441_ARMA4_COEFS[ 7 ];
extern const SKP_int16 SKP_Silk_Resampler_120_441_ARMA4_COEFS[ 7 ];
extern const SKP_int16 SKP_Silk_Resampler_80_441_ARMA4_COEFS[ 7 ];

/* Table with interplation fractions of 1/288 : 2/288 : 287/288 (432 Words) */
extern const SKP_int16 SKP_Silk_resampler_frac_FIR_144[ 144 ][ RESAMPLER_ORDER_FIR_144 / 2 ];

#ifdef  __cplusplus
}
#endif

#endif // _SKP_SILK_FIX_RESAMPLER_ROM_H_
