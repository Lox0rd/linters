# Linter Results

## 1) Python (bandit)

```bash
user@arch ~/D/B/L/python> bash -c "source bandit-env/bin/activate && bandit linter.py"

[main]  INFO  profile include tests: None
[main]  INFO  profile exclude tests: None
[main]  INFO  cli include tests: None
[main]  INFO  cli exclude tests: None
[main]  INFO  running on Python 3.14.3

Run started:2026-03-23 13:36:39.182636+00:00

Test results:
Issue: [B404:blacklist] Consider possible security implications associated with the subprocess module.
Severity: Low   Confidence: High
CWE: CWE-78 (https://cwe.mitre.org/data/definitions/78.html)
More Info: https://bandit.readthedocs.io/en/1.9.4/blacklists/blacklist_imports.html#b404-import-subprocess
Location: ./linter.py:2:0
1  import os
2  import subprocess
3  import sqlite3

Issue: [B105:hardcoded_password_string] Possible hardcoded password: 'super_secret_password_123'
Severity: Low   Confidence: Medium
CWE: CWE-259 (https://cwe.mitre.org/data/definitions/259.html)
More Info: https://bandit.readthedocs.io/en/1.9.4/plugins/b105_hardcoded_password_string.html
Location: ./linter.py:6:17
5  # Опасность 1: жёстко заданный пароль
6  ADMIN_PASSWORD = "super_secret_password_123"
7  

Issue: [B605:start_process_with_a_shell] Starting a process with a shell, possible injection detected, security issue.
Severity: High   Confidence: High
CWE: CWE-78 (https://cwe.mitre.org/data/definitions/78.html)
More Info: https://bandit.readthedocs.io/en/1.9.4/plugins/b605_start_process_with_a_shell.html
Location: ./linter.py:10:4
9      # Опасность 2: использование os.system с пользовательским вводом
10      os.system(f"ls {user_input}")
11  

Issue: [B608:hardcoded_sql_expressions] Possible SQL injection vector through string-based query construction.
Severity: Medium   Confidence: Low
CWE: CWE-89 (https://cwe.mitre.org/data/definitions/89.html)
More Info: https://bandit.readthedocs.io/en/1.9.4/plugins/b608_hardcoded_sql_expressions.html
Location: ./linter.py:16:14
15      # Опасность 3: SQL-инъекция
16      query = f"SELECT * FROM users WHERE id = {user_id}"
17      cursor.execute(query)

Code scanned:
Total lines of code: 23
Total lines skipped (#nosec): 0

Run metrics:
Total issues (by severity):
Undefined: 0
Low: 2
Medium: 1
High: 1

Total issues (by confidence):
Undefined: 0
Low: 1
Medium: 1
High: 2

Files skipped (0):


⸻

2) Go (golangci-lint)

user@arch ~/D/B/L/go [1]> golangci-lint run

WARN [config_reader] The configuration option run.skip-files is deprecated, please use issues.exclude-files.
WARN [config_reader] The configuration option run.skip-dirs is deprecated, please use issues.exclude-dirs.

linter.go:46:1: File is not properly formatted (gofmt)
      Result:   result,
^


⸻

3) C (clang-tidy)

user@arch ~/D/B/L/C> clang-tidy linter.c -- -I.

1 warning generated.

/home/user/Desktop/BRPO/Linters/C/linter.c:6:9: warning: Call to function 'fprintf' is insecure as it does not provide security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'fprintf_s' in case of C11 [clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling]
6 |         fprintf(stderr, "Memory allocation failed\n");
  |         ^~~~~
/home/user/Desktop/BRPO/Linters/C/linter.c:6:9: note: Call to function 'fprintf' is insecure as it does not provide security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'fprintf_s' in case of C11
6 |         fprintf(stderr, "Memory allocation failed\n");
  |         ^~~~~