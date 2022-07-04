# serversdf
Poll multiple servers to find the disk space available for a particular disk over SSH.

# Build
```bash
go build serversdf
```

# Usage
Create a file named "config.toml" with the full path to your `known_hosts` file and add a table with the server's information to the `[[server]]` array. Both password and key file authentication are supported.

Run the program in a directory that contains "config.toml" to execute the program.

```bash
$ ./serverusage
Server A                    2.2G / 25G (9%) used
Server B                    2.0G / 25G (9%) used
```

# Example Configuration
```toml
# config.toml

known_hosts = "/home/username/.ssh/known_hosts"

[[server]]
name = "Server A"                           # A shorthand name to label the server
host = "XXX.XXX.X.XX"                       # IP of the host
user = "username"                           # Username
password = "password"                       # Password authentication for SSH
disk = "/dev/vda1"                          # Disk mount point

[[server]]
name = "Server B"
host = "XXX.XXX.X.XX"
user = "username"
key_file = "/home/username/.ssh/id_rsa"     # Path to your key file
key_password = "password"                   # (Optional) Key file password
disk = "/dev/vda1"
```
