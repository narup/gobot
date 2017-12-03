
## Based on awesome tutorial by https://github.com/justforfunc

## Deploy to Compute Engine Insrance
## Build linux binary
```bash
GOOS=linux go build -o gobot
```
## Copy to server 
```bash
gcloud compute scp gobot <computer_engine_user>@<host_name>:~/
```

## Enable Botsgrid Service 
```bash
sudo mv gobot.service /etc/systemd/system
sudo systemctl enable gobot
```

```bash
sudo systemctl status gobot
sudo systemctl start gobot
```