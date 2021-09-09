module github.com/InVisionApp/opentelemetry-go-contrib/plugins/gorilla/mux

go 1.14

replace github.com/InVisionApp/opentelemetry-go-contrib/ => ../../..

require (
	github.com/gorilla/mux v1.7.4
	github.com/stretchr/testify v1.4.0
	github.com/InVisionApp/opentelemetry-go-contrib/ v0.0.0-00010101000000-000000000000
	github.com/InVisionApp/opentelemetry-go v0.4.2
)
