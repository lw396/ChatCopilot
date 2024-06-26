#ifndef SKP_SILK_ERRORS_H
#define SKP_SILK_ERRORS_H

#ifdef __cplusplus
extern "C"
{
#endif

/******************/
/* Error messages */
/******************/
#define SKP_SILK_NO_ERROR                               0

/**************************/
/* Encoder error messages */
/**************************/

/* Input length is not a multiplum of 10 ms, or length is longer than the packet length */
#define SKP_SILK_ENC_INPUT_INVALID_NO_OF_SAMPLES        -1

/* Sampling frequency not 8000, 12000, 16000 or 24000 Hertz */
#define SKP_SILK_ENC_FS_NOT_SUPPORTED                   -2

/* Packet size not 20, 40, 60, 80 or 100 ms */
#define SKP_SILK_ENC_PACKET_SIZE_NOT_SUPPORTED          -3

/* Allocated payload buffer too short */
#define SKP_SILK_ENC_PAYLOAD_BUF_TOO_SHORT              -4

/* Loss rate not between 0 and 100 percent */
#define SKP_SILK_ENC_INVALID_LOSS_RATE                  -5

/* Complexity setting not valid, use 0, 1 or 2 */
#define SKP_SILK_ENC_INVALID_COMPLEXITY_SETTING         -6

/* Inband FEC setting not valid, use 0 or 1 */
#define SKP_SILK_ENC_INVALID_INBAND_FEC_SETTING         -7

/* DTX setting not valid, use 0 or 1 */
#define SKP_SILK_ENC_INVALID_DTX_SETTING                -8

/* Internal encoder error */
#define SKP_SILK_ENC_INTERNAL_ERROR                     -9

/**************************/
/* Decoder error messages */
/**************************/

/* Output sampling frequency lower than internal decoded sampling frequency */
#define SKP_SILK_DEC_INVALID_SAMPLING_FREQUENCY         -10

/* Payload size exceeded the maximum allowed 1024 bytes */
#define SKP_SILK_DEC_PAYLOAD_TOO_LARGE                  -11

/* Payload has bit errors */
#define SKP_SILK_DEC_PAYLOAD_ERROR                      -12

#ifdef __cplusplus
}
#endif

#endif
