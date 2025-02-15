<h1 align="center">WallMap</h1>

<p align="center">
  <img src="https://github.com/user-attachments/assets/e875995d-5574-4e34-83ca-8d5c6a5de601" alt="Wallmap" width="650" height="600">  
</p>




WallMap is a powerful and efficient IP deduplication and expansion tool written in Go. It allows users to handle CIDR notations, IP ranges, and hostnames to generate unique, sorted lists of IP addresses. This tool is designed for penetration testers, bug bounty hunters, and network engineers.

<h1>🚀 Features</h1>

1.CIDR Expansion: Convert subnet ranges to individual IPs.

2.IP Range Expansion: Handle IP ranges (e.g., 192.168.0.1-192.168.0.255).

3.Hostname Resolution: Resolve hostnames to their IP addresses.

4.IPv4/IPv6 Support: Option to filter IPv6 addresses.

5.Silent Mode: Suppress banners for scripting.

6.File and Stdin Support: Process inputs from files or piped commands.

<h1>📥 Installation</h1>

Make sure you have Go installed. To install WallMap:

    go install github.com/Evil-twinz/Wallmap@latest
           
This will place the wallmap executable in your $GOPATH/bin

<h1>💻 Usage</h1>

    wallmap -l ips.txt > output.txt
          

<h1>Command-Line Flags:</h1>

<table>
  <tr>
    <th>Flag</th><th>Type</th><th>Description</th>
  </tr>
  <tr>
    <td>-l</td><td>string</td><td>Path to file containing IPs, ranges, or hostnames.</td>
  </tr>
  <tr>
    <td>-ipv4-only</td><td>bool</td><td>Filter out IPv6 addresses.</td>
  </tr>
  <tr>
    <td>-silent</td><td>bool</td><td>Run without displaying the banner.</td>
  </tr>
</table>


Examples:
1.From File:

    wallmap -l targets.txt -ipv4-only > ips.txt

2.From Stdin:

    cat targets.txt | wallmap -ipv4-only

3.With Silent Mode:

    wallmap -l targets.txt -silent


<h1>📝 Contributing</h1>

Pull requests and contributions are welcome. Please fork the repository and submit your PRs.


<h1>🏷️ Author</h1>

Created with ❤️ by Evil-Twinz (Srilakivarma).
