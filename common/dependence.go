package common

type ModelDependence struct {
	PackagePath string
	PackageName string
	FiltHook
}

func NewModelDependence(packagePath, packageName string, filtHook FiltHook) *ModelDependence {
	return &ModelDependence{
		PackagePath: packagePath,
		PackageName: packageName,
		FiltHook: filtHook,
	}
}

type DaoDependence struct {
	FiltHook
}

func NewDaoDependence(filtHook FiltHook) *DaoDependence {
	return &DaoDependence{
		FiltHook: filtHook,
	}
}

type ServiceDependence struct {
	FiltHook
}

func NewServiceDependence(filtHook FiltHook) *ServiceDependence {
	return &ServiceDependence{
		FiltHook: filtHook,
	}
}
