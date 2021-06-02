module github.com/aofiee/mpd

go 1.16

replace github.com/aofiee/mindpowerbank/helloworld => ../helloworld

replace github.com/aofiee/album => ../album

require (
	cloud.google.com/go v0.75.0 // indirect
	github.com/GoogleCloudPlatform/functions-framework-go v1.2.0
	github.com/aofiee/album v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.3.0
	golang.org/x/sys v0.0.0-20210223095934-7937bea0104d // indirect
)
