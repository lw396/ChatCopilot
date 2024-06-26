
#if defined(__arm__)

#include "SKP_Silk_AsmPreproc.h"

#if EMBEDDED_ARM >= 5
	VARDEF A, r2
	VARDEF len, r4
	VARDEF S00, r5
	VARDEF ink1, r6
	VARDEF ink2, r8
	VARDEF Y2, sb
	VARDEF outk1, sl
	VARDEF outk2, ip

.globl	SYM(SKP_Silk_allpass_int)
SYM(SKP_Silk_allpass_int):
	stmdb	sp!,  {r4-r10, fp, ip, lr}
	add		fp, sp, #36
.set	len_offset, 40 
	ldr			len, [sp, #len_offset]
	ldr			S00, [r1]
	movs		A, A, lsl #17
	sub			len, len, #1
	ldmia		r0!, {ink1,ink2}
	bge			LR(1, f)/*AgeZero*/
L(0)
	sub			Y2, ink1, S00
	smlawt		outk1, Y2, A, ink1
	add			S00, Y2, outk1
	subs		len, len, #2
	sub			Y2, ink2, S00
	smlawt		outk2, Y2, A, ink2
#ifdef _WINRT
	ble			LR(3, f)
	ldmia		r0!, {ink1,ink2}
L(3)	
#else	
	ldmgtia		r0!, {ink1,ink2}
#endif
	add			S00, Y2, outk2
	stmia		r3!, {outk1, outk2}
	bgt			LR(0, b)
#ifdef _WINRT
	bne			LR(2, f)
	ldr			ink1, [r0], #4
	sub			Y2, ink1, S00
	smlawt		outk1, Y2, A, ink1
	add			S00, Y2, outk1
	str			outk1, [r3]
L(2)	
#else
	ldreq		ink1, [r0], #4
	subeq		Y2, ink1, S00
	smlawteq	outk1, Y2, A, ink1
	addeq		S00, Y2, outk1
	streq		outk1, [r3]
#endif	
	str			S00, [r1]	
	ldmia	sp!,  {r4-r10, fp, ip, pc}

L(1)/*AgeZero:*/
	sub			Y2, ink1, S00
	smlawt		outk1, Y2, A, S00
	add			S00, Y2, outk1
	subs		len, len, #2
	sub			Y2, ink2, S00
	smlawt		outk2, Y2, A, S00
#ifdef _WINRT
	ble			LR(3, f)
	ldmia		r0!, {ink1,ink2}
L(3)
#else	
	ldmgtia		r0!, {ink1,ink2}
#endif
	add			S00, Y2, outk2
	stmia		r3!, {outk1, outk2}
	bgt			LR(1, b)
#ifdef _WINRT
	bne			LR(2, f)
	ldr			ink1, [r0], #4
	sub			Y2, ink1, S00
	smlawt		outk1, Y2, A, S00
	add			S00, Y2, outk1
	str			outk1, [r3]
L(2)	
#else	
	ldreq		ink1, [r0], #4
	subeq		Y2, ink1, S00
	smlawteq	outk1, Y2, A, S00
	addeq		S00, Y2, outk1
	streq		outk1, [r3]
#endif	
	str			S00, [r1]	
	ldmia	sp!,  {r4-r10, fp, ip, pc}
	END
#endif	
#endif
