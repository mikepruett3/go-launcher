# go-laucher: A simple program laucher written in go

This simple program was created to be used to launch a list of programs and URL's on a daily basis. Once setup, this can be called from either a shortcut or from the path (like MacOS Spotlight)

## Install from source

```bash
go install github.com/mikepruett3/go-launcher@latest
```

## Building

When ready to compile, run the following

```bash
go build
```

## Testing

To test your **config.yml**, run the following

```bash
go run ./main.go
```

The resulting executable & config.yml file can be placed in your system or user path.

## Configuration

All programgs and settings are stored in the **config.yml** file, which should be located in the same folder as the go program. An example **config.yml** (**config.yml.example**) is included with this repo. Rename the file and customize to your tailored experience.

```yaml
browser:
  exec: "C:\\Program Files\\Google Chrome\\chrome.exe"
  links:
    - "google.com"
    - "msn.com"
    - "reddit.com"

programs:
- exec: "C:\\Windows\\System32\\calc.exe"
  startdir: "C:\\"
  arg: ""
```

**NOTE**: For Windows Program Executables, please enter the full path and executable name in the **exec:** section. Also ensuring to use **\\** (backslash escaping) each folder in the path. (Same with the **startdir:** paths as well)

ie:

```bash
exec: "C:\\Program Files(x86)\\Application\\app.exe"
```

Hoepfully this makes sense

### Browser Section

For the **Browser** section, define the program executable and path to the executable in the **exec:** setting.

If the program executable is defined in your system or user path, then it is safe to omit the full path from the **exec:** setting. (Just leaving the name of the executable)

Define one or multiple **urls** as **links:**

### Program Section

Define each program you wish to launch in the **Program** section.

Each Program should have a defined program executable and path to the executable in the **exec:** setting.

If the program executable is defined in your system or user path, then it is safe to omit the full path from the **exec:** setting. (Just leaving the name of the executable)

To specify a different Startup Directory (different from the path that the executable is located in), specify the path in the **startdir:** setting

For runtime arguments, you can define them using the **arg:** setting. Support for only one runtime argument per program. **arg:** setting can be omitted if there is no runtime argurments required.

## References

Using [**Configor**](https://github.com/jinzhu/configor) GoLang module from [jinzhu](https://github.com/jinzhu)
