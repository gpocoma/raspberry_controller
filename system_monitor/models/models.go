package models

type RamUsage struct {
    TotalMB     uint64  `json:"totalMB"`
    UsedMB      uint64  `json:"usedMB"`
    FreeMB      uint64  `json:"freeMB"`
    UsedPercent float64 `json:"usedPercent"`
}

type CpuUsage struct {
    CpuUsagePercent float64 `json:"cpuUsagePercent"`
}

type SystemInfo struct {
    Hostname            string `json:"hostname"`
    Uptime              uint64 `json:"uptime"`
    BootTime            uint64 `json:"bootTime"`
    OS                  string `json:"os"`
    Platform            string `json:"platform"`
    PlatformFamily      string `json:"platformFamily"`
    PlatformVersion     string `json:"platformVersion"`
    KernelVersion       string `json:"kernelVersion"`
    VirtualizationSystem string `json:"virtualizationSystem"`
    VirtualizationRole   string `json:"virtualizationRole"`
}

type CpuTemperature struct {
    CpuTemperature float64 `json:"cpuTemperature"`
}