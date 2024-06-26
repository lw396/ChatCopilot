 

#include "SKP_Silk_main.h"

/* NLSF vector decoder */
void SKP_Silk_NLSF_MSVQ_decode(
    SKP_int                         *pNLSF_Q15,     /* O    Pointer to decoded output vector [LPC_ORDER x 1]    */
    const SKP_Silk_NLSF_CB_struct   *psNLSF_CB,     /* I    Pointer to NLSF codebook struct                     */
    const SKP_int                   *NLSFIndices,   /* I    Pointer to NLSF indices          [nStages x 1]      */
    const SKP_int                   LPC_order       /* I    LPC order used                                      */
) 
{
    const SKP_int16 *pCB_element;
          SKP_int    s;
          SKP_int    i;

    /* Check that each index is within valid range */
    SKP_assert( 0 <= NLSFIndices[ 0 ] && NLSFIndices[ 0 ] < psNLSF_CB->CBStages[ 0 ].nVectors );

    /* Point to the first vector element */
    pCB_element = &psNLSF_CB->CBStages[ 0 ].CB_NLSF_Q15[ SKP_MUL( NLSFIndices[ 0 ], LPC_order ) ];

    /* Initialize with the codebook vector from stage 0 */
    for( i = 0; i < LPC_order; i++ ) {
        pNLSF_Q15[ i ] = ( SKP_int )pCB_element[ i ];
    }
          
    for( s = 1; s < psNLSF_CB->nStages; s++ ) {
        /* Check that each index is within valid range */
        SKP_assert( 0 <= NLSFIndices[ s ] && NLSFIndices[ s ] < psNLSF_CB->CBStages[ s ].nVectors );

        if( LPC_order == 16 ) {
            /* Point to the first vector element */
            pCB_element = &psNLSF_CB->CBStages[ s ].CB_NLSF_Q15[ SKP_LSHIFT( NLSFIndices[ s ], 4 ) ];

            /* Add the codebook vector from the current stage */
            pNLSF_Q15[  0 ] += pCB_element[  0 ];
            pNLSF_Q15[  1 ] += pCB_element[  1 ];
            pNLSF_Q15[  2 ] += pCB_element[  2 ];
            pNLSF_Q15[  3 ] += pCB_element[  3 ];
            pNLSF_Q15[  4 ] += pCB_element[  4 ];
            pNLSF_Q15[  5 ] += pCB_element[  5 ];
            pNLSF_Q15[  6 ] += pCB_element[  6 ];
            pNLSF_Q15[  7 ] += pCB_element[  7 ];
            pNLSF_Q15[  8 ] += pCB_element[  8 ];
            pNLSF_Q15[  9 ] += pCB_element[  9 ];
            pNLSF_Q15[ 10 ] += pCB_element[ 10 ];
            pNLSF_Q15[ 11 ] += pCB_element[ 11 ];
            pNLSF_Q15[ 12 ] += pCB_element[ 12 ];
            pNLSF_Q15[ 13 ] += pCB_element[ 13 ];
            pNLSF_Q15[ 14 ] += pCB_element[ 14 ];
            pNLSF_Q15[ 15 ] += pCB_element[ 15 ];
        } else {
            /* Point to the first vector element */
            pCB_element = &psNLSF_CB->CBStages[ s ].CB_NLSF_Q15[ SKP_SMULBB( NLSFIndices[ s ], LPC_order ) ];

            /* Add the codebook vector from the current stage */
            for( i = 0; i < LPC_order; i++ ) {
                pNLSF_Q15[ i ] += pCB_element[ i ];
            }
        }
    }

    /* NLSF stabilization */
    SKP_Silk_NLSF_stabilize( pNLSF_Q15, psNLSF_CB->NDeltaMin_Q15, LPC_order );
}
