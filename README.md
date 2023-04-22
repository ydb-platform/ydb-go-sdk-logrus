# logrus

logrus package helps to create ydb-go-sdk traces with logging driver events with logrus

## Usage
```go
import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"

	ydbLogrus "github.com/ydb-platform/ydb-go-sdk-logrus"
)

func main() {
	// init your logrus.Logger
	log := logrus.New(os.Stdout).With().Timestamp().Logger()

	db, err := ydb.Open(
		context.Background(),
		os.Getenv("YDB_CONNECTION_STRING"),
		ydbLogrus.WithTraces(
			&log,
			trace.DetailsAll,
		),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close(context.Background())
	}()

	// work with db
}
```
