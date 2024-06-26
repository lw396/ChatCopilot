 

#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"

#if EMBEDDED_ARM>=5

	VARDEF	len, r4
	VARDEF	S_0, r5
	VARDEF	S_1, r6
	VARDEF	in0, _r7
	VARDEF	out32, r8
	VARDEF	A_Q140, sb
	VARDEF	A_Q141, sl

.set	sp_S, 0 
.globl	SYM(SKP_Silk_resampler_private_AR2)
SYM(SKP_Silk_resampler_private_AR2):
	stmdb	sp!,  {r4-r10, fp, ip, lr}
	add		fp, sp, #36
	sub		sp, sp, #4
.set	sp_len, 44 
	
	str	r0, [sp, #sp_S]
	ldr	S_0, [r0], #4
	ldr	S_1, [r0], #-4
	ldr	len, [sp, #sp_len]
	ldrsh	A_Q140, [r3], #2
	ldrsh	A_Q141, [r3]
	cmp	len, #0
	beq	LR(1, f)

L(0)
	ldrsh	in0, [r2], #2
	add	out32, S_0, in0, lsl #8
	str	out32, [r1], #4
	subs	len, len, #1
	mov	out32, out32, lsl #2
	smlawb	S_0, out32, A_Q140, S_1
	smulwb	S_1, out32, A_Q141
	bgt	LR(0, b)	
	
	str	S_0, [r0], #4
	str	S_1, [r0]
L(1)	
	add		sp, sp, #4
	ldmia	sp!,  {r4-r10, fp, ip, pc}
	END
#endif
#endif
