// 这些示例展示了旗帜程序包更复杂的用途。
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// 例 1：一个名为 "物种 "的单字符串标记，默认值为 "地鼠"。
var species = flag.String("species", "gopher", "the species we are studying")

// 例 2：两个标志共享一个变量，因此我们可以使用速记方法。初始化的顺序是未定义的，因此要确保两个变量使用相同的默认值。必须使用 init 函数对它们进行设置。
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// 例 3：用户自定义的标记类型，持续时间片。
type interval []time.Duration

// String 是格式化标志值的方法，是 flag.Value 接口的一部分。String 方法的输出将用于诊断。
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set 是设置标志值的方法，是 flag.Value 接口的一部分。Set 的参数是一个要解析的字符串，用于设置标志。这是一个以逗号分隔的列表，因此我们要对其进行拆分。
func (i *interval) Set(value string) error {
	// 如果我们想允许多次设置标志，累积数值，就可以删除 if 语句。这样就可以使用 -deltaT 10s -deltaT 15s 和其他组合。
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// 定义一个累计持续时间的标志。由于它的类型特殊，我们需要使用 Var 函数，因此需要在初始化过程中创建标记。

var intervalFlag interval

func init() {
	// 将命令行标志与 intervalFlag 变量绑定，并设置使用信息。
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// 所有有趣的部分都在上面声明的变量中，但要让 flag 软件包看到定义在那里的标志，必须执行（通常在 main（而不是 init！）开始时）：flag.Parse() 我们在这里不调用它，因为这段代码是一个名为 "Example "的函数，是软件包测试套件的一部分，它已经解析了标志。不过，在 pkg.go.dev 中查看时，该函数已更名为 "main"，可以作为独立示例运行。
}
