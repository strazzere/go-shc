# go-shc
go variant of shc

## usage

```
Usage: go-shc SCRIPT

Options:
  -arch arch
    	arch to compile the script to, defaults to arm64 (default "arm64")
  -direct
    	directly pipe the script to an interpreter (no file used), defaults to false
  -garble
    	use garble instead of go to build the file, defaults to false
  -interpreter direct
    	if direct was used, attempt to use this specific interpreter, defaults to sh with no direct path (default "sh")
  -os arch
    	arch to compile the script to, defaults to arm64 (default "android")
```

If you want to use `garble` then you would need to ensure it is installed on the host machine while compiling a script.
