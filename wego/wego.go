package wego

type hookfunc func() error

var (
	hooks = make([]hookfunc, 0)
)

func InitHooks(h ...hookfunc) {
	hooks = append(hooks, h...)
}

func RunHooks() {
	if len(hooks) > 0 {
		for _, v := range hooks {
			if err := v(); err != nil {
				panic(err)
			}
		}
	}
}

func InitBeforeHttp() {
	//init hooks
	//InitHooks()

	//run hooks
	RunHooks()

	//run http
	WegoApp.Run()
}

func Run() {
	InitBeforeHttp()
}
