.text
	.global _start

_start:
	# Push some 64-bit integers onto the stack.
	push $3
	push $2
	push $1

	# Let's print "Hey\n", using the stack.
	# man ascii
	# |---------+----+-----+-----+-----|
	# | ASCII   |  H |   e |   y | \n  |
	# |---------+----+-----+-----+-----|
	# | Decimal | 10 | 171 | 145 | 110 |
	# |---------+----+-----+-----+-----|
	# | Hex     | 48 |  65 |  79 | 0A  |
	# |---------+----+-----+-----+-----|

	# Remember that the stack is backwards:
	# it starts at high memory addresses, and "grows" down.
	# However, the write system call follows a low-to-high address order.
	# This means we need to reverse the order of our string.
	#    <-- high           low -->
	#               \n y e H
	push $0x000000000A796548

	# %rsp points to the top of the stack.
	# (Where 'H' is -- a lower address).
	# We can pass this address to the kernel.
	# write will start at 'H', and work up to '\n'.
	mov $1, %rax     # write
	mov $1, %rdi     # stdout
	lea (%rsp), %rsi # address of buffer
	# Try changing the number of bytes to 9, 17, 25.
	# Then call the program with strace.
	# You'll be able to see the previous numbers we pushed onto the stack.
	mov $4, %rdx     # number of bytes
	syscall

	# Remember, the stack is LIFO.
	pop %rdi # pop the "Hey\n" string
	pop %rdi # pop 1
	pop %rdi # pop 2

	# We'll use the last popped value in %rdi as our exit code.
	mov $231, %rax
        syscall
