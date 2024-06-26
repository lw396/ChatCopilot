 

#ifndef SIGPROC_COMMON_PITCH_EST_DEFINES_H
#define SIGPROC_COMMON_PITCH_EST_DEFINES_H

#include "SKP_Silk_SigProc_FIX.h"

/************************************************************/
/* Definitions For Fix pitch estimator                      */
/************************************************************/

#define PITCH_EST_MAX_FS_KHZ                24 /* Maximum sampling frequency used */

#define PITCH_EST_FRAME_LENGTH_MS           40 /* 40 ms */

#define PITCH_EST_MAX_FRAME_LENGTH          (PITCH_EST_FRAME_LENGTH_MS * PITCH_EST_MAX_FS_KHZ)
#define PITCH_EST_MAX_FRAME_LENGTH_ST_1     (PITCH_EST_MAX_FRAME_LENGTH >> 2)
#define PITCH_EST_MAX_FRAME_LENGTH_ST_2     (PITCH_EST_MAX_FRAME_LENGTH >> 1)
#define PITCH_EST_MAX_SF_FRAME_LENGTH       (PITCH_EST_SUB_FRAME * PITCH_EST_MAX_FS_KHZ)

#define PITCH_EST_MAX_LAG_MS                18            /* 18 ms -> 56 Hz */
#define PITCH_EST_MIN_LAG_MS                2            /* 2 ms -> 500 Hz */
#define PITCH_EST_MAX_LAG                   (PITCH_EST_MAX_LAG_MS * PITCH_EST_MAX_FS_KHZ)
#define PITCH_EST_MIN_LAG                   (PITCH_EST_MIN_LAG_MS * PITCH_EST_MAX_FS_KHZ)

#define PITCH_EST_NB_SUBFR                  4

#define PITCH_EST_D_SRCH_LENGTH             24

#define PITCH_EST_MAX_DECIMATE_STATE_LENGTH 7

#define PITCH_EST_NB_STAGE3_LAGS            5

#define PITCH_EST_NB_CBKS_STAGE2            3
#define PITCH_EST_NB_CBKS_STAGE2_EXT        11

#define PITCH_EST_CB_mn2                    1
#define PITCH_EST_CB_mx2                    2

#define PITCH_EST_NB_CBKS_STAGE3_MAX        34
#define PITCH_EST_NB_CBKS_STAGE3_MID        24
#define PITCH_EST_NB_CBKS_STAGE3_MIN        16

extern const SKP_int16 SKP_Silk_CB_lags_stage2[PITCH_EST_NB_SUBFR][PITCH_EST_NB_CBKS_STAGE2_EXT];
extern const SKP_int16 SKP_Silk_CB_lags_stage3[PITCH_EST_NB_SUBFR][PITCH_EST_NB_CBKS_STAGE3_MAX];
extern const SKP_int16 SKP_Silk_Lag_range_stage3[ SKP_Silk_PITCH_EST_MAX_COMPLEX + 1 ] [ PITCH_EST_NB_SUBFR ][ 2 ];
extern const SKP_int16 SKP_Silk_cbk_sizes_stage3[ SKP_Silk_PITCH_EST_MAX_COMPLEX + 1 ];
extern const SKP_int16 SKP_Silk_cbk_offsets_stage3[ SKP_Silk_PITCH_EST_MAX_COMPLEX + 1 ];

#endif

