.data
str:
        .ascii "Hello, World\n"

.text
	.global _start

_start:
        mov $1, %rax   # 64-bit system call number for sys_write.
        mov $1, %rdi   # stdout.
        mov $str, %rsi # Address of the buffer.
        mov $13, %rdx  # Number of bytes to write.
        syscall

	mov $231, %rax # 64-bit system call number for exit_group.
        mov $0, %rdi   # Exit code.
        syscall
