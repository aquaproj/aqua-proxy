# aqua-proxy

[![Build Status](https://github.com/aquaproj/aqua-proxy/workflows/test/badge.svg)](https://github.com/aquaproj/aqua-proxy/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/aquaproj/aqua-proxy.svg)](https://github.com/aquaproj/aqua-proxy)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/aquaproj/aqua-proxy/main/LICENSE)

The internal CLI tool of [aqua](https://github.com/aquaproj/aqua).
We develop aqua-proxy for aqua, and we don't assume that aqua-proxy is used in the other purpose.

Basically the user of aqua don't have to know the detail of aqua-proxy.
aqua-proxy is installed to `$AQUA_ROOT_DIR/bin/aqua-proxy` automatically when `aqua install` and `aqua exec` is run, so you don't have to install aqua-proxy explicitly.

aqua-proxy has only the minimum feature and responsibility.
aqua-proxy is stable and isn't changed basically.

aqua-proxy is developed to decide the version of aqua and package managed with aqua dynamically according to the aqua's configuration file when the package is executed.

Please see [How does Lazy Install work?](https://aquaproj.github.io/docs/reference/lazy-install) too.

## LICENSE

[MIT](LICENSE)
