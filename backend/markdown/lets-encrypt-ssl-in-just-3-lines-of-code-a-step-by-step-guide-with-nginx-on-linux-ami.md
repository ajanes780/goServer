# Let’s Encrypt SSL in Just 3 Lines of Code: A Step-by-Step Guide with Nginx on Linux AMI
![hero image](/images/lets-encrypt.png)

Setting up SSL does not have to be costly and complicated; this has been the fastest and easiest way to set up SSL for a new web server for me.

Amazon Linux AMIs come pre-packaged with python and pip, so this is extremely quick and easy
Step one:
SSH into the EC2 instance
Step two:
Install cert bot.

```bash
sudo pip3 install certbot certbot-nginx
```

Now that you have installed Certbot, we can configure our Nginx server with SSL. For your installation to be completed, your DNS records must be pointed at your EC2 instance
Step three:
Install the certificate

```bash
sudo /usr/local/bin/certbot --nginx \ 
-d example.com -d www.example.com \ 
--agree-tos -m your@email.com
```
Now your server is configured with a free SSL cert for 90days

We can also automate the renewal

Step four:

```bash 
echo "0 0,12 * * * root /opt/certbot/bin/python -c \ 
'import random; import time; time.sleep(random.random() * 3600)' \
&& sudo certbot renew -q" | sudo tee -a /etc/crontab > /dev/null
```
As promised, three lines of code and you have working SSL certificates for your Nginx server
If you enjoyed this article, please make sure to clap and follow
your support is appreciated


[Certbot Documentation](https://certbot.eff.org/docs/using.html#certbot-command-line-options)

#### Written on: 2021-10-10 
#### Written by: Aaron Janes
