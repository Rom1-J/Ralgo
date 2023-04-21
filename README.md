# Ralgo
[![Github release action](https://github.com/Rom1-J/Ralgo/workflows/Release/badge.svg)](https://github.com/Rom1-J/Ralgo/actions?query=workflow%3ARelease)
[![Github commit action](https://github.com/Rom1-J/Ralgo/workflows/Building/badge.svg)](https://github.com/Rom1-J/Ralgo/actions?query=workflow%3AGo)

Script to convert large `.txt` files (or any other format) to `.csv` via a regular expression.
Encode text or file into a pseudo-random sequence, and decode it despite the random aspect.

---

## Installation

### From Sources

```bash
$ git clone https://github.com/Rom1-J/Ralgo
$ cd Ralgo
$ make build  # assuming you already have go installed on your system
```

Then you can find the executable inside `dist/` directory.

### From Builds

- [Download the latest release](https://github.com/Rom1-J/Ralgo/releases/latest) compatible with your system

---

## Usage

```bash
$ ./ralgo -h
# -input string
#         Input text
#   -output string
#         Output file (default "stdout")
```

---

## Examples

```bash
$ time ./ralgo 
# todo

$ cd extra/example
$ python verify.py
# Test passed!
```

---

## Performances tests

Specs:
+ CPU: `Intel i7-9750H (12) @ 4.500GHz`
+ Disk: `NVMe`
+ Memory: `32GB`

Sample:
+ todo

Runs:
+ todo


## License
+ todo