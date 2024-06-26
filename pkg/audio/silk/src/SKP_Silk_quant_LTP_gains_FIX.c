 

#include "SKP_Silk_main_FIX.h"

void SKP_Silk_quant_LTP_gains_FIX(
    SKP_int16               B_Q14[],                /* I/O  (un)quantized LTP gains     */
    SKP_int                 cbk_index[],            /* O    Codebook Index              */
    SKP_int                 *periodicity_index,     /* O    Periodicity Index           */
    const SKP_int32         W_Q18[],                /* I    Error Weights in Q18        */
    SKP_int                 mu_Q8,                  /* I    Mu value (R/D tradeoff)     */
    SKP_int                 lowComplexity           /* I    Flag for low complexity     */
)
{
    SKP_int             j, k, temp_idx[ NB_SUBFR ], cbk_size;
    const SKP_int16     *cl_ptr;
    const SKP_int16     *cbk_ptr_Q14;
    const SKP_int16     *b_Q14_ptr;
    const SKP_int32     *W_Q18_ptr;
    SKP_int32           rate_dist_subfr, rate_dist, min_rate_dist;



    /***************************************************/
    /* iterate over different codebooks with different */
    /* rates/distortions, and choose best */
    /***************************************************/
    min_rate_dist = SKP_int32_MAX;
    for( k = 0; k < 3; k++ ) {
        cl_ptr      = SKP_Silk_LTP_gain_BITS_Q6_ptrs[ k ];
        cbk_ptr_Q14 = SKP_Silk_LTP_vq_ptrs_Q14[       k ];
        cbk_size    = SKP_Silk_LTP_vq_sizes[          k ];

        /* Setup pointer to first subframe */
        W_Q18_ptr = W_Q18;
        b_Q14_ptr = B_Q14;

        rate_dist = 0;
        for( j = 0; j < NB_SUBFR; j++ ) {

            SKP_Silk_VQ_WMat_EC_FIX(
                &temp_idx[ j ],         /* O    index of best codebook vector                           */
                &rate_dist_subfr,       /* O    best weighted quantization error + mu * rate            */
                b_Q14_ptr,              /* I    input vector to be quantized                            */
                W_Q18_ptr,              /* I    weighting matrix                                        */
                cbk_ptr_Q14,            /* I    codebook                                                */
                cl_ptr,                 /* I    code length for each codebook vector                    */
                mu_Q8,                  /* I    tradeoff between weighted error and rate                */
                cbk_size                /* I    number of vectors in codebook                           */
            );

            rate_dist = SKP_ADD_POS_SAT32( rate_dist, rate_dist_subfr );

            b_Q14_ptr += LTP_ORDER;
            W_Q18_ptr += LTP_ORDER * LTP_ORDER;
        }

        /* Avoid never finding a codebook */
        rate_dist = SKP_min( SKP_int32_MAX - 1, rate_dist );

        if( rate_dist < min_rate_dist ) {
            min_rate_dist = rate_dist;
            SKP_memcpy( cbk_index, temp_idx, NB_SUBFR * sizeof( SKP_int ) );
            *periodicity_index = k;
        }

        /* Break early in low-complexity mode if rate distortion is below threshold */
        if( lowComplexity && ( rate_dist < SKP_Silk_LTP_gain_middle_avg_RD_Q14 ) ) {
            break;
        }
    }

    cbk_ptr_Q14 = SKP_Silk_LTP_vq_ptrs_Q14[ *periodicity_index ];
    for( j = 0; j < NB_SUBFR; j++ ) {
        for( k = 0; k < LTP_ORDER; k++ ) { 
            B_Q14[ j * LTP_ORDER + k ] = cbk_ptr_Q14[ SKP_MLA( k, cbk_index[ j ], LTP_ORDER ) ];
        }
    }
}

