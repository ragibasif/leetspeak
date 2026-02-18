# L33tspeak Converter

A simple command-line tool written in Go that converts text into leetspeak (l33t speak).

## What is Leetspeak?

Leetspeak is an informal language or code that substitutes letters with numbers or symbols that look visually similar (e.g. `A` → `4`, `E` → `3`, `S` → `5`).

## Usage

```bash
go run main.go <word> [word2] ...
```

### Examples

```bash
go run main.go hello
# Output: #3110

go run main.go leetspeak
# Output: 1_3375p34k

go run main.go hello world
# Output: #3110//0r1_d
```

## Character Map

| Letter | Leet |
|--------|------|
| A | 4 |
| B | 8 |
| C | ( |
| E | 3 |
| G | 6 |
| H | # |
| I | 1 |
| J | _ |
| L | 1_ |
| O | 0 |
| Q | 0_ |
| S | 5 |
| T | 7 |
| V | / |
| W | // |
| X | >< |
| Y | `/ |
| Z | 2 |

Letters not in the map are left unchanged.

## Build

```bash
go build -o leet main.go
./leet hello world
```
