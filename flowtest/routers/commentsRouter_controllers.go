package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowAccessContext",
			Router: `/flowaccesscontext`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowAccessContextList",
			Router: `/flowaccesscontextlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowAction",
			Router: `/flowaction`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowActionList",
			Router: `/flowactionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowDoc",
			Router: `/flowdoc`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowEvent",
			Router: `/flowdocevent`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowEventList",
			Router: `/flowdoceventlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowDocList",
			Router: `/flowdoclist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowDocumentDetail",
			Router: `/flowdocumentdetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowDocumentList",
			Router: `/flowdocumentlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowGroup",
			Router: `/flowgroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowGroupList",
			Router: `/flowgrouplist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowGroupRole",
			Router: `/flowgrouprole`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowGroupRoleList",
			Router: `/flowgrouprolelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowGroupUsersList",
			Router: `/flowgroupuserslist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowNext",
			Router: `/flownext`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowNode",
			Router: `/flownode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowNodeList",
			Router: `/flownodelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowPermission",
			Router: `/flowpermission`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowRole",
			Router: `/flowrole`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowRoleList",
			Router: `/flowrolelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowRolePermissionList",
			Router: `/flowrolepermissionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowState",
			Router: `/flowstate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowStateList",
			Router: `/flowstatelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowTransition",
			Router: `/flowtransition`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowTransitionList",
			Router: `/flowtransitionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowType",
			Router: `/flowtype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowTypeDelete",
			Router: `/flowtypedelete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowTypeList",
			Router: `/flowtypelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowTypeUpdate",
			Router: `/flowtypeupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowUser",
			Router: `/flowuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowUserGroup",
			Router: `/flowusergroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowUserList",
			Router: `/flowuserlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowWorkflow",
			Router: `/flowworkflow`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "FlowWorkflowList",
			Router: `/flowworkflowlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/flowtest/controllers:MainController"],
		beego.ControllerComments{
			Method: "WorkFlow",
			Router: `/workflow`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
