# Otto Example App Plugin

This repository contains an example app type plugin for
[Otto](https://www.ottoproject.io). App types enable Otto to detect and
work with new types of applications.

**NOTE:** Otto 0.2 (currently unreleased) is required for app type plugins.
This repository is used for the documentation and as an example in preparation
for that release. Once it is released, this section will be removed.

## Building the Plugin

To build the plugin, compile it like a normal go executable:

```
$ go get
$ go build -o otto-plugin-example
```

The output name of "otto-plugin-example" is important. Otto looks for
the pattern `otto-plugin-*` to find plugins.
