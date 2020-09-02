package config

var IdCheckUserCodes map[int]string
var UserCodes map[int]string
var PhoneCerCodes map[int]string
var DeviceQueryCodes map[int]string

func Init() {
    //init returnCodes
    UserCodes = make(map[int]string)
    UserCodes[9100] = "余额不足"
    UserCodes[9101] = "无权限操作"
    UserCodes[1100] = "成功"
    UserCodes[1901] = "请求超过单日上限"
    UserCodes[1902] = "参数不合法"
    UserCodes[1903] = "服务失败"

    //init idcheck user codes
    IdCheckUserCodes = make(map[int]string)
    IdCheckUserCodes[9100] = "余额不足"
    IdCheckUserCodes[9101] = "无权限操作"
    IdCheckUserCodes[1100] = "验证一致"
    IdCheckUserCodes[1101] = "验证不一致"
    IdCheckUserCodes[1103] = "库中无此号"
    IdCheckUserCodes[1104] = "卡状态错误"
    IdCheckUserCodes[1105] = "验证失败"
    IdCheckUserCodes[1902] = "参数不合法"
    IdCheckUserCodes[1903] = "服务失败"

    //init phoneCertification user codes
    PhoneCerCodes = make(map[int]string)
    PhoneCerCodes[9100] = "余额不足"
    PhoneCerCodes[9101] = "无权限操作"
    PhoneCerCodes[1100] = "认证通过"
    PhoneCerCodes[1101] = "认证未通过"
    PhoneCerCodes[1102] = "号码状态有误"
    PhoneCerCodes[1103] = "查无数据"
    PhoneCerCodes[1105] = "查有数据"
    PhoneCerCodes[1902] = "参数不合法"
    PhoneCerCodes[1903] = "服务失败"

    DeviceQueryCodes = make(map[int]string)
    DeviceQueryCodes[1100] = "成功"
    DeviceQueryCodes[1101] = "deviceId不存在"
    DeviceQueryCodes[1902] = "参数不合法"
    DeviceQueryCodes[1903] = "服务失败"
    DeviceQueryCodes[1904] = "无效deviceId"
    DeviceQueryCodes[9101] = "无权限操作"
}
