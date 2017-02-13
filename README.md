# Wake up My Site (In heroku free tier)
Heroku free tier must sleep if not active for 30 minutes.<br />
So I wanted to wake up my site every 30 minutes. And I wanted to logging.<br />

## Usage ##
You must have server. And install golang.<br />
I have use my raspberry pi.<br />

### go get
```
$ go get github.com/hwj4477/wakeup-my-site
```

### edit shell script
```
URL="your-site-url"

BIN_PATH="$HOME/go/src/github.com/hwj4477/wakeup-my-site"
export PATH=$BIN_PATH

exec wakeup-my-site $URL
```

### add crontab
```
$ crontab -e
```
