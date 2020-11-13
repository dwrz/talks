.data
romeo:
	.ascii "What's in a name? That which we call a rose,\n"
juliet:
	.ascii "By any other name would smell as sweet.\n"
	.ascii "\n"
stop:
	# Try changing this number. What happens? Why?
	# Hint: try 0, 1, 11, 12, 21, 22.
	.int 12

.text
	.global _start

_start:
	# Print the string before it's messed up.
        mov $1, %rax     # sys_write.
        mov $1, %rdi     # stdout
        mov $romeo, %rsi # address
        mov $86, %rdx    # count bytes
        syscall

	# Setup.
	lea romeo, %rax       # Load the starting address of string into %rax.
	mov $0xb98c9ff0, %ebx # 🌹 = b98c9ff0 (4 bytes)
	mov $0, %rcx          # Set the intial counter.

	# Overwrite the strings with 🌹.
	# Increment the address in %rax by four bytes on each iteration.
loop:
	mov %ebx, (%rax) # Copy
	inc %rcx         # increment our counter
	add $4, %rax     # Try changing this to sub. What happens? Why?
	cmp stop, %rcx
	jne loop

	# Print the string after it's been messed up.
	# NB: nothing stopped us from overwriting past the "romeo" string.
        mov $1, %rax     # sys_write
        mov $1, %rdi     # stdout
        mov $romeo, %rsi # address
        mov $85, %rdx    # count bytes
        syscall

	mov $231, %rax # 64-bit system call number for exit_group.
        mov $0, %rdi   # Exit code.
        syscall        # Call the kernel.
