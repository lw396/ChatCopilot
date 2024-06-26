 

#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"

#if EMBEDDED_ARM>=5

	VARDEF len, r0
	VARDEF S_0, r3
	VARDEF S_1, r4
	VARDEF S_2, r5
	VARDEF S_3, r6
	VARDEF in, _r7
	VARDEF out1, r8
	VARDEF coef01, sb
	VARDEF coef23, sl
	VARDEF coef45, ip
	VARDEF coef6, r8
	VARDEF coef_tmp, _r7
	VARDEF out2, lr

.set	sp_S, 0 
.set	sp_coef, 4 
.globl	SYM(SKP_Silk_resampler_private_ARMA4)
SYM(SKP_Silk_resampler_private_ARMA4):
	stmdb	sp!,  {r4-r10, fp, ip, lr}
	add		fp, sp, #36
	sub		sp, sp, #8
.set	sp_len, 48 
	
	ldr	lr, [sp, #sp_len]
	ldrh	coef01, [r3], #2
	ldrh	S_1, [r3], #2
	ldrh	coef23, [r3], #2
	ldrh	S_2, [r3], #2
	ldrh	coef45, [r3], #2
	ldrh	S_3, [r3], #2
	ldrh	coef6, [r3], #2
	cmp	lr, #0
	str	r0, [sp, #sp_S]
	beq	LR(1, f)			// if len==0
	
	str	coef6, [sp, #sp_coef]
	ldrsh	in, [r2], #2
	add	coef01, coef01, S_1, lsl #16
	add	coef23, coef23, S_2, lsl #16
	add	coef45, coef45, S_3, lsl #16
	ldmia	r0, {S_0, S_1, S_2, S_3}
	mov	len, lr	
	
L(0)
	mov	in, in, lsl #8
	add	out1, in, S_0, lsl #2
	add	out2, out1, S_2, lsl #2
	
	smlawb	S_0, in, coef01, S_1
	smlawb	S_0, out1, coef23, S_0
	smlawt	S_2, out1, coef01, S_3
	smlawb	S_2, out2, coef45, S_2
	smulwt	S_1, out1, coef23
	smulwt	S_3, out2, coef45
	add	S_1, S_1, in, asr #2
	ldr	coef_tmp, [sp, #sp_coef]
	add	S_3, S_3, out1, asr #2
	
	mov	out1, #128
	smlawb	out1, out2, coef_tmp, out1
#if EMBEDDED_ARM<6	
	cmp	out1, #0x800000
	movge	out1, #0x800000
	subge	out1, out1, #1
	cmn	out1, #0x800000
	movlt	out1, #0x800000
	mov	out1, out1, asr #8
#else		
	ssat	out1, #16, out1, asr #8
#endif
	subs	len, len, #1
	strh	out1, [r1], #2
#ifdef _WINRT
	ble		LR(1, f)
	ldrsh	in, [r2], #2
	b	LR(0, b)
L(1)	
#else	
	ldrgtsh	in, [r2], #2
	bgt	LR(0, b)
#endif	
	
	ldr	r0, [sp, #sp_S]
	stmia	r0, {S_0, S_1, S_2, S_3}
L(1)	
	add		sp, sp, #8
	ldmia	sp!,  {r4-r10, fp, ip, pc}
	END
#endif
#endif
