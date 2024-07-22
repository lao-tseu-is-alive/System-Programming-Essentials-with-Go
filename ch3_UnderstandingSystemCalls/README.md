## Chapter 3 : Understanding System Calls
### System calls
System calls, often called “syscalls,” are fundamental to an operating system’s interface. 
They are low-level functions provided by the operating system kernel that allow user-level processes to request services from the kernel.
- [wikipedia syscall](https://en.wikipedia.org/wiki/System_call) 
- System calls are the interface between user space and kernel space.
- System calls are the only way to interact with the kernel.
- System call table is a data structure that contains the addresses of the system calls.[linux syscall table](https://elixir.bootlin.com/linux/v5.14/source/arch/x86/entry/syscalls/syscall_64.tbl)
- System calls are identified by a unique number called a system call number.[Searchable syscall table](https://filippo.io/linux-syscall-table/)
- System calls are implemented in the kernel.

### the syscall package
The syscall package provides low-level interfaces to the operating system.The package has been a foundation for interfacing with system calls and constants across various architectures and operating systems. However, over time, several issues have emerged that led to its deprecation:
- The syscall package is bloated, lacks explicit tests and is a compatibility challenge.
- The syscall package is not portable across different architectures.

for all these reasons, the syscall package has been deprecated in favor of the golang.org/x/sys/unix package.

### The golang.org/x/sys/unix package
The golang.org/x/sys/unix package is a more modern and portable package that provides a more idiomatic way to interact with system calls and constants across different architectures and operating systems. The package is a wrapper around the C library, which provides a more consistent and portable way to interact with system calls and constants.




