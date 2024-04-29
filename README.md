## ocr-translator

**THIS PROJECT IS STILL UNDER DEVELOPMENT**

### Dependencies

- [libleptonica](https://github.com/DanBloomberg/leptonica/)
- [libtesseract](https://github.com/tesseract-ocr/tesseract)

### Build

#### Linux

Arch:

```bash
pacman -S leptonica
pacman -S tesseract
```

Additionally, you should install training data corresponding to the language
you want to use for `tesseract`. For example

English:

```bash
pacman -S tesseract-data-eng
```

Japanese:

```bash
pacman -S tesseract-data-jpn
pacman -S tesseract-data-jpn_vert
```

Then return to the root of the project and run:

```bash
go build .
```

**NOTE**: Screenshot on linux is unstable to use. You may
write `scanner.Scanner` yourself on linux platform.

#### Windows

It's little tricky to install dependencies on Windows. Because this
project uses `cgo` to call `tesseract` and `leptonica` functions, we
should use `MinGW` as `cgo` compiler.

We recommend that use [`MSYS2`](https://www.msys2.org/) to
install `MinGW` and other dependencies, so that we have no
necessary to build `tesseract` and `leptonica` from source.

After `MSYS2` is installed, open `MSYS2 MinGW 64-bit` shell and:

Install compile environment.

```bash
pacman -S base-devel msys2-devel mingw-w64-x86_64-toolchain git
```

Install dependencies for `tesseract` and `leptonica`.

```bash
pacman -S mingw-w64-x86_64-asciidoc mingw-w64-x86_64-cairo mingw-w64-x86_64-curl mingw-w64-x86_64-icu mingw-w64-x86_64-leptonica mingw-w64-x86_64-libarchive mingw-w64-x86_64-pango mingw-w64-x86_64-zlib mingw-w64-x86_64-autotools mingw-w64-x86_64-cmake
```

Install `tesseract` and `leptonica`.

```bash
pacman -S mingw-w64-x86_64-leptonica
pacman -S mingw-w64-x86_64-tesseract-ocr
```

You could install training data for `tesseract` as mentioned in the
Linux section.

Set Windows environment variables:
- Set MinGW path to `PATH` environment variable. Concretely, add
  `YOUR_MSYS2_PATH\mingw64\bin` to Windows environment variables `PATH`.
- Set Windows environment variables `TESSDATA_PREFIX` to `YOUR_MSYS2_PATH\mingw64\share\tessdata`.

Then return to the root of the project and run:

```bash
go build .
```

### Usage

Write config file, here's an example:

```yaml
ocr:
  type: tesseract
  scan-delay: 2000
  lang:
    - jpn
    - eng
translators:
  - type: google
    api: YOUR_GOOGLEAPI_KEY
    target: zh-chs
    source: ja
    socks5: "127.0.0.1:8888"
```
