.data
stop:
	.int 3

.text
	.global _start

_start:
	mov $0, %rcx
loop:
	cmp stop, %rcx
	# If equal, change RIP to exit.
	je exit
	inc %rcx
	# Unconditionnaly change RIP to loop.
	jmp loop
exit:
	# Try this: comment out the following three instructions.
	# What happens if you assemble and link? Why?
	# Think about RIP and the Fetch-Decode-Execute cycle.
	mov $231, %rax
        mov %rcx, %rdi # Exit code -- let's see the value of %rcx.
        syscall
