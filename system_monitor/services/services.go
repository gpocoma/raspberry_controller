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
    "raspberry-controller/system_monitor/models"
)

func GetRamUsage() (models.RamUsage, error) {
    v, err := mem.VirtualMemory()
    if err != nil {
        return models.RamUsage{}, err
    }

    return models.RamUsage{
        TotalMB:     v.Total / 1024 / 1024,
        UsedMB:      v.Used / 1024 / 1024,
        FreeMB:      v.Free / 1024 / 1024,
        UsedPercent: v.UsedPercent,
    }, nil
}

func GetCpuUsage() (models.CpuUsage, error) {
    percentages, err := cpu.Percent(1*time.Second, false)
    if err != nil {
        return models.CpuUsage{}, err
    }

    return models.CpuUsage{
        CpuUsagePercent: percentages[0],
    }, nil
}

func GetSystemInfo() (models.SystemInfo, error) {
    info, err := host.Info()
    if err != nil {
        return models.SystemInfo{}, err
    }

    return models.SystemInfo{
        Hostname:            info.Hostname,
        Uptime:              info.Uptime,
        BootTime:            info.BootTime,
        OS:                  info.OS,
        Platform:            info.Platform,
        PlatformFamily:      info.PlatformFamily,
        PlatformVersion:     info.PlatformVersion,
        KernelVersion:       info.KernelVersion,
        VirtualizationSystem: info.VirtualizationSystem,
        VirtualizationRole:   info.VirtualizationRole,
    }, nil
}

func GetCpuTemperature() (models.CpuTemperature, error) {
    sensors, err := host.SensorsTemperatures()
    if err != nil {
        return models.CpuTemperature{}, err
    }

    for _, sensor := range sensors {
        if sensor.SensorKey == "cpu_thermal" {
            return models.CpuTemperature{
                CpuTemperature: sensor.Temperature,
            }, nil
        }
    }

    return models.CpuTemperature{}, fmt.Errorf("CPU temperature sensor not found")
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