.text
	.global _start

_start:
	mov $0, %rax  # rax = 0
	inc %rax      # rax++
	dec %rax      # rax--
	add $7, %rax  # rax += 7
	sub $5, %rax  # rax -= 5

	# At this point, rax = 2.
	# Multiplication cannot use immediate addressing.
	# So we store the multiplier in %rbx.
	# mul assumes the multiplicand is in %rax.
	# The product is also stored in %rax.
	mov $12, %rbx
	mul %rbx # rax *= 12

	# rax = 24.
	# Division also cannot use immediate addressing.
	# Store the divisor in %rbx.
	# div assumes the dividend is in %rax.
	# The quotient is also stored in %rax.
	mov $3, %rbx
	div %rbx

	# rax = 8
	# Let's use this as our exit code.
	# We're using 64-bits now -- the exit code should be in %rdi.
	mov %rax, %rdi
	mov $231, %rax
	syscall
