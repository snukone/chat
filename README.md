
[![snukone - chat](https://img.shields.io/static/v1?label=snukone&message=chat&color=blue&logo=github)](https://github.com/snukone/chat "Go to GitHub repo")
[![GitHub tag](https://img.shields.io/github/tag/snukone/chat?include_prereleases=&sort=semver&color=blue)](https://github.com/snukone/chat/releases/)
[![Made with GH Actions](https://img.shields.io/badge/CI-GitHub_Actions-blue?logo=github-actions&logoColor=white)](https://github.com/features/actions "Go to GitHub Actions homepage")
[![Build](https://github.com/snukone/chat/actions/workflows/build.yml/badge.svg)](https://github.com/snukone/chat/actions/workflows/build.yml)
[![Made with Go](https://img.shields.io/badge/Go-1.18-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")
# Chatroom for authenticated users via OAuth2 service provider Github

## Makes use of the following design principles:
- option pattern: makes parameters in structs optional; implemented in trace-module
- decorator pattern: add additional functionality to handlers via wrapper types with helper function (e.g. auth.MustAuth())
- chaining pattern: making handlers dependend of each other