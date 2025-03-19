# go-launcher: A simple program launcher written in go

This simple program was created to be used to launch a list of programs and URL's on a daily basis. Once setup, this can be called from either a shortcut or from the path.

## Install

```PowerShell
go install github.com/mikepruett3/go-launcher@latest
```

## Building

When ready to compile, run the following

```powershell
git clone https://github.com/mikepruett3/go-launcher.git

cd go-launcher
go build
```

## Testing

To test your `config.yaml`, run the following

```ps
go run ./main.go
```

The resulting executable & config.yaml file can be placed in your system or user path.

## Configuration

All programs and settings are stored in the `config.yaml` file, which can be located in the following locations

- `/etc/go-launcher/`
- `$HOME/.config/go-launcher/`
- or the current directory

An example `config.yml` has been provided for `Windows (config.yaml.windows.example)` and `Linux (config.yaml.linux.example)` is included with this repo. Rename the file and customize to your tailored experience.

```yaml
browser:
  exec: "C:\\Program Files\\Google Chrome\\chrome.exe"
  profile: Default
  links:
    - https://google.com/
    - https://msn.com/
    - https://reddit.com/

programs:
  calc:
    exec: "C:\\Windows\\System32\\calc.exe"
    start_dir: "C:\\Windows\\System32\\"
    args: ""
```

> `NOTE`: For Windows Program Executables, please enter the full path and executable name in the `exec:` section. Also ensuring to use `\\` (backslash escaping) each folder in the path. (Same with the `start_dir:` paths as well)

Example:

```bash
exec: "C:\\Program Files(x86)\\Application\\app.exe"
```

Hopefully this makes sense

### Browser Section

For the `Browser` section, define the program executable and path to the executable in the `exec:` setting.

If the program executable is defined in your system or user path, then it is safe to omit the full path from the `exec:` setting. (Just leaving the name of the executable)

Define one or multiple `urls` as `links:`

### Program Section

Define each program you wish to launch in the `Program` section.

Each Program should have a name (i.e. `calc:`), and a defined program executable and path to the executable in the `exec:` setting.

If the program executable is defined in your system or user path, then it is safe to omit the full path from the `exec:` setting. (Just leaving the name of the executable)

To specify a different Startup Directory (different from the path that the executable is located in), specify the path in the `start_dir:` setting

For runtime arguments, you can define them using the `args:` setting. Support for only one runtime argument per program. `args:` setting can be omitted if there is no runtime arguments required.

## References

Using [`Viper`](https://github.com/spf13/viper) GoLang module from [spf13](https://github.com/spf13)
