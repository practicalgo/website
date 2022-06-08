---
title: "Errata"
date: 2022-06-06T08:27:41+11:00

---

If you find an error, please let me know: mail@practicalgobook.net

## Chapter 1

*Page 5 (Electronic and Print)*

Current `greetUser()` function is missing a closing paren in the arguments. The fixed version:

```go
func greetUser(c config, name string, w io.Writer) {
        // function body remains same
}
```

(Thanks to Max Wolffe for reporting)
