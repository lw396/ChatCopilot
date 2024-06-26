 

#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"

#if EMBEDDED_ARM>=5
#define SKP_Silk_resampler_up2_lq_0 0x1FA6
#define SKP_Silk_resampler_up2_lq_1 0x8FAF

	VARDEF xy0, r4
	VARDEF xy1, r5
	VARDEF up2_coefs, _r7
	VARDEF S_0, r6
	VARDEF S_1, r8
	VARDEF in, sb
	VARDEF out0, sl
	VARDEF out1, r0

.set	sp_S, 0 
.globl	SYM(SKP_Silk_resampler_up2)
SYM(SKP_Silk_resampler_up2):
	stmdb	sp!,  {r4-r10, fp, ip, lr}
	add		fp, sp, #36
	sub		sp, sp, #4
	mov	up2_coefs, #0x1F000000
	ldrsh	in, [r2], #2
	add	up2_coefs, up2_coefs, #0xA60000
	str	r0, [sp, #sp_S]
	add	up2_coefs, up2_coefs, #0x8F00
	ldmia	r0, {S_0, S_1}
	add	up2_coefs, up2_coefs, #0xAF
	mov	ip, #512
	cmp	r3, #0
	beq	LR(1, f)
L(0)
	mov	in, in, lsl #10
	sub	xy0, in, S_0
	sub	xy1, in, S_1
	smulwt	xy0, xy0, up2_coefs
	smlawb	xy1, xy1, up2_coefs, xy1
	add	out0, S_0, xy0
	add	out1, S_1, xy1
	add	S_0, in, xy0
	add	S_1, in, xy1
	qadd	out0, out0, ip
	qadd	out1, out1, ip
#if EMBEDDED_ARM<6	
	cmp	out0, #0x2000000
	movge	out0, #0x2000000
	subge	out0, out0, #1
	cmn	out0, #0x2000000
	movlt	out0, #0x2000000
	cmp	out1, #0x2000000
	movge	out1, #0x2000000
	subge	out1 ,out1, #1
	cmn	out1, #0x2000000
	movlt	out1, #0x2000000
	mov	out0, out0, asr #10
	mov	out1, out1, asr #10
#else	
	ssat	out0, #16, out0, asr #10
	ssat	out1, #16, out1, asr #10
#endif	
	subs	r3, r3, #1
#ifdef _WINRT
	ble		LR(2, f)
	ldrsh	in, [r2], #2
	strh	out0, [r1], #2
	strh	out1, [r1], #2
	b		LR(0, b)
L(2)
	strh	out0, [r1], #2
	strh	out1, [r1], #2
#else	
	ldrgtsh	in, [r2], #2
	strh	out0, [r1], #2
	strh	out1, [r1], #2
	bgt	LR(0, b)
#endif	
	
	ldr	r0, [sp, #sp_S]
	stmia	r0, {S_0, S_1}
L(1)
	add		sp, sp, #4
	ldmia	sp!,  {r4-r10, fp, ip, pc}
	END
#endif
#endif
