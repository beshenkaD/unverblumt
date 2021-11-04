# Unverblümt
Modularized general purpose telegram bot.

## Installing
### From source
**1. Clone this repo and run:**
``` shell script
go build
```

This command will build the bot itself

**2. Building modules**  
To build needed modules use:
``` shell script
make {needed modules}
```
Or to build all modules:
``` shell script
make modules
```

**3. Configuring**  
To configure bot export some environment variables:
``` shell script
export UNVERBLUMT_TELEGRAM="Your token"
export UNVERBLUMT_TIMEOUT="10s"
export UNVERBLUMT_MODULES="any, modules, you, like"
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
    -e UNVERBLUMT_TELEGRAM="Your token" \
    -e UNVERBLUMT_TIMEOUT="10s" \
    -e UNVERBLUMT_MODULES="any, modules, you, like" \
    u
```

## How to write modules?
**1. Import some packages**
``` go
import (
	"github.com/beshenkaD/unverblumt/core"
	tb "gopkg.in/tucnak/telebot.v3"
)
```

**2. Write your awesome commands**
``` go
func hello(c tb.Context) error {
	return c.Send("Hello from module!")
}
```

**3. Export your module with `Init` function**
``` go
func Init() *core.Module {
	return &core.Module{
		Name:        "Hello Module",
		License:     "My favorite license",
		Author:      "Me",
		Version:     "0.0.1",
		Description: "My module",

		ActiveCommands: map[string]core.Command{
			"/hello": {
				Handler:     hello,
				Description: "My command",
			},
		},
	}
}
```

# License
All the code in this repository is released under the GPL v2 License. Take a look
at the [LICENSE](LICENSE) for more info.
