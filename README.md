# go-shc
go variant of shc

## Why

On a contract the client was using `shc`, which was fine, but annoying in many instances. After fixing a few bugs in (the usage) of it. I wanted to toy around with creating something a bit more managable and with a high level language. The dumping of `shc` "compiled" scripts was always relatively easy, same with this project, though I wanted to add in some obfuscation, randomness and anti-* features. Most of those exist in a private branch, but the general idea of this project is open sourced.

One of the requirements I had was allowing this to work on Android environments, for reasons not given. This turned out to be a bit of a research aspect as you can't often perform memio work as you would on normal linux systems.

Hopefully in the future I can open source the rest of this and also expand upon the Android findings.

## Usage

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

### License

```
Copyright 2022-23 Tim 'diff' Strazzere <diff@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```