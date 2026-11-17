<!--suppress ALL -->
<img height="310" alt="image" src="https://user-images.githubusercontent.com/16114089/121805566-0cd81280-cc4c-11eb-9b7d-b5f8a6db4f8d.png" align="right">

# ExyuLine Software

ExyuLine is a heavily asynchronous server software for Minecraft: Bedrock Edition written in Go. It was written with scalability
and simplicity in mind and aims to make the process of setting up a server and modifying it easy. Unlike other
Minecraft server software, ExyuLine is generally used as a library to extend.

[![Discord Banner 2](https://discordapp.com/api/guilds/623638955262345216/widget.png?style=banner2)](https://discord.gg/U4kFWHhTNR)

## Getting started
Running ExyuLine requires at least **Go 1.26**. After starting the server through one of the methods below,
**ctrl+c** may be used to shut down the server. Also check out the [wiki](https://github.com/texterLine/exyuline/wiki) for
more detailed info.

#### Installation as library
```shell
go mod init github.com/user/module
go get github.com/texterLine/exyuline
```

![SetupLibrary](https://user-images.githubusercontent.com/16114089/121804512-0f843900-cc47-11eb-9320-d195393b5a1f.gif)

#### Installation of the latest commit
```shell
git clone https://github.com/texterLine/exyuline
cd exyuline
go run main.go
```

![SetupClone](https://user-images.githubusercontent.com/16114089/121804495-ff6c5980-cc46-11eb-8e31-df4d94782e5b.gif)

## Developer info
[![Go Reference](https://pkg.go.dev/badge/github.com/texterLine/exyuline/server.svg)](https://pkg.go.dev/github.com/texterLine/exyuline/server)

ExyuLine features a well-documented codebase with an easy-to-use API. Documentation may be found
[here](https://pkg.go.dev/github.com/texterLine/exyuline/server) and in the subpackages found by clicking *Directories*.

Publishing your project on GitHub? Consider adding the **[#df-mc](https://github.com/topic/df-mc)** topic to your
repository and opening a pull request at [df-wiki](https://github.com/df-mc/wiki) adding your project as a Community
Project to improve its visibility.

## Contributing
Contributions are very welcome! Issues, pull requests and feature requests are highly appreciated. Opening a pull
request? Consider joining our [Discord server](https://discord.gg/U4kFWHhTNR) to discuss your changes! Also have a read through the
[CONTRIBUTING.md](https://github.com/texterLine/exyuline/blob/master/.github/CONTRIBUTING.md) for more info.
