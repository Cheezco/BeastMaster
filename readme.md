# BeastMaster

Lightweight monitoring tool for docker containers.

Designed for windows, but can also be used on other platforms without CPU energy monitoring.

# Modules

1. BeastMaster - main application
2. CheekyChipmunk - logging
3. LazyRaven - container monitoring
4. SleepyCapybara - export

# Usage
```
Usage: BeastMaster.exe [OPTIONS]

Options:
    -Config Configuration file name (default: defaultConfig.yml).
```

# Installation

* Windows:

   * Download BeastMaster binary from releases tab.

# Configuration

Configuration file uses yaml format.

## Structure

* BeastMaster
   * configAddress `(string)` `default: localhost:1000`
* CheekyChipmunk
   * address `(string)` `default: localhost:1200`
   * pluginAddress `(string)` `default: localhost:`
   * loggingPlugins `([]Plugin)`
       * fileName `(string)`
       * launchCommand `(string)` `optional`
* LazyRaven
    * containers `([]Container)`
        * id `(string)`
        * alias `(string)` `optional`
    * parserCount `(int)` `default: 1`
    * address `(string)` `default: localhost:1100`
* SleepyCapybara
   * address `(string)` `default: localhost:1300`
   * pluginAddress `(string)` `default: localhost:1400`
   * exportPlugins `([]Plugin)`
     * fileName `(string)`
     * launchCommand `(string)` `optional`

## Example

For most cases you don't need to specify any address. Just providing container ids and plugins is enough.

```yaml
lazyRaven:
  containers:
    - id: "{ContainerID}"
    - id: "{ContainerID}"
      alias: "example alias"
sleepyCapybara:
  exportPlugins:
    - fileName: "test.exe"
CheekyChipmunk:
  loggingPlugins:
    - fileName: "test2.py"
      launchCommand: "python"
```

# Plugins

## Usage

BeastMaster supports two types of plugins: logging and exporting.

* Export plugins must be placed in `./ExportPlugins/`
* Logging plugins must be placed in `./LoggingPlugins/`

After they're added to their respective folders they must be specified in configuration file.

## Default plugins

WIP

## Development

To create plugins for BeastMaster all you need to do is connect to their module's exposed websocket address. All data sent
to plugins is in json format.

### Logging

To get log data connect to CheekyChipmunk pluginAddress using websocket.

#### Log structure

| Field name | Type   |
|------------|--------|
| Text       | string |
| Source     | string |
| Created    | time   |

### Export

To get export data connect to SleepyCapybara pluginAddress using websocket.

#### Container data structure

| Field name         | Type    |
|--------------------|---------|
| Name               | string  |
| CpuDelta           | int64   |
| SystemCpuDelta     | int64   |
| CpuUsagePercent    | float64 |
| UsedMemory         | int64   |
| MemoryUsagePercent | float64 |
| NumberOfCpus       | int     |
| CpuEnergyUsed      | float64 |
| TimeStamp          | time    |