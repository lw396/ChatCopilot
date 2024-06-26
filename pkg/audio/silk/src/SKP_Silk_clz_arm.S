 

#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"
#if (EMBEDDED_ARM>0) && (EMBEDDED_ARM<5)

.globl SYM(SKP_Silk_CLZ16)
.globl SYM(SKP_Silk_CLZ32)

SYM(SKP_Silk_CLZ16):
	str		lr, [sp, #-4]!
	
	mov		r2, r0
	mov		r0, #0
	ands	r1, r2, #0xF000
	movne	r1, r2, asr #12 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF00
	movne	r1, r2, asr #8
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF0
	movne	r1, r2, asr #4
	bne		0f
	add		r0, r0, #4
	mov		r1, r2
0:
	ands	r2, r1, #0x8
	bne		1f
	add		r0, r0, #1
	ands	r2, r1, #0x4
	bne		1f
	add		r0, r0, #1
	ands	r2, r1, #0x2
	bne		1f
	add		r0, r0, #1	
	ands	r2, r1, #0x1
	bne		1f
	add		r0, r0, #1
1:	
	ldr		pc, [sp], #4 

SYM(SKP_Silk_CLZ32):

	str		lr, [sp, #-4]!
	
	mov		r2, r0
	mov		r0, #0
	ands	r1, r2, #0xF0000000
	movne	r1, r2, asr #28 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF000000
	movne	r1, r2, asr #24 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF00000
	movne	r1, r2, asr #20 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF0000
	movne	r1, r2, asr #16 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF000
	movne	r1, r2, asr #12 
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF00
	movne	r1, r2, asr #8
	bne		0f
	add		r0, r0, #4
	ands	r1, r2, #0xF0
	movne	r1, r2, asr #4
	bne		0f
	add		r0, r0, #4
	mov		r1, r2
0:
	ands	r2, r1, #0x8
	bne		1f
	add		r0, r0, #1
	ands	r2, r1, #0x4
	bne		1f
	add		r0, r0, #1
	ands	r2, r1, #0x2
	bne		1f
	add		r0, r0, #1	
	ands	r2, r1, #0x1
	bne		1f
	add		r0, r0, #1
1:	
	ldr		pc, [sp], #4 
	
#endif
#endif
