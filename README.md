## Logger interface with implementations (Zap and Logrus)

When we create go libraries in general we shouldn't be logging but at times we do have to log, debug what the 
library is doing or trace the log. 

We cannot implement a library with one log library and expect applications to use the same log library. We use two 
of the popular log libraries [logrus](https://github.com/sirupsen/logrus) and [zap](https://github.com/uber-go/zap)
and this `go-logger` library allows you to use either by using an interface. 

You can add your implementation if you want to add more log libraries (e.g. zerolog).