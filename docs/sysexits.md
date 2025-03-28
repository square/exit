# Comparison to sysexits.h Exit Codes

| Semantic Exit Code | sysexits.h Exit Code | Name | Meaning | 
| --: | --: | :-- | :-- | 
|  80 | 64 | **Usage Error** | A required argument was omitted or an invalid value was supplied for a flag. |
|  81 |    | **Unknown Command** | An unrecognized subcommand was invoked. |
|  82 |    | **Requirement Not Met** | A prerequisite wasn't met. |
|  83 | 77 | **Forbidden** | The user isn't authorized to perform the requested action. |
|  84 |    | **Moved Permanently** | The program has been migrated to a new location. |
|     | 65 | **Data Error** | The input data was incorrect in some way. This should only be used for user's data and not system file. |
|     | 66 | **No Input** | An input file (not a system file) did not exist or was not readable. This could also include errors like "No message" to a mailer (if it cared to catch it). |
|     | 67 | **No User** | The user specified did not exist. This might be used for mail addresses or remote logins. |
|     | 68 | **No Host** | The host specified did not exist. This is used in mail addresses of network requests. |
|     | 73 | **Cannot Create** | A (user specified) output file cannot be created. |
|     | 78 | **Config Error** | Something was found in an unconfigured or misconfigured state. |
| 100 | 70 | **Internal Error** | An error occurred in the program's own code or in one of its dependencies. |
| 101 | 69 | **Unavailable** | A local daemon or remote service did not respond, a connection was closed unexpectedly, an HTTP service responded with 503. |
|     | 71 | **Os Error** | An operating system error has been detected. This is intended to be used for such things as "cannot fork", "cannot create pipe", or the like. It includes things like getuid returning a user that does not exist in the passwd file. |
|     | 72 | **Os File Error** | Some system file (e.g., `/etc/passwd`, `/var/run/utmp`, etc.) does not exist, cannot be opened, or has some sort of error (e.g., syntax error). |
|     | 74 | **Io Error** | An error occurred while doing I/O on some file. |
|     | 75 | **Temporary Failure** | Temporary failure, indicating something that is not really an error. In `sendmail`, this means that a mailer (e.g.) could not create a connection, and the request should be reattempted later. |
|     | 76 | **Protocol Error** | The remote system returned something that was "not possible" during a protocol exchanged. |

<br>

#### Why did we depart from the `sysexits.h` codes?

This library's goals are to define exit codes that are:
1. Broadly applicable to heterogenous Command Line tools
2. Easy to partition into user errors and system errors

`sendmail`'s exit codes weren't ideally suited to either goal. While several of its codes are broadly applicable (like **Usage Error**, **Internal Error**, and **Forbidden**), others are [over-fit](https://en.wikipedia.org/wiki/Overfitting) to `sendmail` (like **No User** and **No Host**). And its numbering scheme interleaves user errors with software errors.
