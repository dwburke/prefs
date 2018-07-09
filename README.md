# prefs
Preferences or configuration database, with the ability to specify
global settings, and per-user or per-server / region /data center/etc overrides.

# Work in progres... not yet fully functional

For example, say you have an application (or microservice) that you you want all
users to have the same value for the default, but each user could override it
with their own value.  A flag, a server address, a plugin version...

Configuration is handled by Viper so file formats are flexible...

Sample yaml config:

```
  prefs:
    port: 4441
    https: true
    search:
      - "{context}.someapp.{customer_id}.{key}"
      - "{context}.someapp.{key}"
  storage:
    type: "mysql"
    dns: "dbuser:password@/prefs?charset=utf8"
```

Example usage:

  `GET /prefs/dev/foo?customer_id=123456`

From the above config, it will search the keys in bottom up priority.  If 
customer_id is not in the query, or the derrived key does not have a value,
it will move on to the next key up.

":context" and ":key" are part of the http routing ('dev' and 'foo').
Everything else is optional.

# Adding to your own servce

```
import "github.com/dwburke/prefs/api/key"

...
key.SetupRoutes(r)
...

```

