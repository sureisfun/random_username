Random, reddit-style usernames. 

```
import (
        "fmt"

        "github.com/sureisfun/random_username"
)

func main() {
        username := random_username.GenerateUsername()
        fmt.Println(username)
}
```

That'll yield something like "CuriousCheetah-2348" or whatever.
