package services

import (
	"fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/host"
    "time"
    "os/exec"
    "runtime"
    "errors"
)

func GetRamUsage() (map[string]interface{}, error) {
    v, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "totalMB":     v.Total / 1024 / 1024,
        "usedMB":      v.Used / 1024 / 1024,
        "freeMB":      v.Free / 1024 / 1024,
        "usedPercent": v.UsedPercent,
    }, nil
}

func GetCpuUsage() (map[string]interface{}, error) {
    percentages, err := cpu.Percent(1*time.Second, false)
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "cpuUsagePercent": percentages[0],
    }, nil
}

func GetSystemInfo() (map[string]interface{}, error) {
    info, err := host.Info()
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "hostname":       info.Hostname,
        "uptime":         info.Uptime,
        "bootTime":       info.BootTime,
        "os":             info.OS,
        "platform":       info.Platform,
        "platformFamily": info.PlatformFamily,
        "platformVersion": info.PlatformVersion,
        "kernelVersion":  info.KernelVersion,
        "virtualizationSystem": info.VirtualizationSystem,
        "virtualizationRole":   info.VirtualizationRole,
    }, nil
}

func GetCpuTemperature() (map[string]interface{}, error) {
    sensors, err := host.SensorsTemperatures()
    if err != nil {
        return nil, err
    }

    for _, sensor := range sensors {
        if sensor.SensorKey == "cpu_thermal" {
            return map[string]interface{}{
                "cpuTemperature": sensor.Temperature,
            }, nil
        }
    }

    return nil, fmt.Errorf("CPU temperature sensor not found")
}

func ShutdownSystem() error {
    if runtime.GOOS != "linux" {
        return errors.New("unsupported operating system")
    }
    go func() {
        cmd := exec.Command("sudo", "shutdown", "now")
        cmd.Run()
    }()

    return nil
}