# Linter Results

## 1) Python (bandit)

```bash
venv ❯ bandit -r linter.py                                                                                                                                                                                           09:17:13
[main]  INFO    profile include tests: None
[main]  INFO    profile exclude tests: None
[main]  INFO    cli include tests: None
[main]  INFO    cli exclude tests: None
[main]  INFO    running on Python 3.14.4
Run started:2026-04-21 02:18:31.251678+00:00

Test results:
>> Issue: [B608:hardcoded_sql_expressions] Possible SQL injection vector through string-based query construction.
   Severity: Medium   Confidence: Low
   CWE: CWE-89 (https://cwe.mitre.org/data/definitions/89.html)
   More Info: https://bandit.readthedocs.io/en/1.9.4/plugins/b608_hardcoded_sql_expressions.html
   Location: ./linter.py:15:14
14          
15          query = f"SELECT * FROM users WHERE username = '{username}' AND password = '{password}'"
16          cursor.execute(query)

--------------------------------------------------
>> Issue: [B201:flask_debug_true] A Flask app appears to be run with debug=True, which exposes the Werkzeug debugger and allows the execution of arbitrary code.
   Severity: High   Confidence: Medium
   CWE: CWE-94 (https://cwe.mitre.org/data/definitions/94.html)
   More Info: https://bandit.readthedocs.io/en/1.9.4/plugins/b201_flask_debug_true.html
   Location: ./linter.py:55:4
54          
55          app.run(debug=True)

--------------------------------------------------

Code scanned:
        Total lines of code: 34
        Total lines skipped (#nosec): 0

Run metrics:
        Total issues (by severity):
                Undefined: 0
                Low: 0
                Medium: 1
                High: 1
        Total issues (by confidence):
                Undefined: 0
                Low: 1
                Medium: 1
                High: 0
Files skipped (0):


```

## 2) Go (golangci-lint)


```bash
❯ golangci-lint run                                                                                                                                                                                                  10:30:01
linter.go:4:1: File is not properly formatted (gofmt)
 "fmt"
^
linter.go:5:2: SA1019: "io/ioutil" has been deprecated since Go 1.19: As of Go 1.16, the same functionality is now provided by package [io] or package [os], and those implementations should be preferred in new code. See the specific function documentation for details. (staticcheck)
 "io/ioutil"
 ^
linter.go:18:40: ST1013: should use constant http.StatusUnauthorized instead of numeric literal 401 (staticcheck)
  http.Error(w, "Invalid credentials", 401)
                                       ^
3 issues:
* gofmt: 1
* staticcheck: 2
```

## 3) C (clang-tidy)


```bash
❯ clang-tidy linter.c                                                                                                                                                                                                10:31:24
Error while trying to load a compilation database:
Could not auto-detect compilation database for file "linter.c"
No compilation database found in /home/lev/Рабочий стол/linters/C or any parent directory
fixed-compilation-database: Error while opening fixed database: No such file or directory
json-compilation-database: Error while opening JSON database: No such file or directory
Running without flags.
6 warnings generated.
/home/lev/Рабочий стол/linters/C/linter.c:9:5: error: Call to function 'strcat' is insecure as it does not provide bounding of the memory buffer. Replace unbounded copy functions with analogous functions that support length arguments such as 'strlcat'. CWE-119 [clang-analyzer-security.insecureAPI.strcpy,-warnings-as-errors]
    9 |     strcat(path, filename);
      |     ^~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:9:5: note: Call to function 'strcat' is insecure as it does not provide bounding of the memory buffer. Replace unbounded copy functions with analogous functions that support length arguments such as 'strlcat'. CWE-119
    9 |     strcat(path, filename);
      |     ^~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:19:12: error: The 1st argument to 'fgets' is a buffer with size 128 but should be a buffer with size equal to or greater than the value of the 2nd argument (which is 512) [clang-analyzer-unix.StdCLibraryFunctions,-warnings-as-errors]
   19 |     while (fgets(buffer, sizeof(buffer) * 4, file)) {
      |            ^
/home/lev/Рабочий стол/linters/C/linter.c:45:9: note: Assuming 'argc' is >= 2
   45 |     if (argc < 2) {
      |         ^~~~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:45:5: note: Taking false branch
   45 |     if (argc < 2) {
      |     ^
/home/lev/Рабочий стол/linters/C/linter.c:51:5: note: Calling 'read_file'
   51 |     read_file(argv[1]);
      |     ^~~~~~~~~~~~~~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:11:18: note: Assuming that 'fopen' is successful
   11 |     FILE *file = fopen(path, "r");
      |                  ^~~~~~~~~~~~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:12:10: note: 'file' is non-null
   12 |     if (!file) {
      |          ^~~~
/home/lev/Рабочий стол/linters/C/linter.c:12:5: note: Taking false branch
   12 |     if (!file) {
      |     ^
/home/lev/Рабочий стол/linters/C/linter.c:17:5: note: 'buffer' initialized here
   17 |     char buffer[128];
      |     ^~~~~~~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:19:12: note: The 1st argument to 'fgets' is a buffer with size 128 but should be a buffer with size equal to or greater than the value of the 2nd argument (which is 512)
   19 |     while (fgets(buffer, sizeof(buffer) * 4, file)) {
      |            ^     ~~~~~~  ~~~~~~~~~~~~~~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:20:16: error: format string is not a string literal (potentially insecure) [clang-diagnostic-format-security,-warnings-as-errors]
   20 |         printf(buffer);
      |                ^~~~~~
/home/lev/Рабочий стол/linters/C/linter.c:20:16: note: treat the string as an argument to avoid this
   20 |         printf(buffer);
      |                ^
      |                "%s", 
/home/lev/Рабочий стол/linters/C/linter.c:32:5: error: Call to function 'scanf' is insecure as it does not provide bounding of the memory buffer or security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'scanf_s' in case of C11 [clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,-warnings-as-errors]
   32 |     scanf("%s", username);
      |     ^~~~~
/home/lev/Рабочий стол/linters/C/linter.c:32:5: note: Call to function 'scanf' is insecure as it does not provide bounding of the memory buffer or security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'scanf_s' in case of C11
   32 |     scanf("%s", username);
      |     ^~~~~
/home/lev/Рабочий стол/linters/C/linter.c:35:5: error: Call to function 'scanf' is insecure as it does not provide bounding of the memory buffer or security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'scanf_s' in case of C11 [clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,-warnings-as-errors]
   35 |     scanf("%s", password);
      |     ^~~~~
/home/lev/Рабочий стол/linters/C/linter.c:35:5: note: Call to function 'scanf' is insecure as it does not provide bounding of the memory buffer or security checks introduced in the C11 standard. Replace with analogous functions that support length arguments or provides boundary checks such as 'scanf_s' in case of C11
   35 |     scanf("%s", password);
      |     ^~~~~
Suppressed 1 warnings (1 in non-user code).
Use -header-filter=.* or leave it as default to display errors from all non-system headers. Use -system-headers to display errors from system headers as well.
5 warnings treated as errors


```

## 4) Rust (clippy)


```bash
❯ cargo clippy                                                                                                                                                                                                       13:13:47
    Checking lint_test v0.1.0 (/home/lev/Рабочий стол/linters/lint_test)
warning: the loop variable `i` is only used to index `numbers`
 --> src/main.rs:5:14
  |
5 |     for i in 0..numbers.len() {
  |              ^^^^^^^^^^^^^^^^
  |
  = help: for further information visit https://rust-lang.github.io/rust-clippy/rust-1.95.0/index.html#needless_range_loop
  = note: `#[warn(clippy::needless_range_loop)]` on by default
help: consider using an iterator
  |
5 -     for i in 0..numbers.len() {
5 +     for <item> in &numbers {
  |

warning: writing `&Vec` instead of `&[_]` involves a new object where a slice will do
  --> src/main.rs:73:21
   |
73 | fn print_data(data: &Vec<i32>) {
   |                     ^^^^^^^^^
   |
   = help: for further information visit https://rust-lang.github.io/rust-clippy/rust-1.95.0/index.html#ptr_arg
   = note: `#[warn(clippy::ptr_arg)]` on by default
help: change this to
   |
73 - fn print_data(data: &Vec<i32>) {
73 + fn print_data(data: &[i32]) {
   |

warning: the loop variable `i` is used to index `data`
  --> src/main.rs:74:14
   |
74 |     for i in 0..data.len() {
   |              ^^^^^^^^^^^^^
   |
   = help: for further information visit https://rust-lang.github.io/rust-clippy/rust-1.95.0/index.html#needless_range_loop
help: consider using an iterator and enumerate()
   |
74 -     for i in 0..data.len() {
74 +     for (i, <item>) in data.iter().enumerate() {
   |

warning: `lint_test` (bin "lint_test") generated 3 warnings
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.16s


```