# Semantic Exit Codes

### Usage

#### Go

```golang
os.Exit(exit.Forbidden) // The user isn't permitted to perform this action
os.Exit(exit.Unavailable) // An API this program consumes isn't available
```

#### Rust

```rust
use semantic_exit::{exit, Code};

exit(Code::Forbidden);
exit(Code::Unavailable);
```

#### Python

```python
import exit

exit.Code.Forbidden.exit()
exit.Code.Unavailable.exit()
```

See [the complete list of exit codes](#the-codes).

### About

Conventionally, exiting a program with zero indicates success while nonzero indicates failure.

```golang
os.Exit(0) // success
os.Exit(1) // failure
```

But the system call `exit` accepts values between 0 and 255, leaving 254 different ways of expressing failure.

This library's goals are to define exit codes that are:
1. Broadly applicable to heterogenous Command Line tools
2. Easy to partition into user errors and system errors

It defines codes in two [unreserved](#reserved-codes-and-prior-art) ranges: 80-99 for user errors and 100-119 for software or system errors.

### The Codes

| Exit Code | Name | Meaning |
| --: | :-- | :-- |
| 0 | `OK` | The program exited successfully. |
| 1 | `NotOK` | The program exited unsuccessfully but gives no extra context as to what the failure was. |
| 80 | `UsageError` | The program exited unsuccessfully because it was was used incorrectly. (e.g. a required argument was omitted or an invalid value was supplied for a flag.) |
| 81 | `UnknownSubcommand` | The program exited unsuccessfully because an unrecognized subcommand was invoked. (Used by CLI multi-tools.) |
| 82 | `RequirementNotMet` | The program exited unsuccessfully because a prerequisite of it wasn't met. |
| 83 | `Forbidden` | The program exited unsuccessfully because the user isn't authorized to perform the requested action. |
| 84 | `MovedPermanently` | The program exited unsuccessfully because it has been migrated to a new location. |
| 100 | `InternalError` | The program exited unsuccessfully because of a problem in its own code. (Used instead of 1 when the problem is known to be with the program's code or dependencies.) |
| 101 | `Unavailable` | The program exited unsuccessfully because a service it depends on was not available. (e.g. A local daemon or remote service did not respond, a connection was closed unexpectedly, an HTTP service responded with 503.) |

### Reserved Codes and Prior Art

- Values above 128 are reserved for signals. (When a program is terminated with a signal, its exit code is 128 + the signal's numeric value. When you terminate a program with `Ctrl` `C`, for example, you send it the signal `SIGINT` — whose value is 2 — and the program exits with 130.) 
- Bash [reserves 2, 126, and 127](https://tldp.org/LDP/abs/html/exitcodes.html).
- [sysexits.h defines 64-78](https://github.com/bminor/glibc/blob/master/misc/sysexits.h#L96-L110). The `sysexits.h` codes were originally defined for `sendmail` but have been used many places since. ([Compare Semantic Exit Codes to sysexits.h codes](./docs/sysexits.md))
