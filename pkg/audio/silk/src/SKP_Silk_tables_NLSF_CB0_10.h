 

#ifndef SKP_SILK_TABLES_NLSF_CB0_10_H
#define SKP_SILK_TABLES_NLSF_CB0_10_H

#include "SKP_Silk_define.h"

#ifdef __cplusplus
extern "C"
{
#endif

#define NLSF_MSVQ_CB0_10_STAGES       6
#define NLSF_MSVQ_CB0_10_VECTORS      120

/* NLSF codebook entropy coding tables */
extern const SKP_uint16         SKP_Silk_NLSF_MSVQ_CB0_10_CDF[ NLSF_MSVQ_CB0_10_VECTORS + NLSF_MSVQ_CB0_10_STAGES ];
extern const SKP_uint16 * const SKP_Silk_NLSF_MSVQ_CB0_10_CDF_start_ptr[                  NLSF_MSVQ_CB0_10_STAGES ];
extern const SKP_int            SKP_Silk_NLSF_MSVQ_CB0_10_CDF_middle_idx[                 NLSF_MSVQ_CB0_10_STAGES ];

#ifdef __cplusplus
}
#endif

#endif

