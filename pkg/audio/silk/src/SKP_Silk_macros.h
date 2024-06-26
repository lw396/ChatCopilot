 

#ifndef _SKP_SILK_API_C_H_
#define _SKP_SILK_API_C_H_

// This is an inline header file for general platform.

// (a32 * (SKP_int32)((SKP_int16)(b32))) >> 16 output have to be 32bit int
#define SKP_SMULWB(a32, b32)            ((((a32) >> 16) * (SKP_int32)((SKP_int16)(b32))) + ((((a32) & 0x0000FFFF) * (SKP_int32)((SKP_int16)(b32))) >> 16))

// a32 + (b32 * (SKP_int32)((SKP_int16)(c32))) >> 16 output have to be 32bit int
#define SKP_SMLAWB(a32, b32, c32)       ((a32) + ((((b32) >> 16) * (SKP_int32)((SKP_int16)(c32))) + ((((b32) & 0x0000FFFF) * (SKP_int32)((SKP_int16)(c32))) >> 16)))

// (a32 * (b32 >> 16)) >> 16
#define SKP_SMULWT(a32, b32)            (((a32) >> 16) * ((b32) >> 16) + ((((a32) & 0x0000FFFF) * ((b32) >> 16)) >> 16))

// a32 + (b32 * (c32 >> 16)) >> 16
#define SKP_SMLAWT(a32, b32, c32)       ((a32) + (((b32) >> 16) * ((c32) >> 16)) + ((((b32) & 0x0000FFFF) * ((c32) >> 16)) >> 16))

// (SKP_int32)((SKP_int16)(a3))) * (SKP_int32)((SKP_int16)(b32)) output have to be 32bit int
#define SKP_SMULBB(a32, b32)            ((SKP_int32)((SKP_int16)(a32)) * (SKP_int32)((SKP_int16)(b32)))

// a32 + (SKP_int32)((SKP_int16)(b32)) * (SKP_int32)((SKP_int16)(c32)) output have to be 32bit int
#define SKP_SMLABB(a32, b32, c32)       ((a32) + ((SKP_int32)((SKP_int16)(b32))) * (SKP_int32)((SKP_int16)(c32)))

// (SKP_int32)((SKP_int16)(a32)) * (b32 >> 16)
#define SKP_SMULBT(a32, b32)            ((SKP_int32)((SKP_int16)(a32)) * ((b32) >> 16))

// a32 + (SKP_int32)((SKP_int16)(b32)) * (c32 >> 16)
#define SKP_SMLABT(a32, b32, c32)       ((a32) + ((SKP_int32)((SKP_int16)(b32))) * ((c32) >> 16))

// a64 + (b32 * c32)
#define SKP_SMLAL(a64, b32, c32)        (SKP_ADD64((a64), ((SKP_int64)(b32) * (SKP_int64)(c32))))

// (a32 * b32) >> 16
#define SKP_SMULWW(a32, b32)            SKP_MLA(SKP_SMULWB((a32), (b32)), (a32), SKP_RSHIFT_ROUND((b32), 16))

// a32 + ((b32 * c32) >> 16)
#define SKP_SMLAWW(a32, b32, c32)       SKP_MLA(SKP_SMLAWB((a32), (b32), (c32)), (b32), SKP_RSHIFT_ROUND((c32), 16))

// (SKP_int32)(((SKP_int64)a32 * b32) >> 32)
#define SKP_SMMUL(a32, b32)             (SKP_int32)SKP_RSHIFT64(SKP_SMULL((a32), (b32)), 32)

/* add/subtract with output saturated */
#define SKP_ADD_SAT32(a, b)             ((((a) + (b)) & 0x80000000) == 0 ?                              \
                                        ((((a) & (b)) & 0x80000000) != 0 ? SKP_int32_MIN : (a)+(b)) :   \
                                        ((((a) | (b)) & 0x80000000) == 0 ? SKP_int32_MAX : (a)+(b)) )

#define SKP_SUB_SAT32(a, b)             ((((a)-(b)) & 0x80000000) == 0 ?                                        \
                                        (( (a) & ((b)^0x80000000) & 0x80000000) ? SKP_int32_MIN : (a)-(b)) :    \
                                        ((((a)^0x80000000) & (b)  & 0x80000000) ? SKP_int32_MAX : (a)-(b)) )
    
SKP_INLINE SKP_int32 SKP_Silk_CLZ16(SKP_int16 in16)
{
    SKP_int32 out32 = 0;
    if( in16 == 0 ) {
        return 16;
    }
    /* test nibbles */
    if( in16 & 0xFF00 ) {
        if( in16 & 0xF000 ) {
            in16 >>= 12;
        } else {
            out32 += 4;
            in16 >>= 8;
        }
    } else {
        if( in16 & 0xFFF0 ) {
            out32 += 8;
            in16 >>= 4;
        } else {
            out32 += 12;
        }
    }
    /* test bits and return */
    if( in16 & 0xC ) {
        if( in16 & 0x8 )
            return out32 + 0;
        else
            return out32 + 1;
    } else {
        if( in16 & 0xE )
            return out32 + 2;
        else
            return out32 + 3;
    }
}

SKP_INLINE SKP_int32 SKP_Silk_CLZ32(SKP_int32 in32)
{
    /* test highest 16 bits and convert to SKP_int16 */
    if( in32 & 0xFFFF0000 ) {
        return SKP_Silk_CLZ16((SKP_int16)(in32 >> 16));
    } else {
        return SKP_Silk_CLZ16((SKP_int16)in32) + 16;
    }
}

#endif //_SKP_SILK_API_C_H_

