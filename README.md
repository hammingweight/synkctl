# synkctl 
**synkctl** is a CLI and REST Client for SunSynk<sup>:registered:</sup> inverters that allows you to:
 * Read your inverter settings
 * Read the state of the attached power sources (solar panels, grid and battery)
 * Read load statistics
 * Update your inverter settings if you have an [installer account](https://www.sunsynk.org/remote-monitoring)

Since it's written in Go, **synkctl** runs on Linux, Windows and MacOS and on AMD64 and Arm CPUs.

## Getting started with the CLI
Running `synkctl`  without any arguments, provides help

```
$ synkctl 
SynkCtl is a CLI for querying and updating SunSynk hybrid inverters and getting
the state of the battery, grid and input (e.g. solar panels) connected to the
inverter.

Usage:
  synkctl [command]

Available Commands:
  battery       The inverter's battery state and statistics
  completion    Generate the autocompletion script for the specified shell
  configuration Access configuration for the SunSynk API
  grid          The state of the connection to the power grid
  help          Help about any command
  input         The inverter's input (e.g. solar panels, turbine)
  inverter      The inverter's settings
  load          The inverter's load statistics

Flags:
  -c, --config string     synkctl config file location (default "/home/cmeijer/.synk/config")
  -h, --help              help for synkctl
  -i, --inverter string   SunSynk inverter serial number
  -v, --version           version for synkctl

Use "synkctl [command] --help" for more information about a command.
```

### Configuring *synkctl*
To use *synkctl* you need to create a configuration file. The easiest way to do that is to run

```
$ synkctl configuration generate -u <username> -p <password>
```

You might not want to pass your password on the command line; there are two options:
 * Don't specify the password and edit the configuration file afterwards
 * Specify the password via the environment variable `SYNK_PASSWORD`

If you are managing only one inverter, you can specify a default inverter

```
$ synkctl configuration generate -u <username> -p <password> -i <inverter_serial_number>
```

For example,

```
$ synkctl configuration generate -u carl@example.com -p verySecret -i 2401011234
Wrote configuration to '/home/carl/.synk/config'.
```

Then, to view the config file

```
$ cat ~/.synk/config
endpoint: https://api.sunsynk.net
user: carl@example.com
password: verySecret
default_inverter_sn: "2401011234"
```

To check that your credentials work

```
$ synkctl configuration verify
OK.
```
