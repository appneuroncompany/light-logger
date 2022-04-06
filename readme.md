# Logger for production

Use this with to send the logs to [Fluentbit](https://fluentbit.io).

# Usage
```go
package main

import (
	logger "github.com/appneuroncompany/light-logger"
    "github.com/appneuroncompany/light-logger/clogger"
    "errors"
)

func main() {
    logger.Log.App = "App name" // use this on main

    clogger.Error(&logger.Messages{ // use it wherever you want
        "smt err: ": errors.New("erors").Error(),
    })

    clogger.Default(&logger.Messages{ // use it wherever you want
        "Default message: ": "message",
    })
        
    clogger.Info(&logger.Messages{ // use it wherever you want
        "Info message: ": "message",
    })

    clogger.Warning(&logger.Messages{ // use it wherever you want
        "Warning message: ": "message",
    })
}
```

This is how it will look like (will show colorful on linux terminal!):

```ruby
{
        "level": "ERROR",
        "@timestamp": "2022-04-06 13:24:34",
        "@version": "1",
        "app": "App name",
        "C:/path/to/main.go:13",
        "method": "main",
        "message": {
                "smt err: ": "errors"
        }
}
{
        "level": "DEFAULT",
        "@timestamp": "2022-04-06 13:24:34",
        "@version": "1",
        "app": "App name",
        "line": "C:/path/to/main.go:17",
        "method": "main",
        "message": {
                "Default message: ": "message"
        }
}
{
        "level": "INFO",
        "@timestamp": "2022-04-06 13:24:34",
        "@version": "1",
        "app": "App name",
        "line": "C:/path/to/main.go:21",
        "method": "main",
        "message": {
                "Info message: ": "message"
        }
}
{
        "level": "WARNING",
        "@timestamp": "2022-04-06 13:24:34",
        "@version": "1",
        "app": "App name",
        "line": "C:/path/to/main.go:25",
        "method": "main",
        "message": {
                "Warning message: ": "message"
        }
}
```
# Maintainers

Name               | Github          |
------------------ | --------------  |
Appneuron Comapany | [appneuroncompany](https://github.com/appneuroncompany)     |
Burak Halefoğlu    | [burakhalefoglu](https://github.com/burakhalefoglu)  |
Muhammed Sarı      | [muhammeedsari](https://github.com/muhammeedsari)   |

# License

MIT.