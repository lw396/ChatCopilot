 

#include "SKP_Silk_main.h"

/* Resets LBRR buffer, used if packet size changes */
void SKP_Silk_LBRR_reset( 
    SKP_Silk_encoder_state      *psEncC             /* I/O  state                                       */
)
{
    SKP_int i;

    for( i = 0; i < MAX_LBRR_DELAY; i++ ) {
        psEncC->LBRR_buffer[ i ].usage = SKP_SILK_NO_LBRR;
    }
}
