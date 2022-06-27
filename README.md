# clivm-proxy

[![Build Status](https://github.com/clivm/clivm-proxy/workflows/test/badge.svg)](https://github.com/clivm/clivm-proxy/actions)
[![GitHub last commit](https://img.shields.io/github/last-commit/clivm/clivm-proxy.svg)](https://github.com/clivm/clivm-proxy)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/clivm/clivm-proxy/main/LICENSE)

The internal CLI tool of [clivm](https://github.com/clivm/clivm).
We develop clivm-proxy for clivm, and we don't assume that clivm-proxy is used in the other purpose.

Basically the user of clivm don't have to know the detail of clivm-proxy.
clivm-proxy is installed to `$CLIVM_ROOT_DIR/bin/clivm-proxy` automatically when `clivm install` and `clivm exec` is run, so you don't have to install clivm-proxy explicitly.

clivm-proxy has only the minimum feature and responsibility.
clivm-proxy is stable and isn't changed basically.

clivm-proxy is developed to decide the version of clivm and package managed with clivm dynamically according to the clivm's configuration file when the package is executed.

Please see [How does Lazy Install work?](https://clivm.github.io/docs/reference/lazy-install) too.

## LICENSE

[MIT](LICENSE)
