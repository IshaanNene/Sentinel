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
	textView := tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft)

	diskIOStats, _ := monitor.GetDiskIOStats()

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
				"[cyan]SENTINEL - SYSTEM MONITOR[white]\n\n"+
				"[yellow]CPU USAGE:[white] %.2f%% (CORES: %d)\n"+
				"[yellow]MEMORY USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK TOTAL:[white] %s\n"+
				"[yellow]SWAP USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK FREE:[white] %s\n"+
				"[yellow]MEMORY FREE:[white] %s\n"+
				"[yellow]CPU IDLE:[white] %.2f%%\n"+
				"[yellow]DISK INODES:[white] Used: %s, Free: %s\n"+
				"[yellow]NETWORK STATS:[white]\n"+
				"[yellow]CPU TEMP:[white] %.2f°C\n"+
				"[yellow]MEMORY CACHED:[white] %s\n"+
				"[yellow]MEMORY BUFFERS:[white] %s\n"+
				"[yellow]PACKETS SENT:[white] %d\n"+
				"[yellow]PACKETS RECEIVED:[white] %d\n"+
				"[yellow]PACKETS DROPPED:[white] %d\n"+
				"[yellow]CPU USER TIME:[white] %.2f\n"+
				"[yellow]CPU SYSTEM TIME:[white] %.2f\n"+
				"[yellow]SWAP FREE:[white] %s\n"+
				"[yellow]DISK TOTAL INODES:[white] %d\n"+
				"[yellow]MEMORY TOTAL USED:[white] %s\n"+
				"[yellow]DISK TOTAL USED:[white] %s\n",
				cpuUsage, cpuCount, memUsage, utils.FormatBytes(memUsed), utils.FormatBytes(memTotal),
				diskUsage, utils.FormatBytes(diskUsed), utils.FormatBytes(diskTotal),
				utils.FormatBytes(diskTotal),
				(float64(swapUsed)/float64(swapTotal))*100, utils.FormatBytes(swapUsed), utils.FormatBytes(swapTotal),
				utils.FormatBytes(diskFree), utils.FormatBytes(memFree),
				cpuIdle, utils.FormatBytes(diskInodesUsed), utils.FormatBytes(diskInodesFree),
				cpuTemp, utils.FormatBytes(memCached), utils.FormatBytes(memBuffers),
				netPacketsSent, netPacketsReceived, netPacketsDropped,
				cpuUserTime, cpuSystemTime,
				utils.FormatBytes(swapFree), diskTotalInodes,
				utils.FormatBytes(memTotalUsed), utils.FormatBytes(diskTotalUsed),
			)

			for _, stat := range netStats {
				output += fmt.Sprintf("  [green]INTERFACE:[white] %s | [blue]SENT:[white] %s | [blue]RECEIVED:[white] %s\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

			output += fmt.Sprintf(
				"[yellow]DISK READ:[white] %s\n"+
				"[yellow]DISK WRITE:[white] %s\n"+
				"[yellow]CPU FREQUENCY:[white] %s (%.2f MHz)\n\n",
				utils.FormatBytes(diskReadBytes), utils.FormatBytes(diskWriteBytes),
				cpuModelName, cpuMhz,
			)
			output += "[yellow]DISK IO STATS:[white]\n"
			if diskIOStats != nil {
				for device, stats := range diskIOStats {
					output += fmt.Sprintf("  [green]DEVICE:[white] %s | [blue]READ:[white] %s | [blue]WRITE:[white] %s\n",
						device, utils.FormatBytes(stats.ReadBytes), utils.FormatBytes(stats.WriteBytes))
				}
			} else {
				output += "[red]No disk IO stats available.[white]\n"
			}

			output += "\n[red]UPDATING EVERY SECOND... STAY ALERT!\n"

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
