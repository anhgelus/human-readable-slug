# Human Readable Slug

Lightweight Go library generating a human-readable slug by avoiding some bad combinations like:
- I and l
- 0 and O
- m and n
- or if letter1 = letter2

## Usage

To install this library, run this command:
```bash
go install github.com/anhgelus/human-readable-slug@latest
```

And now you generate some slugs with the function `GenerateSlug` in `github.com/anhgelus/human-readable-slug`:
```go
slug := GenerateSlug(time.Now().Unix(), 7) // generates a 7 chars-length slug with the current timestamp as seed 
```
