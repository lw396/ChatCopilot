 

#ifndef SKP_SILK_TABLES_H
#define SKP_SILK_TABLES_H

#include "SKP_Silk_define.h"
#include "SKP_Silk_structs.h"

#define PITCH_EST_MAX_LAG_MS                18          /* 18 ms -> 56 Hz */
#define PITCH_EST_MIN_LAG_MS                2           /* 2 ms -> 500 Hz */

#ifdef __cplusplus
extern "C"
{
#endif

/* entropy coding tables */
extern const SKP_uint16 SKP_Silk_type_offset_CDF[ 5 ];                                              /*   5 */
extern const SKP_uint16 SKP_Silk_type_offset_joint_CDF[ 4 ][ 5 ];                                   /*  20 */
extern const SKP_int    SKP_Silk_type_offset_CDF_offset;

extern const SKP_uint16 SKP_Silk_gain_CDF[ 2 ][ N_LEVELS_QGAIN + 1 ];                               /* 130 */
extern const SKP_int    SKP_Silk_gain_CDF_offset;
extern const SKP_uint16 SKP_Silk_delta_gain_CDF[ MAX_DELTA_GAIN_QUANT - MIN_DELTA_GAIN_QUANT + 2 ]; /*  46 */
extern const SKP_int    SKP_Silk_delta_gain_CDF_offset;

extern const SKP_uint16 SKP_Silk_pitch_lag_NB_CDF[ 8 * ( PITCH_EST_MAX_LAG_MS - PITCH_EST_MIN_LAG_MS ) + 2 ];   /* 130 */
extern const SKP_int    SKP_Silk_pitch_lag_NB_CDF_offset;
extern const SKP_uint16 SKP_Silk_pitch_lag_MB_CDF[ 12 * ( PITCH_EST_MAX_LAG_MS - PITCH_EST_MIN_LAG_MS ) + 2 ];  /* 194 */
extern const SKP_int    SKP_Silk_pitch_lag_MB_CDF_offset;
extern const SKP_uint16 SKP_Silk_pitch_lag_WB_CDF[ 16 * ( PITCH_EST_MAX_LAG_MS - PITCH_EST_MIN_LAG_MS ) + 2 ];  /* 258 */
extern const SKP_int    SKP_Silk_pitch_lag_WB_CDF_offset;
extern const SKP_uint16 SKP_Silk_pitch_lag_SWB_CDF[ 24 * ( PITCH_EST_MAX_LAG_MS - PITCH_EST_MIN_LAG_MS ) + 2 ]; /* 386 */
extern const SKP_int    SKP_Silk_pitch_lag_SWB_CDF_offset;

extern const SKP_uint16 SKP_Silk_pitch_contour_CDF[ 35 ];                                           /*  35 */
extern const SKP_int    SKP_Silk_pitch_contour_CDF_offset;
extern const SKP_uint16 SKP_Silk_pitch_contour_NB_CDF[ 12 ];                                        /*  12 */
extern const SKP_int    SKP_Silk_pitch_contour_NB_CDF_offset;
extern const SKP_uint16 SKP_Silk_pitch_delta_CDF[23];                                               /* 23 */
extern const SKP_int    SKP_Silk_pitch_delta_CDF_offset;

extern const SKP_uint16 SKP_Silk_pulses_per_block_CDF[ N_RATE_LEVELS ][ MAX_PULSES + 3 ];           /* 210 */
extern const SKP_int    SKP_Silk_pulses_per_block_CDF_offset;
extern const SKP_int16  SKP_Silk_pulses_per_block_BITS_Q6[ N_RATE_LEVELS - 1 ][ MAX_PULSES + 2 ];   /* 180 */

extern const SKP_uint16 SKP_Silk_rate_levels_CDF[ 2 ][ N_RATE_LEVELS ];                             /*  20 */
extern const SKP_int    SKP_Silk_rate_levels_CDF_offset;
extern const SKP_int16  SKP_Silk_rate_levels_BITS_Q6[ 2 ][ N_RATE_LEVELS - 1 ];                     /*  18 */

extern const SKP_int    SKP_Silk_max_pulses_table[ 4 ];                                             /*   4 */

extern const SKP_uint16 SKP_Silk_shell_code_table0[  33 ];                                          /*  33 */
extern const SKP_uint16 SKP_Silk_shell_code_table1[  52 ];                                          /*  52 */
extern const SKP_uint16 SKP_Silk_shell_code_table2[ 102 ];                                          /* 102 */
extern const SKP_uint16 SKP_Silk_shell_code_table3[ 207 ];                                          /* 207 */
extern const SKP_uint16 SKP_Silk_shell_code_table_offsets[ 19 ];                                    /*  19 */

extern const SKP_uint16 SKP_Silk_lsb_CDF[ 3 ];                                                      /*   3 */

extern const SKP_uint16 SKP_Silk_sign_CDF[ 36 ];                                                    /*  36 */

extern const SKP_uint16 SKP_Silk_LTP_per_index_CDF[ 4 ];                                            /*   4 */
extern const SKP_int    SKP_Silk_LTP_per_index_CDF_offset;
extern const SKP_int16  * const SKP_Silk_LTP_gain_BITS_Q6_ptrs[ NB_LTP_CBKS ];                      /*   3 */
extern const SKP_uint16 * const SKP_Silk_LTP_gain_CDF_ptrs[ NB_LTP_CBKS ];                          /*   3 */
extern const SKP_int    SKP_Silk_LTP_gain_CDF_offsets[ NB_LTP_CBKS ];                               /*   3 */
extern const SKP_int32  SKP_Silk_LTP_gain_middle_avg_RD_Q14;
extern const SKP_uint16 SKP_Silk_LTPscale_CDF[ 4 ];                                                 /*   4 */
extern const SKP_int    SKP_Silk_LTPscale_offset;

/* Tables for LTPScale */
extern const SKP_int16  SKP_Silk_LTPScales_table_Q14[ 3 ];

extern const SKP_uint16 SKP_Silk_vadflag_CDF[ 3 ];                                                  /*   3 */
extern const SKP_int    SKP_Silk_vadflag_offset;

extern const SKP_int    SKP_Silk_SamplingRates_table[ 4 ];                                          /*   4 */
extern const SKP_uint16 SKP_Silk_SamplingRates_CDF[ 5 ];                                            /*   5 */
extern const SKP_int    SKP_Silk_SamplingRates_offset;

extern const SKP_uint16 SKP_Silk_NLSF_interpolation_factor_CDF[ 6 ];
extern const SKP_int    SKP_Silk_NLSF_interpolation_factor_offset;

/* NLSF codebooks */
extern const SKP_Silk_NLSF_CB_struct SKP_Silk_NLSF_CB0_16, SKP_Silk_NLSF_CB1_16;
extern const SKP_Silk_NLSF_CB_struct SKP_Silk_NLSF_CB0_10, SKP_Silk_NLSF_CB1_10;

/* quantization tables */
extern const SKP_int16 * const SKP_Silk_LTP_vq_ptrs_Q14[ NB_LTP_CBKS ];                             /* 168 */
extern const SKP_int    SKP_Silk_LTP_vq_sizes[ NB_LTP_CBKS ];                                       /*   3 */

/* Piece-wise linear mapping from bitrate in kbps to coding quality in dB SNR */
extern const SKP_int32  TargetRate_table_NB[  TARGET_RATE_TAB_SZ ];
extern const SKP_int32  TargetRate_table_MB[  TARGET_RATE_TAB_SZ ];
extern const SKP_int32  TargetRate_table_WB[  TARGET_RATE_TAB_SZ ];
extern const SKP_int32  TargetRate_table_SWB[ TARGET_RATE_TAB_SZ ];
extern const SKP_int32  SNR_table_Q1[         TARGET_RATE_TAB_SZ ];

extern const SKP_int32  SNR_table_one_bit_per_sample_Q7[ 4 ];

/* Filter coeficicnts for HP filter: 4. Order filter implementad as two biquad filters  */
extern const SKP_int16  SKP_Silk_SWB_detect_B_HP_Q13[ NB_SOS ][ 3 ];
extern const SKP_int16  SKP_Silk_SWB_detect_A_HP_Q13[ NB_SOS ][ 2 ];

/* Decoder high-pass filter coefficients for 24 kHz sampling */
extern const SKP_int16  SKP_Silk_Dec_A_HP_24[ DEC_HP_ORDER ];                                       /*   2 */
extern const SKP_int16  SKP_Silk_Dec_B_HP_24[ DEC_HP_ORDER + 1 ];                                   /*   3 */

/* Decoder high-pass filter coefficients for 16 kHz sampling */
extern const SKP_int16  SKP_Silk_Dec_A_HP_16[ DEC_HP_ORDER ];                                       /*   2 */
extern const SKP_int16  SKP_Silk_Dec_B_HP_16[ DEC_HP_ORDER + 1 ];                                   /*   3 */

/* Decoder high-pass filter coefficients for 12 kHz sampling */
extern const SKP_int16  SKP_Silk_Dec_A_HP_12[ DEC_HP_ORDER ];                                       /*   2 */
extern const SKP_int16  SKP_Silk_Dec_B_HP_12[ DEC_HP_ORDER + 1 ];                                   /*   3 */

/* Decoder high-pass filter coefficients for 8 kHz sampling */
extern const SKP_int16  SKP_Silk_Dec_A_HP_8[ DEC_HP_ORDER ];                                        /*   2 */
extern const SKP_int16  SKP_Silk_Dec_B_HP_8[ DEC_HP_ORDER + 1 ];                                    /*   3 */

/* Table for frame termination indication */
extern const SKP_uint16 SKP_Silk_FrameTermination_CDF[ 5 ];
extern const SKP_int    SKP_Silk_FrameTermination_offset;

/* Table for random seed */
extern const SKP_uint16 SKP_Silk_Seed_CDF[ 5 ];
extern const SKP_int    SKP_Silk_Seed_offset;

/* Quantization offsets */
extern const SKP_int16  SKP_Silk_Quantization_Offsets_Q10[ 2 ][ 2 ];

#if SWITCH_TRANSITION_FILTERING
/* Interpolation points for filter coefficients used in the bandwidth transition smoother */
extern const SKP_int32 SKP_Silk_Transition_LP_B_Q28[ TRANSITION_INT_NUM ][ TRANSITION_NB ];
extern const SKP_int32 SKP_Silk_Transition_LP_A_Q28[ TRANSITION_INT_NUM ][ TRANSITION_NA ];
#endif

#ifdef __cplusplus
}
#endif

#endif
