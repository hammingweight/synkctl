# synkctl 
**synkctl** is a CLI and REST Client for SunSynk<sup>:registered:</sup> inverters that allows you to:
 * Read your inverter settings
 * Read the state of the attached power sources (solar panels, grid and battery)
 * Read load statistics
 * Update your inverter settings if you have an [installer account](https://www.sunsynk.org/remote-monitoring)

Since it's written in Go, **synkctl** runs on Linux, Windows and MacOS and on AMD64 and Arm CPUs.

## The **synkctl** CLI
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
  -k, --keys string       Extract specific keys from response
  -v, --version           version for synkctl

Use "synkctl [command] --help" for more information about a command.
```

### Configuring **synkctl**
To use **synkctl** you need to create a configuration file. The easiest way to do that is to run

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

### Listing all inverters
To see all the inverters that you can inspect, run `synkctl inverters list`, e.g.

```
$ synkctl inverter list
[
    "2401010001",
    "2401020123"
]
```

### Reading state and statistics
**synkctl** can read the state of statistics of the inverter, battery, input (e.g. panels) and grid by passing the "get" verb on the associated object:

```
synkctl inverter get
synkctl battery get
synkctl input get
synkctl grid get
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

The output from `get` can be very lengthy. If you are only interested in certain fields you can specify those fields as comma-separated values following the `-k` switch

```
$ synkctl -i 2201020123 grid get -k etodayFrom,fac
{
    "etodayFrom": "0.5",
    "fac": 49.91
}

```

If you want to extract nested values, you should use a more sophisticated tool like `jq` and pipe the output to the tool. For example

```
$ synkctl -i 2201020123 grid get | jq .vip[0].volt
"236.2"

```


## The **synkctl** REST Client
To use the (Go) REST client, you need to
 * Create a `Configuration` instance (typically by reading it from a config file)
 * Call an `Authenticate` function which, if successful, returns a `SynkClient` object
 * Invoke `Read` or `Update` methods on the `SynkClient`

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
	fmt.Println("Battery SOC:\t\t", battery.SOC())

	// We can update inverter settings. For example, increase the lower
	// threshold for the battery capacity by 5% or allow the inverter to
	// power non-essential circuits if the solar panels are producing
	// more than 1000W.
	inverter, _ := client.Inverter(ctx)
	oldBatteryCapacity := inverter.BatteryCapacity()
	newBatteryCapacity := oldBatteryCapacity + 5
	inverter.SetBatteryCapacity(newBatteryCapacity)
	if input.Power() > 1000 {
		inverter.SetLimitedToLoad(false)
	}

	// Write the updated settings to the API.
	err := client.UpdateInverter(ctx, inverter)
	if err != nil {
		fmt.Println("failed to update inverter settings: ", err)
	}
}
```
