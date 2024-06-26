 
#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"
#if EMBEDDED_ARM>=6

	VARDEF y32, r3
	VARDEF pn1, r4
	VARDEF pn2, r5
	VARDEF lo, r6
	VARDEF hi, r8
.globl	SYM(SKP_Silk_A2NLSF_eval_poly)
SYM(SKP_Silk_A2NLSF_eval_poly):
	stmdb	sp!,  {r4-r8, fp, ip, lr}
	add		fp, sp, #28
	add		r0, r0, r2, lsl #2
	ldr		y32, [r0], #-4
	tst		r2, #1
	beq		LR(1, f)
	ldr		pn1, [r0], #-4
	sub		r2, r2, #1
	mov		r1, r1, lsl #4
L(0)
	smull	lo, hi, y32, r1
	subs	r2, r2, #1
	add		y32, pn1, hi, lsl #16
	ldr		pn1, [r0], #-4
	add		y32, y32, lo, lsr #16
	bgt		LR(0, b)
	smull	lo, hi, y32, r1
	add		y32, pn1, hi, lsl #16
	add		r0, y32, lo, lsr #16
	ldmia	sp!,  {r4-r8, fp, ip, pc}

L(1)/*EVEN:*/
	add		r0, r0, #4
	ldmdb	r0!, {pn1, pn2}
	sub		r2, r2, #2
	mov		r1, r1, lsl #4
L(0)
	smull	lo, hi, y32, r1
	subs	r2, r2, #2
	add		y32, pn2, hi, lsl #16
	add		y32, y32, lo, lsr #16
	smull	lo, hi, y32, r1
	add		y32, pn1, hi, lsl #16
	ldmdb	r0!, {pn1, pn2}
	add		y32, y32, lo, lsr #16
	bgt		LR(0, b)
	smull	lo, hi, y32, r1
	add		y32, pn2, hi, lsl #16
	add		y32, y32, lo, lsr #16
	smull	lo, hi, y32, r1
	add		y32, pn1, hi, lsl #16
	add		r0, y32, lo, lsr #16
	ldmia	sp!,  {r4-r8, fp, ip, pc}
	END
#endif
#endif
