# Contributing

Contributions are **welcome!**

Contributions can be made via a Pull Request on [Github](https://github.com/mike182uk/snpt-alfred-workflow).

## Reporting an Issue

Please report issues via the issue tracker on [Github](https://github.com/mike182uk/snpt-alfred-workflow). For security-related issues, please email the maintainer directly.

## Pull Requests

- **Lint changes** - Make sure you run `make lint` before committing your code.

- **Document any change in behaviour** - Make sure the README and any other relevant documentation are kept up-to-date.

- **Create topic branches** - i.e `feature/some-awesome-feature`.

- **One pull request per feature**

- **Send coherent history** - Make sure each individual commit in your pull request is meaningful. If you had to make multiple intermediate commits while developing, please squash them before submitting.

## Install project dependencies

```bash
make install-env-deps
make install
```

You will need to do this before you make any changes.

## Running the Tests

```bash
make test
```

## Building the Helper

```bash
make build-helper
```

## Building the Workflow

```bash
make build-workflow
```
