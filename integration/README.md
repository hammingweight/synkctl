# Running the Integration Tests
The integration tests need credentials to authenticate against the SunSynk API as well
as the serial number of an inverter. The parameters must be set using environment
variables

```
$ export TEST_USER=carl@example.com
$ export TEST_PASSWORD="verySecret"
$ export TEST_INVERTER_SN=1234567890

```

Then, to run the tests:

```
$ $ go test -count=1 -v -tags integration .
=== RUN   TestBattery
--- PASS: TestBattery (0.31s)
=== RUN   TestGrid
--- PASS: TestGrid (0.26s)
=== RUN   TestInput
--- PASS: TestInput (0.32s)
=== RUN   TestInverter
--- PASS: TestInverter (0.25s)
=== RUN   TestLoad
--- PASS: TestLoad (0.26s)
PASS
ok  	github.com/hammingweight/synkctl/integration	2.372s
```
