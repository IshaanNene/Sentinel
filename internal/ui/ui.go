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

func StartUI() {
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
				"\n[::b]╔══════════════════════ SYSTEM RESOURCES ════════════════════╗[::-]\n\n"+
					"[yellow::b]CPU INFORMATION[white]\n"+
					"  [red]►[white] Usage: [red]%.2f%%[white] | Cores: [green]%d[white] | Idle: [blue]%.2f%%[white]\n"+
					"  [red]►[white] Temperature: [red]%.2f°C[white] | User Time: [cyan]%.2f[white] | System Time: [cyan]%.2f[white]\n"+
					"  [red]►[white] Model: [magenta]%s[white]\n"+
					"  [red]►[white] Frequency: [magenta]%.2f MHz[white]\n\n"+
					"[yellow::b]MEMORY UTILIZATION[white]\n"+
					"  [blue]►[white] Usage: [red]%.2f%%[white] (%s / %s)\n"+
					"  [blue]►[white] Free: [green]%s[white] | Cached: [blue]%s[white] | Buffers: [cyan]%s[white]\n"+
					"  [blue]►[white] Total Used: [magenta]%s[white]\n"+
					"  [blue]►[white] Swap: [magenta]%.2f%%[white] (%s / %s)\n"+
					"  [blue]►[white] Swap Free: [green]%s[white]\n\n"+
					"[yellow::b]STORAGE STATUS[white]\n"+
					"  [green]►[white] Usage: [red]%.2f%%[white] (%s / %s)\n"+
					"  [green]►[white] Free Space: [green]%s[white]\n"+
					"  [green]►[white] Inodes: Used: [blue]%s[white] | Free: [green]%s[white] | Total: [cyan]%d[white]\n"+
					"  [green]►[white] Total Used: [magenta]%s[white]\n"+
					"  [green]►[white] I/O Activity: Read: [magenta]%s/s[white] | Write: [magenta]%s/s[white]\n\n"+
					"[yellow::b]NETWORK METRICS[white]\n"+
					"  [magenta]►[white] Packets: Sent: [green]%d[white] | Received: [blue]%d[white] | Dropped: [red]%d[white]\n",
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

			output += "\n[yellow::b]NETWORK INTERFACES[white]\n"
			for _, stat := range stats.netStats {
				output += fmt.Sprintf("  [cyan]►[white] [green::b]%s[white]\n    ↑ [blue]%s/s[white] | ↓ [blue]%s/s[white]\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

			output += "\n[yellow::b]DISK I/O STATISTICS[white]\n"
			diskIOStatsMutex.RLock()
			if diskIOStats != nil {
				for device, stats := range diskIOStats {
					output += fmt.Sprintf("  [cyan]►[white] [green::b]%s[white]\n    ↑ [blue]%s/s[white] | ↓ [blue]%s/s[white]\n",
						device, utils.FormatBytes(stats.ReadBytes), utils.FormatBytes(stats.WriteBytes))
				}
			} else {
				output += "  [red]No disk I/O statistics available[white]\n"
			}
			diskIOStatsMutex.RUnlock()
			output += "\n[::b]╚═════════════════════════ [white]Auto-refresh: 1s[white] ═════════════════════════╝[::-]\n"
			app.QueueUpdateDraw(func() {
				textView.SetText(output)
				textView.ScrollToBeginning()
			})			
		}
	}()

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(fmt.Sprintf("Failed to start UI: %v", err))
	}
}
