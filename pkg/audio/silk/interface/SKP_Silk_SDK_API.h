#ifndef SKP_SILK_SDK_API_H
#define SKP_SILK_SDK_API_H

#include "SKP_Silk_control.h"
#include "SKP_Silk_typedef.h"
#include "SKP_Silk_errors.h"

#ifdef __cplusplus
extern "C"
{
#endif

#define SILK_MAX_FRAMES_PER_PACKET  5

/* Struct for TOC (Table of Contents) */
typedef struct {
    SKP_int     framesInPacket;                             /* Number of 20 ms frames in packet     */
    SKP_int     fs_kHz;                                     /* Sampling frequency in packet         */
    SKP_int     inbandLBRR;                                 /* Does packet contain LBRR information */
    SKP_int     corrupt;                                    /* Packet is corrupt                    */
    SKP_int     vadFlags[     SILK_MAX_FRAMES_PER_PACKET ]; /* VAD flag for each frame in packet    */
    SKP_int     sigtypeFlags[ SILK_MAX_FRAMES_PER_PACKET ]; /* Signal type for each frame in packet */
} SKP_Silk_TOC_struct;

/****************************************/
/* Encoder functions                    */
/****************************************/

/***********************************************/
/* Get size in bytes of the Silk encoder state */
/***********************************************/
SKP_int SKP_Silk_SDK_Get_Encoder_Size( 
    SKP_int32                           *encSizeBytes   /* O:   Number of bytes in SILK encoder state           */
);

/*************************/
/* Init or reset encoder */
/*************************/
SKP_int SKP_Silk_SDK_InitEncoder(
    void                                *encState,      /* I/O: State                                           */
    SKP_SILK_SDK_EncControlStruct       *encStatus      /* O:   Encoder Status                                  */
);

/***************************************/
/* Read control structure from encoder */
/***************************************/
SKP_int SKP_Silk_SDK_QueryEncoder(
    const void                          *encState,      /* I:   State                                           */
    SKP_SILK_SDK_EncControlStruct       *encStatus      /* O:   Encoder Status                                  */
);

/**************************/
/* Encode frame with Silk */
/**************************/
SKP_int SKP_Silk_SDK_Encode( 
    void                                *encState,      /* I/O: State                                           */
    const SKP_SILK_SDK_EncControlStruct *encControl,    /* I:   Control status                                  */
    const SKP_int16                     *samplesIn,     /* I:   Speech sample input vector                      */
    SKP_int                             nSamplesIn,     /* I:   Number of samples in input vector               */
    SKP_uint8                           *outData,       /* O:   Encoded output vector                           */
    SKP_int16                           *nBytesOut      /* I/O: Number of bytes in outData (input: Max bytes)   */
);

/****************************************/
/* Decoder functions                    */
/****************************************/

/***********************************************/
/* Get size in bytes of the Silk decoder state */
/***********************************************/
SKP_int SKP_Silk_SDK_Get_Decoder_Size( 
    SKP_int32                           *decSizeBytes   /* O:   Number of bytes in SILK decoder state           */
);

/*************************/
/* Init or Reset decoder */
/*************************/
SKP_int SKP_Silk_SDK_InitDecoder( 
    void                                *decState       /* I/O: State                                           */
);

/******************/
/* Decode a frame */
/******************/
SKP_int SKP_Silk_SDK_Decode(
    void*                               decState,       /* I/O: State                                           */
    SKP_SILK_SDK_DecControlStruct*      decControl,     /* I/O: Control Structure                               */
    SKP_int                             lostFlag,       /* I:   0: no loss, 1 loss                              */
    const SKP_uint8                     *inData,        /* I:   Encoded input vector                            */
    const SKP_int                       nBytesIn,       /* I:   Number of input bytes                           */
    SKP_int16                           *samplesOut,    /* O:   Decoded output speech vector                    */
    SKP_int16                           *nSamplesOut    /* I/O: Number of samples (vector/decoded)              */
);

/***************************************************************/
/* Find Low Bit Rate Redundancy (LBRR) information in a packet */
/***************************************************************/
void SKP_Silk_SDK_search_for_LBRR(
    const SKP_uint8                     *inData,        /* I:   Encoded input vector                            */
    const SKP_int                       nBytesIn,       /* I:   Number of input Bytes                           */
    SKP_int                             lost_offset,    /* I:   Offset from lost packet                         */
    SKP_uint8                           *LBRRData,      /* O:   LBRR payload                                    */
    SKP_int16                           *nLBRRBytes     /* O:   Number of LBRR Bytes                            */
);

/**************************************/
/* Get table of contents for a packet */
/**************************************/
void SKP_Silk_SDK_get_TOC(
    const SKP_uint8                     *inData,        /* I:   Encoded input vector                            */
    const SKP_int                       nBytesIn,       /* I:   Number of input bytes                           */
    SKP_Silk_TOC_struct                 *Silk_TOC       /* O:   Table of contents                               */
);

/**************************/
/* Get the version number */
/**************************/
/* Return a pointer to string specifying the version */ 
const char *SKP_Silk_SDK_get_version(void);

#ifdef __cplusplus
}
#endif

#endif
