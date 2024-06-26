 

/***********************************************************
* Pitch analyser function
********************************************************** */
#include "SKP_Silk_SigProc_FIX.h"
#include "SKP_Silk_common_pitch_est_defines.h"

void SKP_Silk_decode_pitch(
    SKP_int          lagIndex,                        /* I                             */
    SKP_int          contourIndex,                    /* O                             */
    SKP_int          pitch_lags[],                    /* O 4 pitch values              */
    SKP_int          Fs_kHz                           /* I sampling frequency (kHz)    */
)
{
    SKP_int lag, i, min_lag;

    min_lag = SKP_SMULBB( PITCH_EST_MIN_LAG_MS, Fs_kHz );

    /* Only for 24 / 16 kHz version for now */
    lag = min_lag + lagIndex;
    if( Fs_kHz == 8 ) {
        /* Only a small codebook for 8 khz */
        for( i = 0; i < PITCH_EST_NB_SUBFR; i++ ) {
            pitch_lags[ i ] = lag + SKP_Silk_CB_lags_stage2[ i ][ contourIndex ];
        }
    } else {
        for( i = 0; i < PITCH_EST_NB_SUBFR; i++ ) {
            pitch_lags[ i ] = lag + SKP_Silk_CB_lags_stage3[ i ][ contourIndex ];
        }
    }
}
