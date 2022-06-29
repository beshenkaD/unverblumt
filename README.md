# Unverblümt
General purpose telegram bot.

## Installing
### From source
**1. Clone this repo and run:**
``` shell script
go build
```

**3. Configuring**  
To configure bot export some environment variables:
``` shell script
export UNVERBLUMT_TOKEN="Your token"
```

Or write it to `.env` file

### Docker
**1. Build image**
``` shell script
docker build . --tag u
```

**2. Run it :)**
``` shell script
docker run \
    -e UNVERBLUMT_TOKEN="Your token" \
    u
```

# License
All the code in this repository is released under the GPL v2 License. Take a look
at the [LICENSE](LICENSE) for more info.
