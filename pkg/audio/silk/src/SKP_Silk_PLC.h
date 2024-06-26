 

#ifndef SKP_SILK_PLC_FIX_H
#define SKP_SILK_PLC_FIX_H

#include "SKP_Silk_main.h"

#define BWE_COEF_Q16                    64880           /* 0.99 in Q16                      */
#define V_PITCH_GAIN_START_MIN_Q14      11469           /* 0.7 in Q14                       */
#define V_PITCH_GAIN_START_MAX_Q14      15565           /* 0.95 in Q14                      */
#define MAX_PITCH_LAG_MS                18
#define SA_THRES_Q8                     50
#define USE_SINGLE_TAP                  1
#define RAND_BUF_SIZE                   128
#define RAND_BUF_MASK                   (RAND_BUF_SIZE - 1)
#define LOG2_INV_LPC_GAIN_HIGH_THRES    3               /* 2^3 = 8 dB LPC gain              */
#define LOG2_INV_LPC_GAIN_LOW_THRES     8               /* 2^8 = 24 dB LPC gain             */
#define PITCH_DRIFT_FAC_Q16             655             /* 0.01 in Q16                      */

void SKP_Silk_PLC_Reset(
    SKP_Silk_decoder_state      *psDec              /* I/O Decoder state        */
);

void SKP_Silk_PLC(
    SKP_Silk_decoder_state      *psDec,             /* I/O Decoder state        */
    SKP_Silk_decoder_control    *psDecCtrl,         /* I/O Decoder control      */
    SKP_int16                   signal[],           /* I/O  signal              */
    SKP_int                     length,             /* I length of residual     */
    SKP_int                     lost                /* I Loss flag              */
);

void SKP_Silk_PLC_update(
    SKP_Silk_decoder_state      *psDec,             /* I/O Decoder state        */
    SKP_Silk_decoder_control    *psDecCtrl,         /* I/O Decoder control      */
    SKP_int16                   signal[],
    SKP_int                     length
);

void SKP_Silk_PLC_conceal(
    SKP_Silk_decoder_state      *psDec,             /* I/O Decoder state        */
    SKP_Silk_decoder_control    *psDecCtrl,         /* I/O Decoder control      */
    SKP_int16                   signal[],           /* O LPC residual signal    */
    SKP_int                     length              /* I length of signal       */
);

void SKP_Silk_PLC_glue_frames(
    SKP_Silk_decoder_state      *psDec,             /* I/O decoder state        */
    SKP_Silk_decoder_control    *psDecCtrl,         /* I/O Decoder control      */
    SKP_int16                   signal[],           /* I/O signal               */
    SKP_int                     length              /* I length of signal       */
);

#endif

