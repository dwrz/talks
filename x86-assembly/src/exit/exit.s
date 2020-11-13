.text
	.global _start

_start:
	# Store system call 1 in the EAX register.
	# mov = instruction
	# $1, %eax = operands
	# %eax = register
	mov $1, %eax

	# Store the exit code 0 in the EBX register.
	mov $0, %ebx

	# Call interrupt handler 0x80 (128) -- on Linux, this is the kernel.
	# Kernel will take over and see 1 in %eax, 0 in %ebx.
	int $0x80
