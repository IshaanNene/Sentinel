package ui

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"goscope/internal/monitor"
	"goscope/pkg/utils"
)

func StartUI() {
	app := tview.NewApplication()
	// Create TextView and store in variable explicitly typed as *tview.TextView
	var textView *tview.TextView = tview.NewTextView()
	textView.
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetBorder(true).
		SetTitle(" 🖥️  SENTINEL SYSTEM MONITOR ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(tview.Styles.BorderColor)

	diskIOStats, _ := monitor.GetDiskIOStats()

	// Launch a goroutine for dynamic updates
	go func() {
		for {
			cpuUsage, _ := monitor.GetCPUUsage()
			cpuCount, _ := monitor.GetCPUCount()
			memUsage, memUsed, _ := monitor.GetMemoryUsage()
			memTotal, _ := monitor.GetMemoryTotal()
			diskUsage, diskUsed, _ := monitor.GetDiskUsage()
			diskTotal, _ := monitor.GetDiskTotal()
			netStats, _ := monitor.GetNetworkStats()
			swapTotal, _ := monitor.GetMemorySwapTotal()
			swapUsed, _ := monitor.GetMemorySwapUsed()
			diskFree, _ := monitor.GetDiskFree()
			memFree, _ := monitor.GetMemoryFree()
			cpuIdle, _ := monitor.GetCPUIdle()
			diskInodesUsed, diskInodesFree, _ := monitor.GetDiskInodes()
			cpuFrequencyInfo, _ := monitor.GetCPUFrequency()
			cpuFrequency := cpuFrequencyInfo[len(cpuFrequencyInfo)-1]
			cpuModelName := cpuFrequency.ModelName
			cpuMhz := cpuFrequency.Mhz
			diskReadBytes, _ := monitor.GetDiskReadBytes()
			diskWriteBytes, _ := monitor.GetDiskWriteBytes()
			cpuTemp, _ := monitor.GetCPUTemperature()
			memCached, _ := monitor.GetMemoryCached()
			memBuffers, _ := monitor.GetMemoryBuffers()
			netPacketsSent, _ := monitor.GetNetworkPacketsSent()
			netPacketsReceived, _ := monitor.GetNetworkPacketsReceived()
			netPacketsDropped, _ := monitor.GetNetworkPacketsDropped()
			cpuUserTime, _ := monitor.GetCPUUserTime()
			cpuSystemTime, _ := monitor.GetCPUSystemTime()
			swapFree, _ := monitor.GetMemorySwapFree()
			diskTotalInodes, _ := monitor.GetDiskTotalInodes()
			memTotalUsed, _ := monitor.GetMemoryTotalUsed()
			diskTotalUsed, _ := monitor.GetDiskTotalUsed()

			output := fmt.Sprintf(
				"\n[::b]═══════════════════ SYSTEM RESOURCES ═══════════════════[::-]\n\n"+
					"[yellow::b]CPU[white]\n"+
					"  Usage: [red]%.2f%%[white] | Cores: [green]%d[white] | Idle: [blue]%.2f%%[white]\n"+
					"  Temperature: [red]%.2f°C[white] | User Time: [cyan]%.2f[white] | System Time: [cyan]%.2f[white]\n"+
					"  Model: [magenta]%s[white] @ [magenta]%.2f MHz[white]\n\n"+
					"[yellow::b]MEMORY[white]\n"+
					"  Usage: [red]%.2f%%[white] (%s / %s)\n"+
					"  Free: [green]%s[white] | Cached: [blue]%s[white] | Buffers: [cyan]%s[white]\n"+
					"  Total Used: [magenta]%s[white]\n"+
					"  Swap: [magenta]%.2f%%[white] (%s / %s) | Free: [green]%s[white]\n\n"+
					"[yellow::b]STORAGE[white]\n"+
					"  Usage: [red]%.2f%%[white] (%s / %s)\n"+
					"  Free: [green]%s[white]\n"+
					"  Inodes: Used: [blue]%s[white], Free: [green]%s[white], Total: [cyan]%d[white]\n"+
					"  Total Used: [magenta]%s[white]\n"+
					"  I/O: Read: [magenta]%s[white] | Write: [magenta]%s[white]\n\n"+
					"[yellow::b]NETWORK STATISTICS[white]\n"+
					"  Packets: Sent: [green]%d[white] | Received: [blue]%d[white] | Dropped: [red]%d[white]\n",
				cpuUsage, cpuCount, cpuIdle,
				cpuTemp, cpuUserTime, cpuSystemTime,
				cpuModelName, cpuMhz,
				memUsage, utils.FormatBytes(memUsed), utils.FormatBytes(memTotal),
				utils.FormatBytes(memFree), utils.FormatBytes(memCached), utils.FormatBytes(memBuffers),
				utils.FormatBytes(memTotalUsed),
				(float64(swapUsed)/float64(swapTotal))*100, utils.FormatBytes(swapUsed), utils.FormatBytes(swapTotal), utils.FormatBytes(swapFree),
				diskUsage, utils.FormatBytes(diskUsed), utils.FormatBytes(diskTotal),
				utils.FormatBytes(diskFree),
				utils.FormatBytes(diskInodesUsed), utils.FormatBytes(diskInodesFree), diskTotalInodes,
				utils.FormatBytes(diskTotalUsed),
				utils.FormatBytes(diskReadBytes), utils.FormatBytes(diskWriteBytes),
				uint64(netPacketsSent), uint64(netPacketsReceived), uint64(netPacketsDropped),
			)

			output += "\n[yellow::b]NETWORK INTERFACES[white]\n"
			for _, stat := range netStats {
				output += fmt.Sprintf("  [green::b]%s[white]\n    ↑ [blue]%s[white] | ↓ [blue]%s[white]\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

			output += "\n[yellow::b]DISK I/O STATISTICS[white]\n"
			if diskIOStats != nil {
				for device, stats := range diskIOStats {
					output += fmt.Sprintf("  [green::b]%s[white]\n    ↑ [blue]%s[white] | ↓ [blue]%s[white]\n",
						device, utils.FormatBytes(stats.ReadBytes), utils.FormatBytes(stats.WriteBytes))
				}
			} else {
				output += "  [red]No disk I/O statistics available[white]\n"
			}

			output += "\n[::b]══════════════════════════════════[white]Refresh every second...[white][::-]\n"

			app.QueueUpdateDraw(func() {
				textView.SetText(output) 
			})
			time.Sleep(1 * time.Second)
		}
	}()
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
