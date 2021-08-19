package default_value

import (
	"fmt"
	"github.com/kuhufu/util/pprint"
	"testing"
	"time"
)

func Test_DefaultValue(t *testing.T) {
	f := struct {
		Name string `default:""`
		Age  int    `default:""`
	}{}

	DefaultValue(&f)

	f2 := struct {
		Name string
	}{}
	DefaultValue(&f2)

	fmt.Println(f)
	fmt.Println(f2)
}

func Test_time(t *testing.T) {
	f := struct {
		T1 time.Time `default:"now"`
		T2 time.Time `default:"1591953335" time_format:"unix"`
		T3 time.Time `default:"1591953335000" time_format:"unixmilli"`
		T4 time.Time `default:"2020-06-12" time_format:"2006-01-02"`
	}{}

	DefaultValue(&f)

	pprint.Print(f)
}
