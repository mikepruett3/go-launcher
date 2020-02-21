# go-laucher: A simple program laucher written in go

This simple program was created to be used to launch a list of programs and URL's on a daily basis. Once setup, this can be called from either a shortcut or from the path (like MacOS Spotlight)

## Configuration

All programgs and settings are stored in the **config.yml** file, which should be located in the same folder as the go program. An example **config.yml** (**config.yml.example**) is included with this repo. Rename the file and customize to your tailored experience.

```yaml
browser:
  exec: "chrome.exe"
  path: "C:/Program\ Files/Google/Chrome/"
  links:
    - "google.com"
    - "msn.com"
    - "reddit.com"

programs:
- exec: "calc.exe"
  path: "C:/Windows/System32"
  arg: ""
```

**NOTE**: Path settings are entered in unix/MacOS notation. For Windows Path settings, use **/** instead of **\\** , and to use **\\** to escape spaces or other non-unix characters.

ie:

```bash
path: "C:/Program\ Files\(x86\)/Application/"
```

Hoepfully this makes sense

### Browser Section

For the **Browser** section, define the program executable (**exec:**) and path to the executable (**path:**) in the corresponding settings.

If the program executable is defined in your system or user path, then it is safe to omit the **path:** setting.

Define one or multiple **urls** as **links:**

### Program Section

Define each program you wish to launch in the **Program** section.

Each Program should have a defined program executable (**exec:**) and path to the executable (**path:**) in the corresponding settings.

If the program executable is defined in your system or user path, then it is safe to omit the **path:** setting.

For runtime arguments, you can define them using the **arg:** setting. Support for only one runtime argument per program. **arg:** setting can be omitted if there is no runtime argurments required.

## Testing

To test your **config.yml**, run the following

```bash
go run *.go
```

## Building

When ready to compile, run the following

```bash
go build
```

The resulting executable & config.yml file can be placed in your system or user path.

## References

Using [**Configor**](https://github.com/jinzhu/configor) GoLang module from [jinzhu](https://github.com/jinzhu)
