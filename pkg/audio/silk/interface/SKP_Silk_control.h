#ifndef SKP_SILK_CONTROL_H
#define SKP_SILK_CONTROL_H

#include "SKP_Silk_typedef.h"

#ifdef __cplusplus
extern "C"
{
#endif

/***********************************************/
/* Structure for controlling encoder operation */
/***********************************************/
typedef struct {
    /* I:   Input signal sampling rate in Hertz; 8000/12000/16000/24000                     */
    SKP_int32 API_sampleRate;

    /* I:   Maximum internal sampling rate in Hertz; 8000/12000/16000/24000                 */
    SKP_int32 maxInternalSampleRate;

    /* I:   Number of samples per packet; must be equivalent of 20, 40, 60, 80 or 100 ms    */
    SKP_int packetSize;

    /* I:   Bitrate during active speech in bits/second; internally limited                 */
    SKP_int32 bitRate;                        

    /* I:   Uplink packet loss in percent (0-100)                                           */
    SKP_int packetLossPercentage;
    
    /* I:   Complexity mode; 0 is lowest; 1 is medium and 2 is highest complexity           */
    SKP_int complexity;

    /* I:   Flag to enable in-band Forward Error Correction (FEC); 0/1                      */
    SKP_int useInBandFEC;

    /* I:   Flag to enable discontinuous transmission (DTX); 0/1                            */
    SKP_int useDTX;
} SKP_SILK_SDK_EncControlStruct;

/**************************************************************************/
/* Structure for controlling decoder operation and reading decoder status */
/**************************************************************************/
typedef struct {
    /* I:   Output signal sampling rate in Hertz; 8000/12000/16000/24000                    */
    SKP_int32 API_sampleRate;

    /* O:   Number of samples per frame                                                     */
    SKP_int frameSize;

    /* O:   Frames per packet 1, 2, 3, 4, 5                                                 */
    SKP_int framesPerPacket;

    /* O:   Flag to indicate that the decoder has remaining payloads internally             */
    SKP_int moreInternalDecoderFrames;

    /* O:   Distance between main payload and redundant payload in packets                  */
    SKP_int inBandFECOffset;
} SKP_SILK_SDK_DecControlStruct;

#ifdef __cplusplus
}
#endif

#endif
