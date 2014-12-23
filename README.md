# USACO Test

Run old USACO contest test cases on your machine from the command-line.

### Installation
    
    $ go get github.com/lordzamy/usacotest
    
### Usage
    
    $ usacotest (path_to_program) [path_to_test_cases_directory]
    
The `path_to_program` argument is required and must be a <b>compiled executable</b> or a Python <b>(.py)</b> script. If your program is named `xyz` then it must take input from `xyz.in` and output to `xyz.out`.

The `path_to_test_cases_directory` argument is optional and assumes the current working directory as the default. This directory must contain the test cases downloaded from the <a href="http://usaco.org/index.php?page=contests">USACO website</a>.
