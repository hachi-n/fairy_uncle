# FAIRY_UNCLE
## Specification.
 This tool launches once every 30 minutes.
 Its specifications are not good, so I plan to change it.

## SubCommands
```
  show
    show notification details.
  add
    add notification.
  destroy
    destroy notification.
  list
    display notification list.
  notify
    notify desktop.
```

## SETUP
```
  make build
    create builded module.
  install
    deploy this command to bin dir.
  setup
    write to crontab.
```

## JSON FORMAT(.config/appName)
```
{
    "NotificationName1": {
        "Title": "hoge",
        "Time": "00 19",
        "SubTitle": "hoge1",
        "Message": "fuga",
        "icon": "/path/to/icon",
        "Open": "https://github.com/hachi-n",
        "Sound": "default" #Blow,Bottle,Frog,Funk,Glass,Hero,Morse,Ping,Pop,Purr,Sosumi,Tink
    },
    "NotificationName2": {
        "Title": "hoge",
        "Time": "00 19",
        "SubTitle": "hoge1",
        "Message": "fuga",
        "icon": "/path/to/icon",
        "Open": "https://github.com/hachi-n",
        "Sound": "default" #Blow,Bottle,Frog,Funk,Glass,Hero,Morse,Ping,Pop,Purr,Sosumi,Tink
    },
    "NotificationName3": {
        "Title": "hoge",
        "Time": "00 19",
        "SubTitle": "hoge1",
        "Message": "fuga",
        "icon": "/path/to/icon",
        "Open": "https://github.com/hachi-n",
        "Sound": "default" #Blow,Bottle,Frog,Funk,Glass,Hero,Morse,Ping,Pop,Purr,Sosumi,Tink
    }
}
```

## Contents to be changed
- Desktop App
- Crontab To Daemon
- Abolish terminal-notify command

