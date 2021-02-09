<p align="center">
<img src="assets/images/logo.svg" alt="frothy Logo">
</p>

# frothy

Cross-platform TOTP client.

## Building

### Prerequisites

* make
* go >= 1.16
* makensis (for building windows)
* create-dmg (for building macos)

### Build
#### on mac

1) install dependencies

```sh
    brew install go nsis create-dmg
```
2) build

```sh
    make
```

#### on windows
1) Install [Git Bash](https://git-scm.com/download/win)
2) Install other dependencies (recommended via [scoop](https://scoop.sh/))

    ```sh
        scoop bucket add nsis https://github.com/NSIS-Dev/scoop-nsis # add bucket
        scoop install go nsis/nsis
    ```
3) build (via [Git Bash](https://git-scm.com/download/win))

    ```sh
        make windows # crossbuilding for macOS is not supported
    ```