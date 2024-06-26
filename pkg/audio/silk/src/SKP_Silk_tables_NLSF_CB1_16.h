 

#ifndef SKP_SILK_TABLES_NLSF_CB1_16_H
#define SKP_SILK_TABLES_NLSF_CB1_16_H

#include "SKP_Silk_define.h"

#ifdef __cplusplus
extern "C"
{
#endif

#define NLSF_MSVQ_CB1_16_STAGES       10
#define NLSF_MSVQ_CB1_16_VECTORS      104

/* NLSF codebook entropy coding tables */
extern const SKP_uint16         SKP_Silk_NLSF_MSVQ_CB1_16_CDF[ NLSF_MSVQ_CB1_16_VECTORS + NLSF_MSVQ_CB1_16_STAGES ];
extern const SKP_uint16 * const SKP_Silk_NLSF_MSVQ_CB1_16_CDF_start_ptr[                  NLSF_MSVQ_CB1_16_STAGES ];
extern const SKP_int            SKP_Silk_NLSF_MSVQ_CB1_16_CDF_middle_idx[                 NLSF_MSVQ_CB1_16_STAGES ];

#ifdef __cplusplus
}
#endif

#endif

