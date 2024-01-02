package usecase

import ioc "github.com/Ignaciojeria/einar-ioc"

type exampleUsecase struct {
}

var ExampleUsecase = ioc.UseCase[exampleUsecase](func() (exampleUsecase, error) {
	usecase := exampleUsecase{}
	return usecase, nil
})
