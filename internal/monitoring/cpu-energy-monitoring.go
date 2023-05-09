package monitoring

import (
	"BeastMaster/internal"
	"BeastMaster/internal/debug"
	"strconv"
)

func RunCpuMonitor(storage *CpuEnergyStorage) error {
	for {
		scanner, cmd := internal.RunExecutableWithScanner("./HardwareMonitor.exe", "-1")
		for scanner.Scan() {
			convertedValue, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			storage.Update(convertedValue)
		}
		err := cmd.Wait()
		if err != nil {
			return err
		}
		debug.Log("CPU monitor stopped. Restarting...")
	}

}
