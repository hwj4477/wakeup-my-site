# Wake up My Site (in heroku free tier)
Heroku free tier must sleep if not active for 30 minutes.<br />
So I wanted to wake up every 30 minutes.<br />

## Usage ##

### go get
```
$ go get github.com/hwj4477/wakeup-my-site
```

### edit script
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
