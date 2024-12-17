package ui

import (
	"fmt"
	"goscope/internal/monitor"
	"goscope/pkg/utils"
	"sync"
	"time"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/net"
)

func StartGUI() {
	app := tview.NewApplication()
	textView := tview.NewTextView()
	textView.SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetBorder(true).
		SetTitle(" 🖥️  SENTINEL SYSTEM MONITOR ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tview.Styles.BorderColor)
	
	var diskIOStatsMutex sync.RWMutex
	diskIOStats, _ := monitor.GetDiskIOStats()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			type systemStats struct {
				cpuUsage         float64
				cpuCount         int
				memUsage         float64
				memUsed          uint64
				memTotal         uint64
				diskUsage        float64
				diskUsed         uint64
				diskTotal        uint64
				netStats         []net.IOCountersStat
				swapTotal        uint64
				swapUsed         uint64
				diskFree         uint64
				memFree          uint64
				cpuIdle          float64
				diskInodesUsed   uint64
				diskInodesFree   uint64
				cpuFrequencyInfo []cpu.InfoStat
				diskReadBytes    uint64
				diskWriteBytes   uint64
				cpuTemp          float64
				memCached        uint64
				memBuffers       uint64
				netPacketsSent   uint64
				netPacketsRecv   uint64
				netPacketsDrop   uint64
				cpuUserTime      float64
				cpuSystemTime    float64
				swapFree         uint64
				diskTotalInodes  uint64
				memTotalUsed     uint64
				diskTotalUsed    uint64
			}

			statsChan := make(chan systemStats, 1)

			go func() {
				stats := systemStats{}
				stats.cpuUsage, _ = monitor.GetCPUUsage()
				stats.cpuCount, _ = monitor.GetCPUCount()
				stats.memUsage, stats.memUsed, _ = monitor.GetMemoryUsage()
				stats.memTotal, _ = monitor.GetMemoryTotal()
				stats.diskUsage, stats.diskUsed, _ = monitor.GetDiskUsage()
				stats.diskTotal, _ = monitor.GetDiskTotal()
				stats.netStats, _ = monitor.GetNetworkStats()
				stats.swapTotal, _ = monitor.GetMemorySwapTotal()
				stats.swapUsed, _ = monitor.GetMemorySwapUsed()
				stats.diskFree, _ = monitor.GetDiskFree()
				stats.memFree, _ = monitor.GetMemoryFree()
				stats.cpuIdle, _ = monitor.GetCPUIdle()
				stats.diskInodesUsed, stats.diskInodesFree, _ = monitor.GetDiskInodes()
				stats.cpuFrequencyInfo, _ = monitor.GetCPUFrequency()
				stats.diskReadBytes, _ = monitor.GetDiskReadBytes()
				stats.diskWriteBytes, _ = monitor.GetDiskWriteBytes()
				stats.cpuTemp, _ = monitor.GetCPUTemperature()
				stats.memCached, _ = monitor.GetMemoryCached()
				stats.memBuffers, _ = monitor.GetMemoryBuffers()
				stats.netPacketsSent, _ = monitor.GetNetworkPacketsSent()
				stats.netPacketsRecv, _ = monitor.GetNetworkPacketsReceived()
				stats.netPacketsDrop, _ = monitor.GetNetworkPacketsDropped()
				stats.cpuUserTime, _ = monitor.GetCPUUserTime()
				stats.cpuSystemTime, _ = monitor.GetCPUSystemTime()
				stats.swapFree, _ = monitor.GetMemorySwapFree()
				stats.diskTotalInodes, _ = monitor.GetDiskTotalInodes()
				stats.memTotalUsed, _ = monitor.GetMemoryTotalUsed()
				stats.diskTotalUsed, _ = monitor.GetDiskTotalUsed()
				statsChan <- stats
			}()

			stats := <-statsChan
			cpuFrequency := stats.cpuFrequencyInfo[len(stats.cpuFrequencyInfo)-1]
			output := fmt.Sprintf(
    "\n[yellow::b]╔══════════════════════ SYSTEM RESOURCES ════════════════════╗[::-]\n\n"+
    "[yellow::b]CPU INFORMATION\n"+
    "  [red]► Usage: %.2f%% | Cores: [green]%d | Idle: [blue]%.2f%%\n"+
    "  [red]► Temperature: %.2f°C | User Time: [cyan]%.2f | System Time: [cyan]%.2f\n"+
    "  [red]► Model: [magenta]%s\n"+
    "  [red]► Frequency: [magenta]%.2f MHz\n\n"+
    "[yellow::b]MEMORY UTILIZATION\n"+
    "  [blue]► Usage: [red]%.2f%% (%s / %s)\n"+
    "  [blue]► Free: [green]%s | Cached: [blue]%s | Buffers: [cyan]%s\n"+
    "  [blue]► Total Used: [magenta]%s\n"+
    "  [blue]► Swap: [magenta]%.2f%% (%s / %s)\n"+
    "  [blue]► Swap Free: [green]%s\n\n"+
    "[yellow::b]STORAGE STATUS\n"+
    "  [green]► Usage: [red]%.2f%% (%s / %s)\n"+
    "  [green]► Free Space: [green]%s\n"+
    "  [green]► Inodes: Used: [blue]%s | Free: [green]%s | Total: [cyan]%d\n"+
    "  [green]► Total Used: [magenta]%s\n"+
    "  [green]► I/O Activity: Read: [magenta]%s/s | Write: [magenta]%s/s\n\n"+
    "[yellow::b]NETWORK METRICS\n"+
    "  [magenta]► Packets: Sent: [green]%d | Received: [blue]%d | Dropped: [red]%d\n",
    stats.cpuUsage, stats.cpuCount, stats.cpuIdle,
    stats.cpuTemp, stats.cpuUserTime, stats.cpuSystemTime,
    cpuFrequency.ModelName, cpuFrequency.Mhz,
    stats.memUsage, utils.FormatBytes(stats.memUsed), utils.FormatBytes(stats.memTotal),
    utils.FormatBytes(stats.memFree), utils.FormatBytes(stats.memCached), utils.FormatBytes(stats.memBuffers),
    utils.FormatBytes(stats.memTotalUsed),
    (float64(stats.swapUsed)/float64(stats.swapTotal))*100, utils.FormatBytes(stats.swapUsed), utils.FormatBytes(stats.swapTotal), utils.FormatBytes(stats.swapFree),
    stats.diskUsage, utils.FormatBytes(stats.diskUsed), utils.FormatBytes(stats.diskTotal),
    utils.FormatBytes(stats.diskFree),
    utils.FormatBytes(stats.diskInodesUsed), utils.FormatBytes(stats.diskInodesFree), stats.diskTotalInodes,
    utils.FormatBytes(stats.diskTotalUsed),
    utils.FormatBytes(stats.diskReadBytes), utils.FormatBytes(stats.diskWriteBytes),
    stats.netPacketsSent, stats.netPacketsRecv, stats.netPacketsDrop,
)

			output += "\n[yellow::b]NETWORK INTERFACES\n"
			for _, stat := range stats.netStats {
				output += fmt.Sprintf("  [cyan]► [green::b]%s\n    ↑ [blue]%s/s | ↓ [blue]%s/s\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

			output += "\n[yellow::b]DISK I/O STATISTICS\n"
			diskIOStatsMutex.RLock()
			if diskIOStats != nil {
				for device, stats := range diskIOStats {
					output += fmt.Sprintf("  [cyan]► [green::b]%s\n    ↑ [blue]%s/s | ↓ [blue]%s/s\n",
						device, utils.FormatBytes(stats.ReadBytes), utils.FormatBytes(stats.WriteBytes))
				}
			} else {
				output += "  [red]No disk I/O statistics available\n"
			}
			diskIOStatsMutex.RUnlock()
			output += "\n[yellow::b]╚═════════════════════════ Auto-refresh: 1s ═════════════════════════╝[::-]\n"
			app.QueueUpdateDraw(func() {
				textView.SetText(output)
			})			
		}
	}()

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(fmt.Sprintf("Failed to start UI: %v", err))
	}
}
