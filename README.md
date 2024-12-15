#  🛡️**Sentinel**🛡 

**Sentinel** is a powerful real-time system monitoring tool built with Go. It provides a terminal-based user interface (TUI) to track crucial system metrics such as CPU usage, memory consumption, disk space, and network activity. **Sentinel** delivers live updates and detailed insights into your system’s health, making it the ideal tool for developers, sysadmins, and anyone passionate about tracking their system performance in real-time. ⚡

## 🌟 Features

- **🔄 Real-Time Monitoring**: Stay up-to-date with live stats on CPU, memory, disk, and network usage.
- **🖥️ Terminal User Interface (TUI)**: Sleek, dynamic, and intuitive UI for easy navigation and fast insights.
- **🌍 Cross-Platform**: Runs seamlessly on Linux, macOS, and Windows (via WSL).
- **⚡ Lightweight & Fast**: Built with Go for minimal resource usage while providing accurate, real-time metrics.
- **🔧 Customizable**: Easily extend and modify for additional system metrics or new features.

## 🔧 Technologies Used

- **Go**: The main language powering the application, known for its speed and efficiency.
- **tview**: A Go package used to create rich terminal-based UIs.
- **gopsutil**: A library for accessing system metrics like CPU, memory, and disk stats.
- **Terminal UI**: A dynamic and responsive interface built to give you real-time system insights.

## 🚀 Installation

### Prerequisites

- Go 1.20 or later installed on your system. If you don’t have Go installed yet, check out the [Go Installation Guide](https://go.dev/doc/install). 👨‍💻

### Steps to Install

1. **Clone the repository**:
   ```bash
   git clone https://github.com/<your-username>/sentinel.git
   cd sentinel
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```
   
   Once the application runs, you’ll be presented with a terminal-based UI displaying live system metrics. 🎉

## ⚙️ Usage

- **⚡ CPU**: Real-time CPU usage and load stats, displayed dynamically.
- **💾 Memory**: See how much memory is being used and what’s available.
- **📂 Disk**: Keep track of disk space usage and free space.
- **🌐 Network**: Monitor network activity, including bytes sent and received.

## 📸 Screenshots
<img width="256" alt="Screenshot 2024-12-15 at 2 25 08 PM" src="https://github.com/user-attachments/assets/40935754-4b0e-44e6-9321-bf3279208368" />


## 🛠️ Future Enhancements

- **🔔 Alerting System**: Get notified when system metrics exceed predefined thresholds (e.g., CPU usage spikes). 
- **📊 Export Feature**: Export system stats to CSV or JSON for analysis.
- **🔌 Support for More Metrics**: Add support for additional metrics like GPU usage, system temperature, etc.
- **🐋 Docker Support**: Make deployment even easier by containerizing **Sentinel** for use in Docker environments.

## 🤝 Contributing

I welcome contributions to **Sentinel**! If you want to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes (`git checkout -b feature-branch`).
3. Implement your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to your fork (`git push origin feature-branch`).
5. Open a pull request.

## 📜 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.
