//! Defines semantic exit codes which may be used by Command Line tools to aid in debugging and instrumentation.
//!
//! This package defines exit codes in two ranges:
//! - Exit Codes 80-99 indicate a user error of some sort.
//! - Exit Codes 100-119 indicate software or system error of some sort.

/// The exit code that is passed to the system call `exit` when the program terminates.
/// Conventionally, the value zero indicates success and all other values (1-255) indicate failure.
#[derive(Clone, Copy, Debug, PartialEq, num_derive::FromPrimitive, num_derive::ToPrimitive)]
#[repr(i32)]
pub enum Code {
    /// Indicates that the program exited successfully.
    OK = 0,

    /// Indicates that the program exited unsuccessfully but gives no extra context as to what the failure was.
    NotOK = 1,

    // Exit Codes 80-99 are reserved for user errors.
    /// UsageError indicates that the program exited unsuccessfully because it was was used incorrectly.
    ///
    /// Examples: a required argument was omitted or an invalid value was supplied for a flag.
    UsageError = 80,

    /// Indicates that the program exited unsuccessfully because an unrecognized subcommand was invoked.
    ///
    /// This is intended for CLI multi-tools.
    /// When you run a command that doesn't exist from the shell, the shell exits 127.
    /// This is distinct from that value in that the command itself exists but the subcommand does not (e.g. `git nope` could exit 81).
    UnknownSubcommand = 81,

    /// Indicates that the program exited unsuccessfully because a precondition wasn't satisfied.
    ///
    /// Examples: the user must be on a VPN before using the program or have a minimum version of some other software installed.
    RequirementNotMet = 82,

    /// Indicates that the program exited unsuccessfully because the user isn't authorized to perform the requested action.
    Forbidden = 83,

    /// Indicates that the program exited unsuccessfully because it has been migrated to a new location.
    MovedPermanently = 84,

    // Exit Codes 100-119 are reserved for software or system errors.
    /// Indicates that the program exited unsuccessfully because of a problem in its own code.
    ///
    /// Used instead of 1 when the problem is known to be with the program's code or dependencies.
    InternalError = 100,

    /// Indicates that the program exited unsuccessfully because a service it depends on was not available.
    ///
    /// Examples: A local daemon or remote service did not respond, a connection was closed unexpectedly, an HTTP service responded with 503.
    Unavailable = 101,
}

#[derive(Debug, thiserror::Error, PartialEq)]
pub enum Error {
    #[error("unknown exit code: {0}")]
    UnknownExitCode(i32),
}

/// Reports whether an exit code is a user error.
/// It returns true if the code is in the range 80-99 and false if not.
pub fn is_user_error(code: Code) -> bool {
    (code as i32) >= 80 && (code as i32) <= 99
}

/// Reports whether an exit code is a software error.
/// It returns true if the code is in the range 100-119 and false if not.
pub fn is_software_error(code: Code) -> bool {
    (code as i32) >= 100 && (code as i32) <= 119
}

/// Reports whether an exit code is derived from a signal.
/// It returns true if the code is in the range 128-255 and false if not.
pub fn is_signal(code: Code) -> bool {
    (code as i32) > 128 && (code as i32) < 255
}

/// Returns the exit code that corresponds to when a program exits in response to a signal.
pub fn from_signal(signal: i32) -> i32 {
    128 + signal
}

impl TryFrom<i32> for Code {
    type Error = Error;

    fn try_from(value: i32) -> Result<Self, Self::Error> {
        num_traits::FromPrimitive::from_i32(value).ok_or(Error::UnknownExitCode(value))
    }
}

impl From<Code> for i32 {
    fn from(value: Code) -> Self {
        value as i32
    }
}

pub fn exit(code: Code) {
    std::process::exit(code.into());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_exit_code_match_readme() {
        let readme = include_str!("../README.md");
        let re = regex::Regex::new(r"\| (\d+) \| `(\w+)` \| .* \|\n").unwrap();
        for captures in re.captures_iter(readme) {
            let expected_code_num: i32 = captures.get(1).unwrap().as_str().parse().unwrap();
            let expected_name = captures.get(2).unwrap().as_str();
            let actual_code: Code = expected_code_num.try_into().expect(&format!(
                "does not define the exit code {expected_name} ({expected_code_num})"
            ));
            let actual_name = format!("{:?}", actual_code);

            assert_eq!(expected_name, actual_name, "maps {actual_name} to {expected_code_num}, README.md expected it to be {expected_name}");
        }
    }

    #[test]
    fn test_unknown_exit_code() {
        let err = <i32 as TryInto<Code>>::try_into(-1).unwrap_err();
        assert_eq!(err, Error::UnknownExitCode(-1))
    }

    #[test_case::test_case(libc::SIGINT, 130)]
    fn test_from_signal(signal: i32, expected: i32) {
        assert_eq!(from_signal(signal), expected);
    }
}
