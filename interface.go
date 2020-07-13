package timing

// Job job interface
type Job interface {
	Run()
}

// JobFunc job function
type JobFunc func()

// Run implement job interface
func (sf JobFunc) Run() {
	sf()
}
