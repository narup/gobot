## Build linux binary
```bash
GOOS=linux go build -o pulsenepal
```
## Copy to server 
```bash
gcloud compute scp pulsenepal mailtopuran@botsgrid:~/
```

## Enable Botsgrid Service 
```bash
sudo mv botsgrid.service /etc/systemd/system
sudo systemctl enable botsgrid
```

```bash
sudo systemctl status botsgrid
sudo systemctl start botsgrid
```

## Web parser
```bash
curl -H "x-api-key: <API_KEY>" "https://mercury.postlight.com/parser?url=https://www.youtube.com/watch?v=yod7ygeIHNo"
```