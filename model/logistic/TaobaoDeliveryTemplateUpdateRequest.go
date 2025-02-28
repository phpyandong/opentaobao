package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改运费模板 APIRequest
taobao.delivery.template.update

修改运费模板
*/
type TaobaoDeliveryTemplateUpdateRequest struct {
    model.Params

    // 模板名称，长度不能大于50个字节
    name   string 

    // 需要修改的模板对应的模板ID
    templateId   int64 

    // 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费
    assumer   int64 

    // 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 <br/>如: post;express;ems;cod <br/><br/><font color=red><br/>注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. <br/> </font><br/><br/><br/><font color=blue><br/>普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。<br/>货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font>
    templateTypes   string 

    // 邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
    templateDests   string 

    // 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    templateStartStandards   string 

    // 首费：输入0.01-999.99（最多包含两位小数）<br/><br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    templateStartFees   string 

    // 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    templateAddStandards   string 

    // 增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>
    templateAddFees   string 

}

func NewTaobaoDeliveryTemplateUpdateRequest() *TaobaoDeliveryTemplateUpdateRequest{
    return &TaobaoDeliveryTemplateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetApiMethodName() string {
    return "taobao.delivery.template.update"
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDeliveryTemplateUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateId(templateId int64) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateId() int64 {
    return r.templateId
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetAssumer(assumer int64) error {
    r.assumer = assumer
    r.Set("assumer", assumer)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetAssumer() int64 {
    return r.assumer
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateTypes(templateTypes string) error {
    r.templateTypes = templateTypes
    r.Set("template_types", templateTypes)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateTypes() string {
    return r.templateTypes
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateDests(templateDests string) error {
    r.templateDests = templateDests
    r.Set("template_dests", templateDests)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateDests() string {
    return r.templateDests
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateStartStandards(templateStartStandards string) error {
    r.templateStartStandards = templateStartStandards
    r.Set("template_start_standards", templateStartStandards)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateStartStandards() string {
    return r.templateStartStandards
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateStartFees(templateStartFees string) error {
    r.templateStartFees = templateStartFees
    r.Set("template_start_fees", templateStartFees)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateStartFees() string {
    return r.templateStartFees
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateAddStandards(templateAddStandards string) error {
    r.templateAddStandards = templateAddStandards
    r.Set("template_add_standards", templateAddStandards)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateAddStandards() string {
    return r.templateAddStandards
}

func (r *TaobaoDeliveryTemplateUpdateRequest) SetTemplateAddFees(templateAddFees string) error {
    r.templateAddFees = templateAddFees
    r.Set("template_add_fees", templateAddFees)
    return nil
}

func (r TaobaoDeliveryTemplateUpdateRequest) GetTemplateAddFees() string {
    return r.templateAddFees
}

