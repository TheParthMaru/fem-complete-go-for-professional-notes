# Installing Go

- Download go from https://go.dev/dl/
- Installation instruction https://go.dev/doc/install

```bash
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz
```

- Add the path environment variable.
- $HOME/.profile or /etc/profile (for system wide installation)

```bash
export PATH=$PATH:/usr/local/go/bin
```

- Check the go version after installation.

```bash
go version
```

## Starting a go project

- We can create a directory and within the directory create a module.

```bash
go mod init <module_name>
```

- Module naming conventions: `hosted_on/author_name/module_name`
- Eg: `module github.com/melkeydev/go-blueprint`
- Any go file is saved with the extension `.go`.
- The main file is usually saved with `main.go`.

## Setting up IDEs

### VSCode

- Use the official Go extension by Go team at Google that will be prompted when you create a go file.

### Neovim

- Install the `gopls` LSP using Mason or any other plugin.
