"""
Defines semantic exit codes which may be used by Command Line tools to aid in
debugging and instrumentation.

This package defines exit codes in two ranges. Exit Codes 80-99 indicate a user error
of some sort. Exit Codes 100-119 indicate software or system error of some sort.
"""

from enum import Enum
import sys

# Code is the exit code that is passed to the system call `exit` when the
# program terminates. Conventionally, the value zero indicates success and all
# other values (1-255) indicate failure.
class Code(Enum):
    # there doesn't really appear to be any convention about letter casing
    # here, so I'm adopting the Go format

    # Indicates that the program exited successfully.
    OK = 0

    # Indicates that the program exited unsuccessfully but gives no extra
    # context as to what the failure was.
    NotOK = 1

    # Exit Codes 80-99 are reserved for user errors.

    # Indicates that the program exited unsuccessfully because it was used
    # incorrectly.
    #
    # Examples: a required argument was omitted or an invalid value was
    # supplied for a flag.
    UsageError = 80

    # Indicates that the program exited unsuccessfully because an unrecognized
    # subcommand was invoked.
    #
    # This is intended for CLI multi-tools. When you run a command that doesn't
    # exist from the shell, the shell exits 127. This is distinct from that
    # value in that the command itself exists but the subcommand does not (e.g.
    # `git nope` could exit 81).
    UnknownSubcommand = 81

    # Indicates that the program exited unsuccessfully because a precondition
    # wasn't satisfied.
    #
    # Examples: the user must be on a VPN before using the program or have
    # a minimum version of some other software installed.
    RequirementNotMet = 82

    # Indicates that the program exited unsuccessfully because the user isn't
    # authorized to perform the requested action.
    Forbidden = 83

    # Indicates that the program exited unsuccessfully because it has been
    # migrated to a new location.
    MovedPermanently = 84

    # Exit Codes 100-119 are reserved for software or system errors.

    # Indicates that the program exited unsuccessfully because of a problem in
    # its own code.
    #
    # Used instead of 1 when the problem is known to be with the program's
    # code or dependencies.
    InternalError = 100

    # Indicates that the program exited unsuccessfully because a service it
    # depends on was not available.
    #
    # Examples: A local daemon or remote service did not respond, a connection
    # was closed unexpectedly, an HTTP service responded with 503.
    Unavailable = 101

    def is_ok(self):
        """Reports whether an exit code is okay.

        Returns True if the code is 0.
        """

        return self.value == 0

    def is_error(self):
        """Reports whether an exit code is an error.

        Returns True if the code is *not* 0.
        """

        return self.value != 0

    def is_user_error(self):
        """Reports whether an exit code is a user error.

        Returns True if the code is in the range 80-99 and False otherwise.
        """

        return 80 <= self.value <= 99

    def is_software_error(self):
        """Reports whether an exit code is a software error.

        Returns True if the code is in the range 100-119 and False otherwise.
        """

        return 100 <= self.value <= 119

    def is_signal(self):
        """Reports whether an exit code is derived from a signal.

        Returns True if the code is in the range 128-255 and False if not.
        """

        return 128 < self.value < 255

    def exit(self):
        """Invoke sys.exit() with this as an exit code."""

        sys.exit(self.value)
