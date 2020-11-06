.text
	.global _start

_start:
	# Store system call 1 in the EAX register.
	mov $1, %eax

	# Store the exit code 0 in the EBX register.
	mov $0, %ebx

	# Call interrupt handler 0x80 (128) -- on Linux, this is the kernel.
	int $0x80
