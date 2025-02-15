package main

import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
        "sort"
        "strings"
)

// Banner function
func printBanner() {
        fmt.Println(`


██     ██  █████  ██      ██      ███    ███  █████  ██████
██     ██ ██   ██ ██      ██      ████  ████ ██   ██ ██   ██
██  █  ██ ███████ ██      ██      ██ ████ ██ ███████ ██████
██ ███ ██ ██   ██ ██      ██      ██  ██  ██ ██   ██ ██
 ███ ███  ██   ██ ███████ ███████ ██      ██ ██   ██ ██


-----------------------------------------------------------
       WallMap - Created by Evil-Twinz (Srilakivarma)
-----------------------------------------------------------
`)
}

// Expand CIDR notation to individual IPs
func expandCIDR(cidr string) ([]string, error) {
        var ips []string
        _, ipNet, err := net.ParseCIDR(cidr)
        if err != nil {
                return nil, err
        }
        for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
                ips = append(ips, ip.String())
        }
        return ips, nil
}

// Increment an IP address
func inc(ip net.IP) {
        for j := len(ip) - 1; j >= 0; j-- {
                ip[j]++
                if ip[j] > 0 {
                        break
                }
        }
}

// Expand IP range notation (e.g., 192.168.10.11-192.168.10.55)
func expandIPRange(rangeStr string) ([]string, error) {
        parts := strings.Split(rangeStr, "-")
        if len(parts) != 2 {
                return nil, fmt.Errorf("invalid range format: %s", rangeStr)
        }
        ipStart := net.ParseIP(parts[0])
        ipEnd := net.ParseIP(parts[1])
        if ipStart == nil || ipEnd == nil {
                return nil, fmt.Errorf("invalid IP range: %s", rangeStr)
        }

        var ips []string
        for ip := ipStart; !ip.Equal(ipEnd); inc(ip) {
                ips = append(ips, ip.String())
        }
        ips = append(ips, ipEnd.String()) // Include last IP
        return ips, nil
}

// Resolve hostnames to IP addresses
func resolveHostname(hostname string) ([]string, error) {
        ips, err := net.LookupHost(hostname)
        if err != nil {
                return nil, err
        }
        return ips, nil
}

func main() {
        // Define flags
        filePath := flag.String("l", "", "Path to the file containing IPs/subnets/hostnames")
        ipv4Only := flag.Bool("ipv4-only", false, "Filter out IPv6 addresses")
        silent := flag.Bool("silent", false, "Run in silent mode (no banner)")
        flag.Parse()

        // Print the banner unless silent mode is enabled
        if !*silent {
                printBanner()
        }

        // Map to store unique IPs
        uniqueIPs := make(map[string]bool)

        // Function to process each line (from file or stdin)
        processLine := func(line string) {
                line = strings.Split(line, "#")[0] // Remove comments
                line = strings.TrimSpace(line)      // Trim spaces after removing comments
                if line == "" {
                        return
                }

                // Handle CIDR notation
                if strings.Contains(line, "/") {
                        ips, err := expandCIDR(line)
                        if err == nil {
                                for _, ip := range ips {
                                        uniqueIPs[ip] = true
                                }
                        } else {
                                fmt.Fprintf(os.Stderr, "Invalid CIDR: %s\n", line)
                        }
                        return
                }

                // Handle IP range notation
                if strings.Contains(line, "-") {
                        ips, err := expandIPRange(line)
                        if err == nil {
                                for _, ip := range ips {
                                        uniqueIPs[ip] = true
                                }
                        } else {
                                fmt.Fprintf(os.Stderr, "Invalid IP range: %s\n", line)
                        }
                        return
                }

                // Handle hostnames (resolve to IPs)
                if net.ParseIP(line) == nil {
                        ips, err := resolveHostname(line)
                        if err == nil {
                                for _, ip := range ips {
                                        uniqueIPs[ip] = true
                                }
                        } else {
                                fmt.Fprintf(os.Stderr, "Failed to resolve hostname: %s\n", line)
                        }
                        return
                }

                // Otherwise, treat it as a normal IP
                uniqueIPs[line] = true
        }

        // Read from file if -l is provided
        if *filePath != "" {
                file, err := os.Open(*filePath)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
                        os.Exit(1)
                }
                defer file.Close()

                scanner := bufio.NewScanner(file)
                for scanner.Scan() {
                        processLine(scanner.Text())
                }
                if err := scanner.Err(); err != nil {
                        fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
                        os.Exit(1)
                }
        } else {
                // Read from stdin (for one-liner usage)
                scanner := bufio.NewScanner(os.Stdin)
                for scanner.Scan() {
                        processLine(scanner.Text())
                }
                if err := scanner.Err(); err != nil {
                        fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
                        os.Exit(1)
                }
        }

        // Convert map to sorted slice
        var sortedIPs []string
        for ip := range uniqueIPs {
                sortedIPs = append(sortedIPs, ip)
        }
        sort.Strings(sortedIPs)

        // Print unique sorted IPs, filtering IPv6 if needed
        for _, ip := range sortedIPs {
                if *ipv4Only && strings.Contains(ip, ":") {
                        continue // Skip IPv6 addresses
                }
                fmt.Println(ip)
        }
}
