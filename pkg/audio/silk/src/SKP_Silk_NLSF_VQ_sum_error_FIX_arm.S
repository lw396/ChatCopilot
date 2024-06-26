 

#if __arm__
#include "SKP_Silk_AsmPreproc.h"

#if ( EMBEDDED_ARM >= 6 )

	VARDEF val_N, r4
	VARDEF val_K, r5
	VARDEF val_order, r6
	VARDEF val_in_Q15, _r7
	VARDEF val_cb_Q15, r8
	VARDEF val_wtmp_Q6, sb
	VARDEF val_sum_error, sl
	VARDEF val_diff, ip
	VARDEF val_tmp, lr

	VARDEF val_in_Q15_tmp, ip
	VARDEF val_wtmp_Q6_tmp, lr

	VARDEF ptr_wtmp, _r7
	VARDEF val0_wtmp, sb
	VARDEF val1_wtmp, r5
	VARDEF val2_wtmp, r8

	VARDEF ptr_err_Q20, r0
	VARDEF ptr_in_Q15, r1
	VARDEF ptr_w_Q6, r2
	VARDEF ptr_pCB_Q15, r3

.set	sp_wtmp, 0 
.set	sp_err_Q20, 32 
.set	sp_in_Q15, 36 
.set	sp_w_Q6, 40 
.set	sp_pCB_Q15, 44 
.globl	SYM(SKP_Silk_NLSF_VQ_sum_error_FIX)
SYM(SKP_Silk_NLSF_VQ_sum_error_FIX):
	stmdb	sp!,  {r4-r10, fp, ip, lr}
	add		fp, sp, #36
	sub		sp, sp, #48
.set	ptr_N, 88 
.set	ptr_K, 92 
.set	ptr_LPC_order, 96 

	str	r0, [sp, #sp_err_Q20]
	str	r1, [sp, #sp_in_Q15]
	str	r2, [sp, #sp_w_Q6]
	str	r3, [sp, #sp_pCB_Q15]
	ldr	val_N, [sp, #ptr_N]

	ldr	val_order, [sp, #ptr_LPC_order]
	ldr	val1_wtmp, [r2], #4
	ldr	val2_wtmp, [r2], #4
	sub	val_order, val_order, #2
	add	ptr_wtmp, sp, #sp_wtmp
L(3)
	pkhbt	val0_wtmp, val1_wtmp, val2_wtmp, lsl #16
	subs	val_order, val_order, #2
	ldr	val1_wtmp, [r2], #4
	ldr	val2_wtmp, [r2], #4
	str	val0_wtmp, [ptr_wtmp], #4
	bgt	LR(3, b)
	pkhbt	val0_wtmp, val1_wtmp, val2_wtmp, lsl #16
	str	val0_wtmp, [ptr_wtmp], #4

	ands	val_tmp, ptr_pCB_Q15, #3
	bgt	LR(4, f)
/*OUTTER_LOOP*/
L(2)
	ldr	ptr_pCB_Q15, [sp, #sp_pCB_Q15]
	ldr	val_K, [sp, #ptr_K]
/*MIDDLE_LOOP*/  
L(1)
	ldr	ptr_in_Q15, [sp, #sp_in_Q15]
	add	ptr_w_Q6, sp, #sp_wtmp
	ldr	val_order, [sp, #ptr_LPC_order]
	mov	val_sum_error, #0
/*INNER_LOOP*/
L(0)
	ldmia	ptr_in_Q15!, {val_in_Q15, val_in_Q15_tmp}
	ldr	val_wtmp_Q6, [ptr_w_Q6], #4
	ldr	val_cb_Q15, [ptr_pCB_Q15], #4
	subs	val_order, val_order, #2
	pkhbt	val_in_Q15, val_in_Q15, val_in_Q15_tmp, lsl #16
	ssub16	val_diff, val_in_Q15, val_cb_Q15
	smulbb	val_tmp, val_diff, val_diff
	smultt	val_diff, val_diff, val_diff
	smlawb	val_sum_error, val_tmp, val_wtmp_Q6, val_sum_error
	smlawt	val_sum_error, val_diff, val_wtmp_Q6, val_sum_error
	bgt	LR(0, b)
	subs	val_K, val_K, #1
	str	val_sum_error, [ptr_err_Q20], #4
	bgt	LR(1, b)
	subs	val_N, val_N, #1
	str	ptr_in_Q15, [sp, #sp_in_Q15]
	bgt	LR(2, b)

	add		sp, sp, #48
	ldmia	sp!,  {r4-r10, fp, ip, pc}

L(4)
	ldr	ptr_pCB_Q15, [sp, #sp_pCB_Q15]
	ldr	val_K, [sp, #ptr_K]
/*MIDDLE_LOOP*/  
L(1)
	ldr	ptr_in_Q15, [sp, #sp_in_Q15]
	add	ptr_w_Q6, sp, #sp_wtmp
	ldr	val_order, [sp, #ptr_LPC_order]
	mov	val_sum_error, #0
/*INNER_LOOP*/
L(0)
	ldmia	ptr_in_Q15!, {val_in_Q15, val_in_Q15_tmp}
	ldr	val_wtmp_Q6, [ptr_w_Q6], #4
	ldrh	val_cb_Q15, [ptr_pCB_Q15], #2
	ldrh	val_tmp, [ptr_pCB_Q15], #2
	subs	val_order, val_order, #2
	pkhbt	val_in_Q15, val_in_Q15, val_in_Q15_tmp, lsl #16
	pkhbt	val_cb_Q15, val_cb_Q15, val_tmp, lsl #16
	ssub16	val_diff, val_in_Q15, val_cb_Q15
	smulbb	val_tmp, val_diff, val_diff
	smultt	val_diff, val_diff, val_diff
	smlawb	val_sum_error, val_tmp, val_wtmp_Q6, val_sum_error
	smlawt	val_sum_error, val_diff, val_wtmp_Q6, val_sum_error
	bgt	LR(0, b)
	subs	val_K, val_K, #1
	str	val_sum_error, [ptr_err_Q20], #4
	bgt	LR(1, b)
	subs	val_N, val_N, #1
	str	ptr_in_Q15, [sp, #sp_in_Q15]
	bgt	LR(4, b)

	add		sp, sp, #48
	ldmia	sp!,  {r4-r10, fp, ip, pc}
	END
#endif
#endif
