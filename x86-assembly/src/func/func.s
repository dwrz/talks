.text
	.global _start

_start:
	# Store 3 on the stack.
	push $3
	# call stores return address on the stack.
	call square
	# Use the return value, in %rax, as our exit status number.
	mov %rax, %rdi
	mov $231, %rax
        syscall

# square expects an integer on the stack, 16 bytes down from RBP.
# The squared value is returned on %rax.
square:
	# Function Epilogue --> {
	push %rbp
        mov %rsp, %rbp

	# Retrieve the parameter from the stack.
	# 8 bytes down from RSP is the return address from square.
	# 8 bytes down from that is the parameter to multiple.
	mov 16(%rbp), %rax
	mul %rax

	# Function Prologue --> }
	mov %rbp, %rsp
        pop %rbp
	# ret pops and return to address on stack.
	ret
