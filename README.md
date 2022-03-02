# <img src="resources/icon.png" height="24"> alfred-snpt

[![Version](https://img.shields.io/github/release/mike182uk/alfred-snpt.svg?style=flat-square)](https://github.com/mike182uk/alfred-snpt)
[![Build Status](https://img.shields.io/github/workflow/status/mike182uk/alfred-snpt/CI/master?style=flat-square)](https://github.com/mike182uk/alfred-snpt/actions?query=workflow%3ACI)
[![Downloads](https://img.shields.io/github/downloads/mike182uk/alfred-snpt/total.svg?style=flat-square)](https://github.com/mike182uk/alfred-snpt)
[![License](https://img.shields.io/github/license/mike182uk/alfred-snpt.svg?style=flat-square)](https://github.com/mike182uk/alfred-snpt)

An [Alfred](https://www.alfredapp.com/) workflow for [snpt](https://github.com/mike182uk/snpt).

![](example.gif)

## Prerequisites

- [snpt](https://github.com/mike182uk/snpt)

## Installation

Download the latest version of the workflow from  [here](https://github.com/mike182uk/alfred-snpt/releases).

## Usage

1. Trigger Alfred
2. Type `snpt` and press enter or tab
3. Start typing the name of the snippet you want to copy
4. Use the up / down key to navigate to the snippet you want to copy
5. Press enter on the snippet that you want to copy

A notification will be displayed once the snippet has been copied to the clipboard.

## Troubleshooting

### The workflow is reporting that `snpt` can not be found

The workflow will search for `snpt` in:

- `/usr/local/bin/`
- `/usr/bin`
- `/bin`
- `/usr/sbin`
- `/sbin`

If you have installed `snpt` to a different location you will need to modify the workflow to use this path:

1. Go to Alfred preferences
2. Select the `snpt` workflow
3. Double click on the `Script Filter` node
4. In the script field, you should see `export PATH=/usr/local/bin/:$PATH` at the top of the field. Modify this to include your path for `snpt`:

```sh
export PATH=/usr/local/bin/:/my/custom/path/bin/:$PATH
```

Repeat the above for each `Run Script` node in the Alfred workflow.
