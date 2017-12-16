# Nightmare Monkey
Nightmare Monkey was inspired by [Netflix Chaos Monkey](https://github.com/Netflix/chaosmonkey).
Nightmare Monkey randomly makes accidents to OS and processes. 
Quoting from Chaos Monkey, 

>Exposing engineers to failures more frequently incentivizes them to build resilient services.


## Requirements

- Linux

## Install locally
See [releases](https://github.com/etsxxx/nightmare-monkey/releases).

## How to use

You must specify `--execute` option. 
If you do not, then the monkey will be gentle. (a.k.a. Dryrun mode)

```
OPTIONS:
   --execute               You must specify this option to see nightmare. If you do not, then the monkey will be gentle.
   --min-interval value    minimum interval between nightmare (sec) (default: 3600)
   --max-interval value    max interval between nightmare (sec) (default: 21600)
   --day value, -d value   specify the day of week to allow nightmare, like crontab. 0-7 (0 or 7 is Sun) (default: "1-5")
   --time value, -t value  specify the time to allow nightmare. 'HH:MM-HH:MM' format. (default: "11:00-18:00")
   --port value, -p value  specify the port on which the API listens for connections (default: 8080)
   --help, -h              show help
   --version, -v           print the version
```

## Notes

- **DO NOT USE IN PRODUCTION ENVIRONMENTS**

