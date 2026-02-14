A go formatter that uses goimports under the hood but allows for disabling
formatting for certain lines.

**Example:**

```
//nofmt:enable
[custom formatted code]
//nofmt:disable
```


**Usage:**

```
Usage:
  nofmt [OPTIONS] [File]

Options:
  -h, --help  Show this help message and exit

Positional arguments:
  File        The file to format. Reads from Stdin if not specified.

```
