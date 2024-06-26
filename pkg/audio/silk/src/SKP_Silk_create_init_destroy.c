#include "SKP_Silk_main.h"

/************************/
/* Init Decoder State   */
/************************/
SKP_int SKP_Silk_init_decoder(
    SKP_Silk_decoder_state      *psDec              /* I/O  Decoder state pointer                       */
)
{
    SKP_memset( psDec, 0, sizeof( SKP_Silk_decoder_state ) );
    /* Set sampling rate to 24 kHz, and init non-zero values */
    SKP_Silk_decoder_set_fs( psDec, 24 );

    /* Used to deactivate e.g. LSF interpolation and fluctuation reduction */
    psDec->first_frame_after_reset = 1;
    psDec->prev_inv_gain_Q16 = 65536;

    /* Reset CNG state */
    SKP_Silk_CNG_Reset( psDec );

    SKP_Silk_PLC_Reset( psDec );
    
    return(0);
}

