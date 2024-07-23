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
The syscall package provides low-level interfaces to the operating system.The package has been a foundation for interfacing with system calls and constants across various architectures and operating systems.
However, over time, several issues have emerged that led to its deprecation:
- The syscall package is bloated, lacks explicit tests and is a compatibility challenge.
- The syscall package is not portable across different architectures.

for all these reasons, the syscall package has been deprecated in favor of the golang.org/x/sys/unix package.

### The golang.org/x/sys/unix package
The golang.org/x/sys/unix package is a more modern and portable package that provides a more idiomatic way to interact with system calls and constants across different architectures and operating systems.
The package is a wrapper around the C library, which provides a more consistent and portable way to interact with system calls and constants.

### The  os package
The os package provides a higher-level interface to the operating system. 
It is a wrapper around the golang.org/x/sys/unix package and provides a more idiomatic way to interact with the operating system.
The os package provides a more user-friendly way to interact with the operating system, such as file and directory operations, environment variables, and command-line arguments.

#### file and directory operations:

+ os.Create(): Creates or opens a file for writing
+ os.Mkdir() and os.MkdirAll(): Create directories
+ os.Remove() and os.RemoveAll(): Remove files and directories
+ os.Stat(): Get file or directory information (metadata)
+ os.IsExist(), os.IsNotExist(), and os.IsPermission(): Check file/directory existence or permission errors
+ os.Open(): Open a file for reading
+ os.Rename(): Rename or move a file
+ os.Truncate(): Resize a file
+ os.Getwd(): Get the current working directory
+ os.Chdir(): Change the current working directory
+ os.Args: Command-line arguments
+ os.Getenv(): Get environment variables
+ os.Setenv(): Set environment variables

#### process operations:

+ os.Getpid(): Get the current process ID
+ os.Getppid(): Get the parent process ID
+ os.Getuid() and os.Getgid(): Get the user and group IDs
+ os.Geteuid() and os.Getegid(): Get the effective user and group IDs
+ os.StartProcess(): Start a new process
+ os.Exit(): Exit the current process
+ os.Signal: Represents signals (for example, SIGINT, SIGTERM)
+ os/signal.Notify(): Notify on signal reception

### key points:
+ use the os package for most tasks, as it provides a safer and portable interface to the underlying operating system.
+ use the golang.org/x/sys/unix package for low-level system calls and constants.
+ avoid using the syscall package, as it has been deprecated in favor of the golang.org/x/sys/unix package.
+ handle errors returned by system calls to ensure the robustness of your application.
+ test your code to ensure that it works as expected across different architectures and operating systems.