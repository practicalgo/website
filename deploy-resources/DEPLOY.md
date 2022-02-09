1. Build: `go build -o server`
2. SCP to EC2 host
3. SCP systemd service and socket files to ec2 host (if this is the first time) to `/etc/systemd/system/`
3. Install caddy: https://caddyserver.com/docs/install#fedora-redhat-centos (if this is the first time)
4. Copy Caddyfile to /etc/caddy/Cadddyfile
5. Start and enable systemd socket service, `practicalgo-website`
6. Start caddy


*new site update*

- GOOS=linux GOARCH=amd64 go build
- scp server root@ip:/usr/local/bin/practicalgo-website-1
- ssh root@ip mv /usr/local/bin/practicalgo-website-1 /usr/local/bin/practicalgo-website
- automatic restart and update!

*gotchas*

- firewalld
- selinux
