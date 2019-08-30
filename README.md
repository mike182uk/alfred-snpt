# <img src="resources/icon.png" height="24"> Snpt Alfred Workflow

[![Version](https://img.shields.io/github/release/mike182uk/snpt-alfred-workflow.svg?style=flat-square)](https://github.com/mike182uk/snpt-alfred-workflow)
[![Build Status](https://img.shields.io/travis/mike182uk/snpt-alfred-workflow.svg?style=flat-square)](http://travis-ci.org/mike182uk/snpt-alfred-workflow)
[![Downloads](https://img.shields.io/github/downloads/mike182uk/snpt-alfred-workflow/total.svg?style=flat-square)](https://github.com/mike182uk/snpt-alfred-workflow)
[![License](https://img.shields.io/github/license/mike182uk/snpt-alfred-workflow.svg?style=flat-square)](https://github.com/mike182uk/snpt-alfred-workflow)

An [Alfred](https://www.alfredapp.com/) workflow for [Snpt](https://github.com/mike182uk/snpt).

![](example.gif)

## Prerequisites

- [Snpt](https://github.com/mike182uk/snpt)
- [fzf](https://github.com/junegunn/fzf)

## Installation

Download the latest version of the workflow from  [here](https://github.com/mike182uk/snpt-alfred-workflow/releases).

## Usage

1. Trigger Alfred
2. Type `snpt` and press enter or tab
3. Start typing the name of the snippet you want to copy
4. Use the up / down key to navigate to the snippet you want to copy
5. Press enter on the snippet you want to copy

A notification will be displayed once the snippet has been copied to the clipboard.

## Troubleshooting

### The workflow is reporting that `snpt` or `fzf` can not be found

The workflow will search for `snpt` / `fzf` in:

- `/usr/local/bin/`
- `/usr/bin`
- `/bin`
- `/usr/sbin`
- `/sbin`

If you have installed `snpt` / `fzf` to a different location you will need to modify the workflow to use this path:

1. Go to Alfred preferences
2. Select the Snpt workflow
3. Double click on the `Script Filter` node
4. In the script field, you should see `export PATH=/usr/local/bin/:$PATH` at the top of the field. Modify this to include your path for `snpt` / `fzf`:

```sh
export PATH=/usr/local/bin/:/my/custom/path/bin/:$PATH
```

Repeat the above for each `Run Script` node in the Alfred workflow.
