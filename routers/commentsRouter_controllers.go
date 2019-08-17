package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:AcudienteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:FamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/informacion_familiar_crud/controllers:RelacionFamiliarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
