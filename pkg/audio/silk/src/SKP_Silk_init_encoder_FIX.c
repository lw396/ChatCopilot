 

#include "SKP_Silk_main_FIX.h"

/*********************************/
/* Initialize Silk Encoder state */
/*********************************/
SKP_int SKP_Silk_init_encoder_FIX(
    SKP_Silk_encoder_state_FIX  *psEnc                  /* I/O  Pointer to Silk FIX encoder state       */
) {
    SKP_int ret = 0;
    /* Clear the entire encoder state */
    SKP_memset( psEnc, 0, sizeof( SKP_Silk_encoder_state_FIX ) );

#if HIGH_PASS_INPUT
    psEnc->variable_HP_smth1_Q15 = 200844; /* = SKP_Silk_log2(70)_Q0; */
    psEnc->variable_HP_smth2_Q15 = 200844; /* = SKP_Silk_log2(70)_Q0; */
#endif

    /* Used to deactivate e.g. LSF interpolation and fluctuation reduction */
    psEnc->sCmn.first_frame_after_reset = 1;

    /* Initialize Silk VAD */
    ret += SKP_Silk_VAD_Init( &psEnc->sCmn.sVAD );

    /* Initialize NSQ */
    psEnc->sCmn.sNSQ.prev_inv_gain_Q16      = 65536;
    psEnc->sCmn.sNSQ_LBRR.prev_inv_gain_Q16 = 65536;

    return( ret );
}
