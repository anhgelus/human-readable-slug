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

And now you generate some slugs with the function `GenerateSlug` in `github.com/anhgelus/human-readable-slug`. E.g.:
```go
slug := GenerateSlug(time.Now().Unix(), 7) // generates a 7 chars-length slug with the current timestamp as seed 
```

You can also check if a slug is human-readable with the function `IsHumanReadable`. E.g.:
```go
IsHumanReadable("aab") // returns false with the default options
IsHumanReadable("Aab") // returns true with the default options
```

If you want to add a custom bad combination, you can use the function `AddBadCombination`.E.g.:
```go
AddBadCombination('b', 'd') // add the combination "bd" and "db"
AddBadCombination('q', 'p') // add the combination "qp" and "pq"
```
### Customize
If you want to customize the options of the slug generation, you must use the `CustomSlug` type.
- `HasCapitalLetters` is true if you want to have capital letters (default: true)
- `HasNumbers` is true if you want to have numbers (default: true)
- `CanBe2ConsecutiveLetters` is true if you want to have 2 consecutive letters or more (default: false)

With this type you can now call every other functions except the `AddBadCombination` which is global.
```go
options := CustomSlug{
    CanBe2ConsecutiveLetters: false,
    HasNumbers:               false,
    HasCapitalLetters:        true,
}
options.GenerateSlug(time.Now().Unix(), 7) // Will generate a 7 chars length slug without numbers
options.IsHumanReadable("helloWorld") // returns true
options.IsHumanReadable("hell0World") // returns false due to the `0`
```
