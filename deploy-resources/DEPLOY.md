1. Build: `go build -o server`
2. SCP to EC2 host
3. SCP systemd service to ec2 host (if this is the first time)
3. Install caddy: https://caddyserver.com/docs/install#fedora-redhat-centos (if this is the first time)
4. Copy Caddyfile to /etc/caddy/Cadddyfile
5. Start systemd service, practicalgowebsite
6. Start caddy


*gotchas*

- firewalld
- selinux
