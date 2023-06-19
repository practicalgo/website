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

*Listing 1.1*

The function definition of `validateArgs()` should check that, the user hasn't
specified the `-h` flag.

Reported as a [GitHub issue](https://github.com/practicalgo/code/issues/5) and
[fixed]() by [titoe218](https://github.com/titoe218). 

(Thanks titoe218 and daniel-lee-tech)

## Chapter 3

*Page 79 (Electronic and Print)*

Listing 3.9, the call to `registerPackageData()` function is missing the second argument.

The published code is correct, and you can find it [here](https://github.com/practicalgo/code/blob/master/chap3/pkgregister-data/pkgregister_test.go#LL57C2-L57C2).

Reported as a [GitHub issue](https://github.com/practicalgo/code/issues/13).

(Thanks rikyhidayat21) 

## Chapter 7

There is a bug in Listing 7.5 described in detail [here](https://github.com/practicalgo/code/issues/11). 

The fixed code is available in the [v1.0.2](https://github.com/practicalgo/code/releases/tag/v1.0.2) release.

(Thanks joergjo)
