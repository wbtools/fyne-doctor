package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

type Dependency struct {
	Name        string
	Command     string
	Optional    bool
	Status      string
	Version     string
	Description string
	Platform    string // "all", "windows", "macos", "linux", "android", "ios"
	Category    string // "core", "mobile", "web", "performance", "compatibility"
	Severity    string // "critical", "warning", "info"
	IssueLink   string // Link to related GitHub issue
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "fyne-doctor",
		Short: "Fyne Environment Check Tool",
	}

	doctorCmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check Fyne development environment",
		Run:   runDoctor,
	}

	rootCmd.AddCommand(doctorCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runDoctor(cmd *cobra.Command, args []string) {
	fmt.Println("          Fyne Doctor")
	fmt.Println()

	// Fyne version (we simulate if fyne CLI not installed)
	fyneVersion := getFyneVersion()

	fmt.Printf("# Fyne\nVersion | %s\n\n", fyneVersion)

	// System info
	printSystemInfo()

	// Dependencies
	deps := getDependencies()

	for i := range deps {
		// Only check dependencies for current platform
		if deps[i].Platform == "all" || deps[i].Platform == runtime.GOOS {
			status, version := checkDependency(deps[i].Command)
			deps[i].Status = status
			deps[i].Version = version
		} else {
			deps[i].Status = "N/A"
			deps[i].Version = "Not applicable"
		}
	}

	printDependencyTable(deps)

	// Diagnosis
	success := true
	missingDeps := []string{}
	for _, d := range deps {
		if d.Platform == "all" || d.Platform == runtime.GOOS {
			if d.Status != "Installed" && !d.Optional {
				success = false
				missingDeps = append(missingDeps, d.Name)
			}
		}
	}

	fmt.Println("\n# Diagnosis")
	if success {
		fmt.Println(" SUCCESS  Your system is ready for Fyne development!")
	} else {
		fmt.Println(" FAILURE  Some required dependencies are missing or not properly installed.")
		fmt.Printf("Missing dependencies: %s\n", strings.Join(missingDeps, ", "))
		printInstallationTips()
	}

	// Check for common issues based on GitHub Issues
	checkCommonIssues(deps)
}

// getFyneVersion simulates fyne version if CLI not installed
func getFyneVersion() string {
	if _, err := exec.LookPath("fyne"); err != nil {
		return "Not installed"
	}
	out, err := exec.Command("fyne", "version").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}

// getDependencies returns platform-specific dependencies based on Fyne documentation and GitHub Issues
func getDependencies() []Dependency {
	baseDeps := []Dependency{
		{
			Name:        "Go",
			Command:     "go version",
			Optional:    false,
			Description: "Go programming language (minimum version 1.12)",
			Platform:    "all",
			Category:    "core",
			Severity:    "critical",
		},
		{
			Name:        "Fyne CLI",
			Command:     "fyne version",
			Optional:    false,
			Description: "Fyne command line tools",
			Platform:    "all",
			Category:    "core",
			Severity:    "critical",
		},
		{
			Name:        "Fyne-cross",
			Command:     "fyne-cross --version",
			Optional:    true,
			Description: "Cross-platform build tool for Fyne",
			Platform:    "all",
			Category:    "core",
			Severity:    "info",
		},
	}

	// Platform-specific dependencies
	switch runtime.GOOS {
	case "windows":
		baseDeps = append(baseDeps, []Dependency{
			{
				Name:        "C Compiler (MSYS2)",
				Command:     "gcc --version",
				Optional:    false,
				Description: "C compiler for Windows (MSYS2/MingW-w64 recommended)",
				Platform:    "windows",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "MSYS2",
				Command:     "pacman --version",
				Optional:    false,
				Description: "MSYS2 package manager",
				Platform:    "windows",
				Category:    "core",
				Severity:    "critical",
			},
		}...)
	case "darwin": // macOS
		baseDeps = append(baseDeps, []Dependency{
			{
				Name:        "Xcode CLI Tools",
				Command:     "xcode-select -p",
				Optional:    false,
				Description: "Xcode command line tools",
				Platform:    "macos",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "C Compiler",
				Command:     "clang --version",
				Optional:    false,
				Description: "Clang compiler (comes with Xcode)",
				Platform:    "macos",
				Category:    "core",
				Severity:    "critical",
			},
		}...)
	case "linux":
		baseDeps = append(baseDeps, []Dependency{
			{
				Name:        "C Compiler",
				Command:     "gcc --version",
				Optional:    false,
				Description: "GCC compiler",
				Platform:    "linux",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "pkg-config",
				Command:     "pkg-config --version",
				Optional:    false,
				Description: "Package configuration tool",
				Platform:    "linux",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "Mesa GL",
				Command:     "pkg-config --exists gl && echo 'Found'",
				Optional:    false,
				Description: "Mesa OpenGL library",
				Platform:    "linux",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "X11 Development",
				Command:     "pkg-config --exists x11 && echo 'Found'",
				Optional:    false,
				Description: "X11 development libraries",
				Platform:    "linux",
				Category:    "core",
				Severity:    "critical",
			},
			{
				Name:        "Wayland Support",
				Command:     "pkg-config --exists wayland-client && echo 'Found'",
				Optional:    true,
				Description: "Wayland client library (for Wayland support)",
				Platform:    "linux",
				Category:    "compatibility",
				Severity:    "warning",
				IssueLink:   "https://github.com/fyne-io/fyne/issues/5908",
			},
		}...)
	}

	// Mobile development dependencies (based on GitHub Issues)
	baseDeps = append(baseDeps, []Dependency{
		{
			Name:        "Android SDK",
			Command:     "echo $ANDROID_HOME",
			Optional:    true,
			Description: "Android SDK for mobile development",
			Platform:    "all",
			Category:    "mobile",
			Severity:    "info",
		},
		{
			Name:        "Android NDK",
			Command:     "echo $ANDROID_NDK_HOME",
			Optional:    true,
			Description: "Android NDK for mobile development",
			Platform:    "all",
			Category:    "mobile",
			Severity:    "info",
		},
		{
			Name:        "Android Studio",
			Command:     "which studio || which android-studio",
			Optional:    true,
			Description: "Android Studio IDE (recommended for mobile dev)",
			Platform:    "all",
			Category:    "mobile",
			Severity:    "info",
		},
		{
			Name:        "iOS Simulator",
			Command:     "xcrun simctl list devices",
			Optional:    true,
			Description: "iOS Simulator for iOS development",
			Platform:    "darwin",
			Category:    "mobile",
			Severity:    "info",
		},
	}...)

	// Web development dependencies
	baseDeps = append(baseDeps, []Dependency{
		{
			Name:        "WebAssembly Support",
			Command:     "go version | grep -q 'go1.16' && echo 'Supported' || echo 'Requires Go 1.16+'",
			Optional:    true,
			Description: "WebAssembly support for web builds",
			Platform:    "all",
			Category:    "web",
			Severity:    "info",
		},
		{
			Name:        "Node.js",
			Command:     "node --version",
			Optional:    true,
			Description: "Node.js for web development tools",
			Platform:    "all",
			Category:    "web",
			Severity:    "info",
		},
	}...)

	// Performance and compatibility checks
	baseDeps = append(baseDeps, []Dependency{
		{
			Name:        "GPU Acceleration",
			Command:     "glxinfo | grep 'direct rendering' || echo 'Not available'",
			Optional:    true,
			Description: "GPU acceleration support (Linux)",
			Platform:    "linux",
			Category:    "performance",
			Severity:    "warning",
		},
		{
			Name:        "Display Server",
			Command:     "echo $XDG_SESSION_TYPE",
			Optional:    true,
			Description: "Current display server (X11/Wayland)",
			Platform:    "linux",
			Category:    "compatibility",
			Severity:    "info",
		},
	}...)

	return baseDeps
}

func checkDependency(cmd string) (status string, version string) {
	// Handle environment variable checks
	if strings.HasPrefix(cmd, "echo $") {
		envVar := strings.TrimPrefix(cmd, "echo $")
		value := os.Getenv(envVar)
		if value == "" {
			return "Missing", ""
		}
		return "Installed", value
	}

	// Handle pkg-config checks
	if strings.Contains(cmd, "pkg-config --exists") {
		c := exec.Command("sh", "-c", cmd)
		if runtime.GOOS == "windows" {
			c = exec.Command("cmd", "/C", cmd)
		}
		err := c.Run()
		if err != nil {
			return "Missing", ""
		}
		return "Installed", "Found"
	}

	// Handle which command checks
	if strings.HasPrefix(cmd, "which ") {
		c := exec.Command("sh", "-c", cmd)
		if runtime.GOOS == "windows" {
			c = exec.Command("cmd", "/C", cmd)
		}
		err := c.Run()
		if err != nil {
			return "Missing", ""
		}
		return "Installed", "Found"
	}

	// Handle complex command chains
	if strings.Contains(cmd, "||") || strings.Contains(cmd, "&&") {
		c := exec.Command("sh", "-c", cmd)
		if runtime.GOOS == "windows" {
			c = exec.Command("cmd", "/C", cmd)
		}
		out, err := c.CombinedOutput()
		if err != nil {
			return "Error", ""
		}
		result := strings.TrimSpace(string(out))
		if result == "" {
			return "Missing", ""
		}
		return "Installed", result
	}

	// Regular command checks
	firstCmd := strings.Fields(cmd)[0]
	if _, err := exec.LookPath(firstCmd); err != nil {
		return "Missing", ""
	}
	c := exec.Command("sh", "-c", cmd)
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", cmd)
	}
	out, err := c.CombinedOutput()
	if err != nil {
		return "Error", ""
	}
	return "Installed", strings.TrimSpace(string(out))
}

func printDependencyTable(deps []Dependency) {
	// Group dependencies by category
	categories := make(map[string][]Dependency)
	for _, d := range deps {
		if d.Platform == "all" || d.Platform == runtime.GOOS {
			categories[d.Category] = append(categories[d.Category], d)
		}
	}

	// Define category order and display names
	categoryOrder := []string{"core", "mobile", "web", "performance", "compatibility"}
	categoryNames := map[string]string{
		"core":          "Core Dependencies",
		"mobile":        "Mobile Development",
		"web":           "Web Development",
		"performance":   "Performance",
		"compatibility": "Compatibility",
	}

	for _, cat := range categoryOrder {
		if deps, exists := categories[cat]; exists && len(deps) > 0 {
			fmt.Printf("\n## %s\n", categoryNames[cat])
			fmt.Println("┌─────────────────────────────┬────────────┬────────────┬─────────────┬─────────────────────────────────────┐")
			fmt.Println("| Dependency                  | Optional   | Status     | Version     | Description                         |")
			fmt.Println("├─────────────────────────────┼────────────┼────────────┼─────────────┼─────────────────────────────────────┤")

			for _, d := range deps {
				opt := ""
				if d.Optional {
					opt = "*"
				}
				// Truncate long descriptions
				desc := d.Description
				if len(desc) > 35 {
					desc = desc[:32] + "..."
				}
				fmt.Printf("| %-27s | %-10s | %-10s | %-11s | %-35s |\n",
					d.Name, opt, d.Status, d.Version, desc)
			}
			fmt.Println("└─────────────────────────────┴────────────┴────────────┴─────────────┴─────────────────────────────────────┘")
		}
	}
}

// printSystemInfo prints OS, Go version, CPU, Memory
func printSystemInfo() {
	fmt.Println("# System")
	cpuInfo, _ := cpu.Info()
	memInfo, _ := mem.VirtualMemory()

	fmt.Println("┌───────────────────────────┐")
	fmt.Printf("| OS           | %s\n", runtime.GOOS)
	fmt.Printf("| Architecture | %s\n", runtime.GOARCH)
	if len(cpuInfo) > 0 {
		fmt.Printf("| CPU          | %s @ %.0fMHz\n", cpuInfo[0].ModelName, cpuInfo[0].Mhz)
	}
	fmt.Printf("| Memory       | %.2f GB\n", float64(memInfo.Total)/1024/1024/1024)
	fmt.Println("└───────────────────────────┘")
	fmt.Println()
}

// printInstallationTips provides platform-specific installation instructions
func printInstallationTips() {
	fmt.Println("\n# Installation Tips")
	fmt.Println("Based on your platform, here are the recommended installation steps:")
	fmt.Println()

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows:")
		fmt.Println("1. Install Go from https://golang.org/dl/")
		fmt.Println("2. Install MSYS2 from https://www.msys2.org/")
		fmt.Println("3. Open MSYS2 MinGW 64-bit terminal and run:")
		fmt.Println("   pacman -Syu")
		fmt.Println("   pacman -S git mingw-w64-x86_64-toolchain")
		fmt.Println("4. Add C:\\msys64\\mingw64\\bin to your PATH")
		fmt.Println("5. Install Fyne CLI: go install fyne.io/fyne/v2/cmd/fyne@latest")

	case "darwin": // macOS
		fmt.Println("macOS:")
		fmt.Println("1. Install Go from https://golang.org/dl/")
		fmt.Println("2. Install Xcode from Mac App Store")
		fmt.Println("3. Install Xcode CLI tools: xcode-select --install")
		fmt.Println("4. Install Fyne CLI: go install fyne.io/fyne/v2/cmd/fyne@latest")

	case "linux":
		fmt.Println("Linux:")
		fmt.Println("For Debian/Ubuntu:")
		fmt.Println("  sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev pkg-config")
		fmt.Println()
		fmt.Println("For Fedora:")
		fmt.Println("  sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel")
		fmt.Println()
		fmt.Println("For Arch Linux:")
		fmt.Println("  sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi")
		fmt.Println()
		fmt.Println("Then install Fyne CLI: go install fyne.io/fyne/v2/cmd/fyne@latest")
	}

	fmt.Println()
	fmt.Println("For mobile development:")
	fmt.Println("- Android: Install Android Studio and NDK")
	fmt.Println("- iOS: Requires macOS with Xcode and Apple Developer account")
	fmt.Println()
	fmt.Println("For cross-platform builds:")
	fmt.Println("  go install github.com/fyne-io/fyne-cross@latest")
}

// checkCommonIssues checks for common problems based on GitHub Issues
func checkCommonIssues(deps []Dependency) {
	fmt.Println("\n# Common Issues Check")
	issues := []string{}
	warnings := []string{}

	// Check for Wayland issues
	displayServer := os.Getenv("XDG_SESSION_TYPE")
	if displayServer == "wayland" {
		waylandSupport := false
		for _, d := range deps {
			if d.Name == "Wayland Support" && d.Status == "Installed" {
				waylandSupport = true
				break
			}
		}
		if !waylandSupport {
			issues = append(issues, "Running on Wayland but Wayland support may be incomplete")
			fmt.Println("⚠️  WARNING: You're running on Wayland. Some Fyne apps may have issues.")
			fmt.Println("   Related issue: https://github.com/fyne-io/fyne/issues/5908")
		}
	}

	// Check for mobile development issues
	androidSDK := false
	androidNDK := false
	for _, d := range deps {
		if d.Name == "Android SDK" && d.Status == "Installed" {
			androidSDK = true
		}
		if d.Name == "Android NDK" && d.Status == "Installed" {
			androidNDK = true
		}
	}
	if androidSDK && !androidNDK {
		warnings = append(warnings, "Android SDK found but NDK missing - mobile builds may fail")
	}

	// Check for performance issues
	gpuAccel := false
	for _, d := range deps {
		if d.Name == "GPU Acceleration" && d.Status == "Installed" {
			gpuAccel = true
			break
		}
	}
	if !gpuAccel && runtime.GOOS == "linux" {
		warnings = append(warnings, "GPU acceleration not available - performance may be reduced")
	}

	// Check for web development issues
	goVersion := ""
	for _, d := range deps {
		if d.Name == "Go" && d.Status == "Installed" {
			goVersion = d.Version
			break
		}
	}
	if goVersion != "" && !strings.Contains(goVersion, "go1.16") {
		warnings = append(warnings, "Go version < 1.16 - WebAssembly builds not supported")
	}

	// Display results
	if len(issues) == 0 && len(warnings) == 0 {
		fmt.Println("✅ No common issues detected!")
	} else {
		if len(issues) > 0 {
			fmt.Println("\n🚨 Issues found:")
			for _, issue := range issues {
				fmt.Printf("   • %s\n", issue)
			}
		}
		if len(warnings) > 0 {
			fmt.Println("\n⚠️  Warnings:")
			for _, warning := range warnings {
				fmt.Printf("   • %s\n", warning)
			}
		}
	}

	// Display GitHub Issues summary
	fmt.Println("\n# GitHub Issues Summary")
	fmt.Println("Based on recent Fyne GitHub Issues, common problems include:")
	fmt.Println("• Mobile web builds: Paste functionality issues (#5916)")
	fmt.Println("• Android: NewMultiLineEntry scrolling problems (#5915)")
	fmt.Println("• Mobile: Grid container performance issues (#5914)")
	fmt.Println("• Wayland: App crashes on some systems (#5908)")
	fmt.Println("• X11: Display wake-up crashes (#5899)")
	fmt.Println("• Windows: UI position issues after minimize/restore (#5898)")
	fmt.Println("\nFor more details, visit: https://github.com/fyne-io/fyne/issues")
}
