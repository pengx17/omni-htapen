## Run https server with the following:
```
# Renew cert
sudo certbot renew --dry-run
# Run with docker
sudo docker run -p 443:443 -v /etc/letsencrypt/:/etc/letsencrypt/ pengxiao/omni-htapen -port 443
```
