# synkctl 
<img src="/images/synkctl.jpg" align="right" width="200px">

![](https://github.com/hammingweight/synkctl/actions/workflows/build.yml/badge.svg)
![](https://github.com/hammingweight/synkctl/actions/workflows/integrationtest.yml/badge.svg)

SunSynk<sup>:registered:</sup> is a manufacturer of popular hybrid inverters that can be managed through a mobile app, web interface,
or direct access to the inverter. While the inverter's user interfaces are functional, they may not be ideal for automated management
of settings. 

The **synkctl** tool simplifies the monitoring of an inverter's state and the application of automatic updates to its settings.

**synkctl** is a CLI and Go REST Client for SunSynk<sup>:registered:</sup> inverters that allows you to:
 * Read your inverter settings
 * Read the state of the attached power sources (solar panels, grid and battery)
 * Read load statistics
 * Update your inverter settings if you have an [installer account](https://www.sunsynk.org/remote-monitoring)

**synkctl** is written in Go and uses the [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) libraries.

## The **synkctl** CLI

### Getting the **synkctl** CLI
You can download **synkctl** for Windows, macOS or Linux from [Releases](https://github.com/hammingweight/synkctl/releases).

If you have Go installed, you can get the latest version by running

```
go install github.com/hammingweight/synkctl@latest
```


### Using the **synkctl** CLI
Running **synkctl**  without any arguments, provides help

```
$ synkctl 
synkctl is a CLI for querying and updating SunSynk hybrid inverters and getting
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
  -c, --config string     synkctl config file location
  -h, --help              help for synkctl
  -i, --inverter string   SunSynk inverter serial number
  -v, --version           version for synkctl

Use "synkctl [command] --help" for more information about a command.
```

#### Configuring **synkctl**
To use **synkctl** you need to create a configuration file; the easiest way to generate a
file is to run

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

#### Listing all inverters
To see all the inverters that you can inspect, run `synkctl inverters list`, e.g.

```
$ synkctl inverters list
[
    "2401010001",
    "2401020123"
]
```

#### Reading state and statistics
**synkctl** can read the state of the inverter, battery, input (e.g. panels), grid and load by passing the "get" verb on the associated object:

```
synkctl inverter get
synkctl battery get
synkctl input get
synkctl grid get
synkctl load get
```

If you have not specified a default inverter serial number (or you want to override the default serial number), you can pass the serial number as a command line argument (using the `-i` or `--inverter` switch); for example,

```
$ synkctl -i 2401020123 grid get
{
    "acRealyStatus": 1,
    "etodayFrom": "0.5",
    "etodayTo": "0.0",
    "etotalFrom": "1944.6",
    "etotalTo": "20.4",
    "fac": 49.91,
    "limiterPowerArr": [
        0,
        0
    ],
    "limiterTotalPower": 0,
    "pac": 0,
    "pf": 1,
    "qac": 0,
    "status": 0,
    "vip": [
        {
            "current": "1.3",
            "power": 0,
            "volt": "236.2"
        }
    ]
}
```

The output from `get` can be very lengthy; if you are only interested in certain fields, you can specify those fields as comma-separated values following the `-k` switch

```
$ synkctl -i 2201020123 grid get -k etodayFrom,fac
{
    "etodayFrom": "0.5",
    "fac": 49.91
}

```

If you want to extract nested values, you should use a more sophisticated tool like **jq** and pipe the output to the tool. For example

```
$ synkctl -i 2201020123 grid get | jq .vip[0].volt
"236.2"

```

#### Updating the inverter settings
There are two verbs for updating the inverter's settings:
 * `update` for the common use-cases
 * `apply` for fine-grained updates to the inverter

  _Unless you have an installer account, attempts to update your inverter settings will fail._
  
##### `update`
The `update` operation allows you to
 * Set the minimum battery SOC (i.e. the SOC at which the inverter will use the grid to power circuits - assuming that the grid is up, obviously)
 * Enable or disable providing power to the CT coil (i.e. allowing or preventing the inverter from powering non-essential circuits)
 * Enable or disable recharging of the battery from the grid
 * Prioritize charging the battery or powering the load

For example, to ensure that the battery won't be discharged below 50% if the grid is up

```
$ synkctl inverter update --battery-capacity 50
```

To allow the inverter to power non-essential circuits via the CT coil

```
$ synkctl inverter update --essential-only off
```

To prioritize recharging the battery if the battery SOC drops below the discharge threshold

```
$ synkctl inverter update --battery-first on
```

To allow the grid to recharge the battery if the battery SOC drops below the discharge threshold

```
$ synkctl inverter update --grid-charge on
```
To see the current inverter settings that can be updated, run

```
$ synkctl inverter settings 
{
    "battery-capacity": 30,
    "battery-first": "on",
    "essential-only": "on",
    "grid-charge": "off"
}
```

##### `apply`
The `update` operation is limited and coarse. For example, a SunSynk inverter allows an operator to set up to six different battery SOCs (`cap1` to `cap6`) depending on the time of day,
but the `update` operation will set all the values `cap1` to `cap6` to the same value. For fine-grained updates to the inverter settings, use the `apply` subcommand and a JSON file (or pipe JSON to stdin)
with the inverter settings.

To get the current inverter settings and all fields that can be updated (and to redirect the output to a JSON file),
you could run

```
$ synkctl inverter get --short` > settings.json
```

(The `--short` flag directive ensures that only updateable inverter settings will be returned.) 

Edit the JSON file with your new settings and use the `apply` subcommand

```
$ vi settings.json
$ synkctl inverter apply -f settings.json force`
```

Note that the `force` argument must be supplied to acknowledge that you are doing something potentially dangerous: There is no validation of the settings.

## The **synkctl** REST Client
A CLI can be useful but for more complex scenarios, it's better to run a program that monitors and adjust settings by making API calls. For example:
 * At the end of each day, check the battery SOC and adjust the minimum SOC (for example, increase the minimum SOC as the seasons change from summer to winter)
 * Only allow the inverter to power non-essential circuits if the battery SOC is above some threshold and the input is producing some minimum amount of power (this ensures that the battery won't be drained too rapidly)

### Installing the REST Client
To add the **synkctl** module as a dependency to a project, run

```
$ go get github.com/hammingweight/synkctl@latest
```


### Using the REST Client
To use the REST client, you need to
 * Create a `Configuration` instance (typically by reading it from a config file)
 * Call an `Authenticate` function which, if successful, returns a `SynkClient` object
 * Invoke `Read` or `Update` methods on the `SynkClient`

The Go code below is illustrative of what's required (although you shouldn't discard returned `error`s)

```
package main

import (
	"context"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
)

func main() {
	configFile, _ := configuration.DefaultConfigurationFile()
	config, _ := configuration.ReadConfigurationFromFile(configFile)
	ctx := context.Background()
	client, _ := rest.Authenticate(ctx, config)

	// Read Input (e.g. solar panels) and display the total energy
	// generated (etotal)
	input, _ := client.Input(ctx)
	eTotal, _ := input.Get("etotal")
	fmt.Println("Total energy generated:\t", eTotal)

	// For some useful attributes, there are convenience methods. For exmple
	// battery.SOC() is equivalent to battery.Get("bmsSoc")
	battery, _ := client.Battery(ctx)
	soc, _ := battery.SOC()
	fmt.Println("Battery SOC:\t\t", soc)

	// We can update inverter settings. For example, increase the lower
	// threshold for the battery capacity by 5% or allow the inverter to
	// power non-essential circuits if the solar panels are producing
	// more than 1000W.
	inverter, _ := client.Inverter(ctx)
	oldBatteryCapacity, _ := inverter.BatteryCapacity()
	newBatteryCapacity := oldBatteryCapacity + 5
	inverter.SetBatteryCapacity(newBatteryCapacity)
	power, _ := input.Power()
	if power > 1000 {
		inverter.SetLimitedToLoad(false)
	}

	// Write the updated settings to the API.
	err := client.UpdateInverter(ctx, inverter)
	if err != nil {
		fmt.Println("failed to update inverter settings: ", err)
	}
}
```
