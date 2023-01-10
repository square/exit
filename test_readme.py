import pytest
import exit
import re

with open("README.md") as fp:
    pattern = re.compile("\\| (\\d+) \\| `(\\w+)` \\| .* \\|")
    expected_constants = {}
    for line in fp:
        if result := pattern.match(line):
            expected_constants[result.group(2)] = int(result.group(1))

actual_constants = {code.name: code.value for code in exit.Code}

@pytest.mark.parametrize("code_name", list(set(expected_constants) | set(actual_constants)))
def test_foo(code_name):
    if code_name not in actual_constants:
        assert False, f"exit.py does not define the exit code {code_name} ({expected_constants[code_name]})"

    if code_name not in expected_constants:
        assert False, f"exit.py defines an undocumented exit code {code_name} ({actual_constants[code_name]})"

    if actual_constants[code_name] != expected_constants[code_name]:
        assert False, f"exit.py defines {code_name} as {actual_constants[code_name]}, README.md defines it as {expected_constants[code_name]}"
