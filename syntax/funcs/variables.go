package main

// 不定参数，可以传入多个值

func YourName(name string, aliases ...string) {
	if len(aliases) > 0 {
		for i := 0; i < len(aliases); i++ {
			println(aliases[i])
		}
	}

}

func YourNameInvoke() {
	YourName("邓明")
	YourName("邓明", "deng ming")
	YourName("邓明", "deng ming", "Den")
}
