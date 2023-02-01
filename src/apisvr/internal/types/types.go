// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Uid         int64  `json:"uid,string,optional"`         // 用户id
	UserName    string `json:"userName,optional"`           // 用户名(唯一)
	Password    string `json:"password,omitempty"`          // 登录密码
	Email       string `json:"email,optional"`              // 邮箱
	Phone       string `json:"phone,optional"`              // 手机号
	Wechat      string `json:"wechat,optional"`             // 微信UnionID
	LastIP      string `json:"lastIP,optional"`             // 最后登录ip
	RegIP       string `json:"regIP,optional"`              // 注册ip
	NickName    string `json:"nickName,optional"`           // 用户的昵称
	City        string `json:"city,optional"`               // 用户所在城市
	Country     string `json:"country,optional"`            // 用户所在国家
	Province    string `json:"province,optional"`           // 用户所在省份
	Language    string `json:"language,optional"`           // 用户的语言，简体中文为zh_CN
	HeadImgUrl  string `json:"headImgUrl,optional"`         // 用户头像
	CreatedTime int64  `json:"createdTime,string,optional"` // 创建时间
	Role        int64  `json:"role"`                        // 用户角色
	Sex         int64  `json:"sex,optional"`                // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
}

type UserCreateReq struct {
	ReqType     string `json:"reqType,options=phone|wxopen|wxin|wxminip|pwd"` //注册方式:	phone手机号注册 wxopen 微信开放平台登录 wxin 微信内登录 wxminip 微信小程序 密码方式 pwd 账密方式 必输
	UserName    string `json:"userName"`                                      //手机号注册时填写手机号,账密登录时填写用户账号 必输
	Password    string `json:"password"`                                      //明文密码 必输，且做大小写校验
	Wechat      string `json:"wechat,optional"`                               // 微信UnionID
	LastIP      string `json:"lastIP,optional"`                               // 最后登录ip
	RegIP       string `json:"regIP,optional"`                                // 注册ip
	NickName    string `json:"nickName,optional"`                             // 用户的昵称
	City        string `json:"city,optional"`                                 // 用户所在城市
	Country     string `json:"country,optional"`                              // 用户所在国家
	Province    string `json:"province,optional"`                             // 用户所在省份
	Language    string `json:"language,optional"`                             // 用户的语言，简体中文为zh_CN
	HeadImgUrl  string `json:"headImgUrl,optional"`                           // 用户头像
	CreatedTime int64  `json:"createdTime,string,optional"`                   // 创建时间
	Role        int64  `json:"role"`                                          // 用户角色
	Sex         int64  `json:"sex,optional"`                                  // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
}

type UserCreateResp struct {
	Uid int64 `json:"uid,string"` //用户id
}

type UserCaptchaReq struct {
	Data string `json:"data,optional"`              //短信验证时填写手机号
	Type string `json:"type,options=sms|image"`     //验证方式:短信验证,图片验证码
	Use  string `json:"use,options=login|register"` //用途
}

type UserCaptchaResp struct {
	CodeID string `json:"codeID"`       //验证码编号
	Url    string `json:"url,optional"` //图片验证码url
	Expire int64  `json:"expire"`       //过期时间
}

type UserIndexReq struct {
	Page     PageInfo `json:"page"`              //分页信息
	UserName string   `json:"userName,optional"` //用户名(唯一)
	Phone    string   `json:"phone,optional"`    // 手机号
	Email    string   `json:"email,optional"`    // 邮箱
}

type UserIndexResp struct {
	List  []*UserInfo `json:"list,omitempty"`           //用户信息列表
	Total int64       `json:"total,optional,omitempty"` //总数
}

type UserUpdateReq struct {
	Uid        int64  `json:"uid,string"`          // 用户id
	UserName   string `json:"userName,optional"`   // 用户名(唯一)
	Email      string `json:"email,optional"`      // 邮箱
	NickName   string `json:"nickName,optional"`   // 用户的昵称
	City       string `json:"city,optional"`       // 用户所在城市
	Country    string `json:"country,optional"`    // 用户所在国家
	Province   string `json:"province,optional"`   // 用户所在省份
	Language   string `json:"language,optional"`   // 用户的语言，简体中文为zh_CN
	HeadImgUrl string `json:"headImgUrl,optional"` // 用户头像
	Sex        int64  `json:"sex,optional"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Role       int64  `json:"role,optional"`       // 用户角色
}

type UserReadReq struct {
	Uid int64 `json:"uid,string"` // 用户id
}

type UserReadResp struct {
	Uid         int64  `json:"uid,string,optional"`         // 用户id
	UserName    string `json:"userName,optional"`           // 用户名(唯一)
	Email       string `json:"email,optional"`              // 邮箱
	Phone       string `json:"phone,optional"`              // 手机号
	Wechat      string `json:"wechat,optional"`             // 微信UnionID
	LastIP      string `json:"lastIP,optional"`             // 最后登录ip
	RegIP       string `json:"regIP,optional"`              // 注册ip
	NickName    string `json:"nickName,optional"`           // 用户的昵称
	City        string `json:"city,optional"`               // 用户所在城市
	Country     string `json:"country,optional"`            // 用户所在国家
	Province    string `json:"province,optional"`           // 用户所在省份
	Language    string `json:"language,optional"`           // 用户的语言，简体中文为zh_CN
	HeadImgUrl  string `json:"headImgUrl,optional"`         // 用户头像
	CreatedTime int64  `json:"createdTime,string,optional"` // 创建时间
	Role        int64  `json:"role"`                        // 用户角色
	Sex         int64  `json:"sex,optional"`                // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
}

type UserDeleteReq struct {
	Uid string `json:"uid,omitempty"` //用户id
}

type UserLoginReq struct {
	UserID    string `json:"userID"`                                        //登录账号(支持用户名,手机号登录) 账号密码登录时需要填写
	PwdType   int32  `json:"pwdType"`                                       //账号密码登录时需要填写.1,无密码 2，明文 3，md5加密
	Password  string `json:"password"`                                      //密码，建议md5转换 密码登录时需要填写
	LoginType string `json:"loginType,options=sms|pwd|wxopen|wxin|wxminip"` //验证类型 sms 短信验证码 pwd 账号密码登录 wxopen 微信开放平台登录 wxin 微信内登录 wxminip 微信小程序
	Code      string `json:"code,optional"`                                 //验证码    微信登录填code
	CodeID    string `json:"codeID,optional"`                               //验证码编号 微信登录填state
}

type UserLoginResp struct {
	Info  UserInfo `json:"info"`  //用户信息
	Token JwtToken `json:"token"` //用户token
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`         //用户token
	AccessExpire int64  `json:"accessExpire,string,omitempty"` //token过期时间
	RefreshAfter int64  `json:"refreshAfter,string,omitempty"` //token刷新时间
}

type UserResourceReadResp struct {
	Menu []*MenuData `json:"menu"` //菜单资源
}

type PageInfo struct {
	Page int64 `json:"page,optional" form:"page,optional"` // 页码
	Size int64 `json:"size,optional" form:"size,optional"` // 每页大小
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Point struct {
	Longitude float64 `json:"longitude,range=[0:180]"` //经度
	Latitude  float64 `json:"latitude,range=[0:90]"`   //纬度
}

type DateRange struct {
	Start string `json:"start,optional "` //开始时间 格式：yyyy-mm-dd
	End   string `json:"end,optional "`   //结束时间 格式：yyyy-mm-dd
}

type MenuCreateReq struct {
	Name       string `json:"name"`                // 菜单名称
	ParentID   int64  `json:"parentID,optional"`   // 父菜单ID，一级菜单为1
	Type       int64  `json:"type,optional"`       // 类型   1：目录   2：菜单   3：按钮
	Path       string `json:"path,optional"`       // 系统的path
	Component  string `json:"component,optional"`  // 页面
	Icon       string `json:"icon,optional"`       // 菜单图标
	Redirect   string `json:"redirect,optional"`   // 路由重定向
	Order      int64  `json:"order,optional"`      // 左侧table排序序号
	HideInMenu int64  `json:"hideInMenu,optional"` // 菜单是否隐藏 1：是 2：否
}

type MenuIndexReq struct {
	Name string `json:"name,optional"` // 按菜单名称筛选
	Path string `json:"path,optional"` // 按菜单路径筛选
}

type MenuData struct {
	ID         int64  `json:"id"`                  // 编号
	Name       string `json:"name"`                // 菜单名称
	ParentID   int64  `json:"parentID"`            // 父菜单ID，一级菜单为1
	Type       int64  `json:"type"`                // 类型   1：目录   2：菜单   3：按钮
	Path       string `json:"path"`                // 系统的path
	Component  string `json:"component"`           // 页面
	Icon       string `json:"icon"`                // 菜单图标
	Redirect   string `json:"redirect"`            // 路由重定向
	CreateTime int64  `json:"createTime"`          // 创建时间
	Order      int64  `json:"order"`               // 左侧table排序序号
	HideInMenu int64  `json:"hideInMenu,optional"` // 菜单是否隐藏 1：是 2：否
}

type MenuIndexResp struct {
	List []*MenuData `json:"list"` //菜单列表
}

type MenuUpdateReq struct {
	ID         int64  `json:"id"`                  // 编号
	Name       string `json:"name"`                // 菜单名称
	ParentID   int64  `json:"parentID"`            // 父菜单ID，一级菜单为1
	Type       int64  `json:"type,optional"`       // 类型   1：目录   2：菜单   3：按钮
	Path       string `json:"path,optional"`       // 系统的path
	Component  string `json:"component,optional"`  // 页面
	Icon       string `json:"icon,optional"`       // 菜单图标
	Redirect   string `json:"redirect,optional"`   // 路由重定向
	Order      int64  `json:"order"`               // 左侧table排序序号
	HideInMenu int64  `json:"hideInMenu,optional"` // 菜单是否隐藏 1：是 2：否
}

type MenuDeleteReq struct {
	ID int64 `json:"id"` // 编号
}

type RoleCreateReq struct {
	Name   string `json:"name"`            // 角色名称
	Remark string `json:"remark,optional"` // 备注
	Status int64  `json:"status,optional"` // 状态 1:启用,2:禁用
}

type RoleIndexReq struct {
	Page   PageInfo `json:"page"`             //分页信息,只获取一个则不填
	Name   string   `json:"name,optional "`   //按名称查找角色
	Status int64    `json:"status,optional "` //按状态查找角色
}

type RoleIndexData struct {
	ID          int64   `json:"id"`          // 编号
	Name        string  `json:"name"`        // 角色名称
	Remark      string  `json:"remark"`      // 备注
	CreatedTime int64   `json:"createdTime"` // 创建时间
	Status      int64   `json:"status"`      // 角色状态
	RoleMenuID  []int64 `json:"roleMenuID"`  // 角色对应的菜单id列表
}

type RoleIndexResp struct {
	List  []*RoleIndexData `json:"list"`  //角色列表数据
	Total int64            `json:"total"` //角色列表总数
}

type RoleUpdateReq struct {
	ID     int64  `json:"id"`     // 编号
	Name   string `json:"name"`   // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64  `json:"status"` // 状态
}

type RoleDeleteReq struct {
	ID int64 `json:"id"` //编号
}

type RoleMenuUpdateReq struct {
	ID     int64   `json:"id"`     //角色编号
	MenuID []int64 `json:"menuID"` //菜单编号列表
}

type SysLogLoginIndexReq struct {
	Page          PageInfo  `json:"page"`                    //分页信息,只获取一个则不填
	IpAddr        string    `json:"ipAddr,optional "`        //按ip地址查找
	LoginLocation string    `json:"loginLocation,optional "` //按登录地址查找
	Date          DateRange `json:"date,optional "`          //按时间范围查找
}

type SysLogLoginIndexData struct {
	UserName      string `json:"userName"`      // 登录账号
	IpAddr        string `json:"ipAddr"`        // 登录IP地址
	LoginLocation string `json:"loginLocation"` // 登录地点
	Browser       string `json:"browser"`       // 浏览器类型
	Os            string `json:"os"`            // 操作系统
	Code          int64  `json:"code"`          // 登录状态（200成功 其它失败）
	Msg           string `json:"msg"`           // 提示消息
	CreatedTime   int64  `json:"createdTime"`   // 登录时间
}

type SysLogLoginIndexResp struct {
	List  []*SysLogLoginIndexData `json:"list"`  //登录日志列表数据
	Total int64                   `json:"total"` //登录日志列表总记录数
}

type SysLogOperIndexReq struct {
	Page         PageInfo `json:"page"`                   //分页信息,只获取一个则不填
	OperName     string   `json:"operName,optional "`     //按操作名称查找
	OperUserName string   `json:"operUserName,optional "` //按操作人员名称查找
	BusinessType int64    `json:"businessType,optional "` //按业务类型（1新增 2修改 3删除 4查询）查找
}

type SysLogOperIndexData struct {
	OperUserName string `json:"operUserName"` //操作人员名称
	OperName     string `json:"operName"`     //操作名称
	BusinessType int64  `json:"businessType"` //业务类型（1新增 2修改 3删除 4查询）
	Uri          string `json:"uri"`          //请求地址
	OperIpAddr   string `json:"operIpAddr"`   //操作主机地址
	OperLocation string `json:"operLocation"` //操作地点
	Req          string `json:"req"`          //请求参数
	Resp         string `json:"resp"`         //返回参数
	Code         int64  `json:"code"`         //登录状态（200成功 其它失败）
	Msg          string `json:"msg"`          //提示消息
	CreatedTime  int64  `json:"createdTime"`  //操作时间
}

type SysLogOperIndexResp struct {
	List  []*SysLogOperIndexData `json:"list"`  //操作日志列表数据
	Total int64                  `json:"total"` //操作日志列表总记录数
}

type Map struct {
	Mode      string `json:"mode，options=baidu"` //坐标系 默认百度坐标系
	AccessKey string `json:"accessKey"`          //设备地图key
}

type ConfigResp struct {
	Map Map `json:"map"` //设备地图相关配置
}

type DeviceAuthLoginReq struct {
	Username    string `json:"username"`                       //用户名
	Password    string `json:"password,optional"`              //密码
	ClientID    string `json:"clientID"`                       //clientID
	Ip          string `json:"ip"`                             //访问的ip地址
	Certificate string `json:"certificate,optional,omitempty"` //客户端证书 base64后传过来
}

type DeviceAuthAccessReq struct {
	Username string `json:"username,omitempty"` //用户名
	Topic    string `json:"topic,omitempty"`    //主题
	ClientID string `json:"clientID,omitempty"` //clientID
	Access   string `json:"access,omitempty"`   //操作
	Ip       string `json:"ip,omitempty"`       //访问的ip地址
}

type DeviceAuthRootCheckReq struct {
	Username    string `json:"username,omitempty"`             //用户名
	Password    string `json:"password,optional,omitempty"`    //密码
	ClientID    string `json:"clientID,omitempty"`             //clientID
	Ip          string `json:"ip,omitempty"`                   //访问的ip地址
	Certificate []byte `json:"certificate,optional,omitempty"` //客户端证书
}

type DeviceMsgHubLogIndexReq struct {
	DeviceName string    `json:"deviceName,omitempty"`                //设备名
	ProductID  string    `json:"productID,omitempty"`                 //产品id 获取产品id下的所有设备信息
	TimeStart  int64     `json:"timeStart,string,optional,omitempty"` //获取时间的开始
	TimeEnd    int64     `json:"timeEnd,string,optional,omitempty"`   //时间的结束
	Page       *PageInfo `json:"page,optional"`                       //分页信息
	Actions    []string  `json:"actions,optional"`                    //过滤操作类型 connected:上线 disconnected:下线  property:属性 event:事件 action:操作 thing:物模型提交的操作为匹配的日志
	Topics     []string  `json:"topics,optional"`                     //过滤主题
	Content    string    `json:"content,optional"`                    //过滤内容
	RequestID  string    `json:"requestID,optional"`                  //过滤请求ID
}

type DeviceMsgHubLogIndexResp struct {
	List  []*DeviceMsgHubLogIndex `json:"list"`  //数据
	Total int64                   `json:"total"` //总数
}

type DeviceMsgHubLogIndex struct {
	Timestamp  int64  `json:"timestamp,string"`
	Action     string `json:"action"` //connected:上线 disconnected:下线  property:属性 event:事件 action:操作 thing:物模型提交的操作为匹配的日志
	RequestID  string `json:"requestID"`
	TranceID   string `json:"tranceID"`
	Topic      string `json:"topic"`
	Content    string `json:"content"`
	ResultType int64  `json:"resultType,string"`
}

type DeviceMsgSdkLogIndexReq struct {
	DeviceName string    `json:"deviceName,omitempty"`                //设备名
	ProductID  string    `json:"productID,omitempty"`                 //产品id 获取产品id下的所有设备信息
	TimeStart  int64     `json:"timeStart,string,optional,omitempty"` //获取时间的开始
	TimeEnd    int64     `json:"timeEnd,string,optional,omitempty"`   //时间的结束
	LogLevel   int       `json:"logLevel,optional"`                   //等级
	Page       *PageInfo `json:"page,optional"`                       //分页信息
}

type DeviceMsgSdkIndexResp struct {
	List  []*DeviceMsgSdkIndex `json:"list"`  //数据
	Total int64                `json:"total"` //总数
}

type DeviceMsgSdkIndex struct {
	Timestamp int64  `json:"timestamp,string"` //发生时间戳
	Loglevel  int64  `json:"loglevel"`         //日志级别 1)关闭 2)错误 3)告警 4)信息 5)调试
	Content   string `json:"content"`          //具体内容
}

type DeviceMsgPropertyLogIndexReq struct {
	DeviceNames []string  `json:"deviceNames,omitempty"`               //设备名(不填获取产品下所有设备)
	ProductID   string    `json:"productID,omitempty"`                 //产品id 获取产品id下的所有设备信息
	DataID      string    `json:"dataID,optional,omitempty"`           //获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值
	TimeStart   int64     `json:"timeStart,string,optional,omitempty"` //获取时间的开始
	TimeEnd     int64     `json:"timeEnd,string,optional,omitempty"`   //时间的结束
	Page        *PageInfo `json:"page,optional"`                       //分页信息
	Interval    int64     `json:"interval,optional"`                   //分页信息
	ArgFunc     string    `json:"argFunc,optional"`                    //分页信息
	Fill        string    `json:"fill,optional"`                       //填充模式 参考:https://docs.taosdata.com/taos-sql/distinguished/
	Order       int32     `json:"order,optional"`                      //时间排序 0:aes(默认,从久到近排序) 1:desc(时间从近到久排序)
}

type DeviceMsgPropertyLatestIndexReq struct {
	DeviceName string   `json:"deviceName,omitempty"`       //设备名
	ProductID  string   `json:"productID,omitempty"`        //产品id 获取产品id下的所有设备信息
	DataIDs    []string `json:"dataIDs,optional,omitempty"` //获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值
}

type DeviceMsgPropertyIndexResp struct {
	List  []*DeviceMsgPropertyIndex `json:"list"`  //数据
	Total int64                     `json:"total"` //总数
}

type DeviceMsgPropertyIndex struct {
	Timestamp int64  `json:"timestamp,string"` //发生时间戳
	DataID    string `json:"dataID"`           //获取的具体属性值
	Value     string `json:"value,omitempty"`  //获取到的值
}

type DeviceMsgEventLogIndexReq struct {
	DeviceNames []string  `json:"deviceNames,optional"`                //设备名(不填获取产品下所有设备)
	ProductID   string    `json:"productID,optional"`                  //产品id 获取产品id下的所有设备信息
	DataID      string    `json:"dataID,optional,omitempty"`           //获取的具体标识符的数据 如果不指定则获取所有属性数据,一个属性一条,如果没有获取到的不会返回值
	TimeStart   int64     `json:"timeStart,string,optional,omitempty"` //获取时间的开始
	TimeEnd     int64     `json:"timeEnd,string,optional,omitempty"`   //时间的结束
	Page        *PageInfo `json:"page,optional"`                       //分页信息
	Types       []string  `json:"types,optional"`                      //类型 事件类型: 信息:info  告警alert  故障:fault
}

type DeviceMsgEventIndexResp struct {
	List  []*DeviceMsgEventIndex `json:"list"`  //数据
	Total int64                  `json:"total"` //总数
}

type DeviceMsgEventIndex struct {
	Timestamp int64  `json:"timestamp,string"` //发生时间戳
	Type      string `json:"type,omitempty"`   //类型 事件类型: 信息:info  告警alert  故障:fault
	DataID    string `json:"dataID"`           //获取的具体属性值
	Params    string `json:"params,omitempty"` //获取到的值
}

type DeviceGateWayIndexReq struct {
	Page              *PageInfo `json:"page,optional"`     //分页信息 只获取一个则不填
	GateWayProductID  string    `json:"gateWayProductID"`  //产品ID
	GateWayDeviceName string    `json:"gateWaydeviceName"` //设备名称
}

type DeviceGateWayIndexResp struct {
	List  []*DeviceInfo `json:"list"`  //分组信息
	Total int64         `json:"total"` //总数(只有分页的时候会返回)
}

type DeviceGateWayMultiCreateReq struct {
	GateWayProductID  string        `json:"gateWayProductID"`  //产品ID
	GateWayDeviceName string        `json:"gateWaydeviceName"` //设备名称
	List              []*DeviceCore `json:"list,optional"`     //分组tag
}

type DeviceGateWayMultiDeleteReq struct {
	GateWayProductID  string        `json:"gateWayProductID"`  //产品ID
	GateWayDeviceName string        `json:"gateWaydeviceName"` //设备名称
	List              []*DeviceCore `json:"list,optional"`     //分组tag
}

type DeviceInfo struct {
	ProductID   string  `json:"productID"`                   //产品id 只读
	DeviceName  string  `json:"deviceName"`                  //设备名称 读写
	CreatedTime int64   `json:"createdTime,optional,string"` //创建时间 只读
	Secret      string  `json:"secret,optional"`             //设备秘钥 只读
	FirstLogin  int64   `json:"firstLogin,optional,string"`  //激活时间 只读
	LastLogin   int64   `json:"lastLogin,optional,string"`   //最后上线时间 只读
	Version     *string `json:"version,optional"`            // 固件版本  读写
	LogLevel    int64   `json:"logLevel,optional"`           // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写
	Cert        string  `json:"cert,optional"`               // 设备证书  只读
	Tags        []*Tag  `json:"tags,optional"`               // 设备tag
	IsOnline    int64   `json:"isOnline,optional"`           // 在线状态  1离线 2在线 只读
	Address     *string `json:"address,optional"`            //所在地址
	Position    *Point  `json:"position,optional"`           //设备定位,默认百度坐标系
}

type DeviceInfoCreateReq struct {
	ProductID  string  `json:"productID"`         //产品id 只读
	DeviceName string  `json:"deviceName"`        //设备名称 读写
	LogLevel   int64   `json:"logLevel,optional"` // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写
	Address    *string `json:"address,optional"`  //所在地址
	Position   *Point  `json:"position,optional"` //设备定位,默认百度坐标系
	Tags       []*Tag  `json:"tags,optional"`     // 设备tag
}

type DeviceInfoUpdateReq struct {
	ProductID  string  `json:"productID"`         //产品id 只读
	DeviceName string  `json:"deviceName"`        //设备名称 读写
	LogLevel   int64   `json:"logLevel,optional"` // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写
	Address    *string `json:"address,optional"`  //所在地址
	Position   *Point  `json:"position,optional"` //设备定位,默认百度坐标系
	Tags       []*Tag  `json:"tags,optional"`     // 设备tag
}

type DeviceInfoDeleteReq struct {
	ProductID  string `json:"productID"`  //产品id 只读
	DeviceName string `json:"deviceName"` //设备名称 读写
}

type DeviceInfoReadReq struct {
	ProductID  string `json:"productID,optional"` //产品id 为空时获取所有产品
	DeviceName string `json:"deviceName"`         //设备名称 读写
}

type DeviceInfoIndexReq struct {
	Page       *PageInfo `json:"page,optional"`       //分页信息 只获取一个则不填
	ProductID  string    `json:"productID,optional"`  //产品id 为空时获取所有产品
	DeviceName string    `json:"deviceName,optional"` //过滤条件:模糊查询 设备名
	Position   *Point    `json:"position,optional"`   //设备定位,默认百度坐标系，用于获取以该点为中心，Range范围内的设备列表，与Range连用
	Range      int64     `json:"range,optional"`      //过滤条件:距离坐标点固定范围内的设备 单位：米
	Tags       []*Tag    `json:"tags,optional"`       // key tag过滤查询,非模糊查询 为tag的名,value为tag对应的值
}

type DeviceInfoIndexResp struct {
	List  []*DeviceInfo `json:"list"`  //设备信息
	Total int64         `json:"total"` //总数(只有分页的时候会返回)
	Num   int64         `json:"num"`   //返回的数量
}

type DeviceCore struct {
	ProductID  string `json:"productID"`  //产品ID
	DeviceName string `json:"deviceName"` //设备名称
}

type DeviceCountReq struct {
	StartTime int64 `json:"startTime,optional" form:"startTime,optional"` //查询区间的开始时间（秒）
	EndTime   int64 `json:"endTime,optional" form:"endTime,optional"`     //查询区间的结束时间（秒）
}

type DeviceInfoCount struct {
	Online   int64 `json:"online"`   // 在线设备数
	Offline  int64 `json:"offline"`  // 离线设备数
	Inactive int64 `json:"inactive"` // 未激活设备数
	Unknown  int64 `json:"unknown"`  // 未知设备数（all = 在线+离线+未激活+未知）
}

type DeviceTypeCount struct {
	Device  int64 `json:"device"`  // 设备类型数
	Gateway int64 `json:"gateway"` // 网关类型数
	Subset  int64 `json:"subset"`  // 子设备类型数
	Unknown int64 `json:"unknown"` // 未知设备类型
}

type DeviceCountResp struct {
	DeviceInfoCount DeviceInfoCount `json:"deviceCount"`
	DeviceTypeCount DeviceTypeCount `json:"deviceTypeCount"`
}

type DeviceInteractSendMsgReq struct {
	Topic   string `json:"topic"`   //发送的topic
	Payload string `json:"payload"` //发送的数据
}

type DeviceInteractSendPropertyReq struct {
	ProductID     string `json:"productID"`     //产品id 获取产品id下的所有设备信息
	DeviceName    string `json:"deviceName"`    //设备名
	Data          string `json:"data"`          //属性数据, JSON格式字符串, 注意字段需要在物模型属性里定义
	DataTimestamp int64  `json:"dataTimestamp"` //上报数据UNIX时间戳, 仅对Method:reported有效
	Method        string `json:"method"`        //请求类型 , 不填该参数或者 desired 表示下发属性给设备, reported 表示模拟设备上报属性
}

type DeviceInteractSendPropertyResp struct {
	Code        int64  `json:"code"`        //设备返回状态码
	Status      string `json:"status"`      //返回状态
	ClientToken string `json:"clientToken"` //调用id
	Data        string `json:"data"`        //返回信息
}

type DeviceInteractSendActionReq struct {
	ProductID   string `json:"productID"`   //产品id 获取产品id下的所有设备信息
	DeviceName  string `json:"deviceName"`  //设备名
	ActionID    string `json:"actionId"`    //产品数据模板中行为功能的标识符，由开发者自行根据设备的应用场景定义
	InputParams string `json:"inputParams"` //输入参数
}

type DeviceInteractSendActionResp struct {
	ClientToken  string `json:"clientToken"`  //调用id
	OutputParams string `json:"outputParams"` //输出参数 注意：此字段可能返回 null，表示取不到有效值。
	Status       string `json:"status"`       //返回状态
	Code         int64  `json:"code"`         //设备返回状态码
}

type ProductInfo struct {
	CreatedTime  int64   `json:"createdTime,optional,string"` //创建时间 只读
	ProductID    string  `json:"productID,optional"`          //产品id 只读
	ProductName  string  `json:"productName,optional"`        //产品名称
	AuthMode     int64   `json:"authMode,optional"`           //认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64   `json:"deviceType,optional"`         //设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64   `json:"categoryID,optional"`         //产品品类
	NetType      int64   `json:"netType,optional"`            //通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64   `json:"dataProto,optional"`          //数据协议:1:自定义,2:数据模板
	AutoRegister int64   `json:"autoRegister,optional"`       //动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Secret       string  `json:"secret,optional"`             //动态注册产品秘钥 只读
	Desc         *string `json:"desc,optional"`               //描述
	Tags         []*Tag  `json:"tags,optional"`               // 产品tag
}

type ProductInfoReadReq struct {
	ProductID string `json:"productID"` //产品id
}

type ProductInfoCreateReq struct {
	ProductName  string  `json:"productName"`           //产品名称
	AuthMode     int64   `json:"authMode,optional"`     //认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64   `json:"deviceType,optional"`   //设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64   `json:"categoryID,optional"`   //产品品类
	NetType      int64   `json:"netType,optional"`      //通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64   `json:"dataProto,optional"`    //数据协议:1:自定义,2:数据模板
	AutoRegister int64   `json:"autoRegister,optional"` //动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Desc         *string `json:"desc,optional"`         //描述
	Tags         []*Tag  `json:"tags,optional"`         // 产品tag
}

type ProductInfoUpdateReq struct {
	ProductID    string  `json:"productID"`             //产品id 只读
	ProductName  string  `json:"productName,optional"`  //产品名称
	AuthMode     int64   `json:"authMode,optional"`     //认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64   `json:"deviceType,optional"`   //设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64   `json:"categoryID,optional"`   //产品品类
	NetType      int64   `json:"netType,optional"`      //通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64   `json:"dataProto,optional"`    //数据协议:1:自定义,2:数据模板
	AutoRegister int64   `json:"autoRegister,optional"` //动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Desc         *string `json:"desc,optional"`         //描述
	Tags         []*Tag  `json:"tags,optional"`         // 产品tag
}

type ProductInfoDeleteReq struct {
	ProductID string `json:"productID"` //产品id 只读
}

type ProductInfoIndexReq struct {
	Page        *PageInfo `json:"page,optional"`        //分页信息,只获取一个则不填
	ProductName string    `json:"productName,optional"` //过滤产品名称
	DeviceType  int64     `json:"deviceType,optional"`  //过滤设备类型:0:全部,1:设备,2:网关,3:子设备
	ProductIDs  []string  `json:"productIDs,optional"`  //过滤产品id列表
	Tags        []*Tag    `json:"tags,optional"`        // key tag过滤查询,非模糊查询 为tag的名,value为tag对应的值
}

type ProductInfoIndexResp struct {
	List  []*ProductInfo `json:"list"`           //产品信息
	Total int64          `json:"total,optional"` //拥有的总数
	Num   int64          `json:"num,optional"`   //返回的数量
}

type ProductSchemaTslReadReq struct {
	ProductID string `json:"productID"` //产品id
}

type ProductSchemaTslReadResp struct {
	Tsl string `json:"tsl"` //物模型tsl
}

type ProductSchemaUpdateReq struct {
	*ProductSchemaInfo
}

type ProductSchemaTslImportReq struct {
	ProductID string `json:"productID"` //产品id 只读
	Tsl       string `json:"tsl"`       //物模型tsl
}

type ProductSchemaCreateReq struct {
	*ProductSchemaInfo
}

type ProductSchemaDeleteReq struct {
	ProductID  string `json:"productID"`  //产品id
	Identifier string `json:"identifier"` //标识符
}

type ProductSchemaIndexReq struct {
	Page        *PageInfo `json:"page,optional"`        //分页信息,只获取一个则不填
	ProductID   string    `json:"productID"`            //产品id
	Type        int64     `json:"type,optional"`        //物模型类型 1:property属性 2:event事件 3:action行为
	Tag         int64     `json:"tag,optional"`         //过滤条件: 物模型标签 1:自定义 2:可选 3:必选
	Identifiers []string  `json:"identifiers,optional"` //过滤标识符列表
}

type ProductSchemaIndexResp struct {
	List  []*ProductSchemaInfo `json:"list"`  //分页信息,只获取一个则不填
	Total int64                `json:"total"` //总数(只有分页的时候会返回)
}

type ProductSchemaInfo struct {
	ProductID  string  `json:"productID"`           //产品id 只读
	Type       int64   `json:"type"`                //物模型类型 1:property属性 2:event事件 3:action行为
	Tag        int64   `json:"tag"`                 //物模型标签 1:自定义 2:可选 3:必选  必选不可删除
	Identifier string  `json:"identifier"`          //标识符
	Name       *string `json:"name,optional"`       //功能名称
	Desc       *string `json:"desc,optional"`       //描述
	Required   int64   `json:"required,optional"`   //是否必须 1:是 2:否
	Affordance *string `json:"affordance,optional"` //各功能类型的详细参数定义
}

type SchemaAction struct {
	Input  []*SchemaParam `json:"input,optional"`  //调用参数
	Output []*SchemaParam `json:"output,optional"` //返回参数
}

type SchemaProperty struct {
	Mode   string        `json:"mode,optional"` //读写类型: r(只读) rw(可读可写)
	Define *SchemaDefine `json:"define"`        //参数定义
}

type SchemaEvent struct {
	Type   string         `json:"type"`            //事件类型: 信息:info  告警alert  故障:fault
	Params []*SchemaParam `json:"params,optional"` //事件参数
}

type SchemaDefine struct {
	Type      string            `json:"type"`                //参数类型: bool int string struct float timestamp array enum
	Mapping   map[string]string `json:"mapping,omitempty"`   //枚举及bool类型:bool enum
	Min       string            `json:"min,omitempty"`       //数值最小值:int  float
	Max       string            `json:"max,omitempty"`       //数值最大值:int string float
	Start     string            `json:"start,omitempty"`     //初始值:int float
	Step      string            `json:"step,omitempty"`      //步长:int float
	Unit      string            `json:"unit,omitempty"`      //单位:int float
	Specs     []*SchemaSpec     `json:"specs,omitempty"`     //结构体:struct
	ArrayInfo *SchemaDefine     `json:"arrayInfo,omitempty"` //数组:array
}

type SchemaSpec struct {
	Identifier string        `json:"identifier"` //参数标识符
	Name       string        `json:"name"`       //参数名称
	DataType   *SchemaDefine `json:"dataType"`   //参数定义
}

type SchemaParam struct {
	Identifier string        `json:"identifier"`       //参数标识符
	Name       string        `json:"name"`             //参数名称
	Define     *SchemaDefine `json:"define,omitempty"` //参数定义
}

type ProductRemoteConfig struct {
	ID         int64  `json:"id"`         //配置编号
	Content    string `json:"content"`    //配置内容
	CreateTime string `json:"createTime"` //创建时间
}

type ProductRemoteConfigCreateReq struct {
	ProductID string `json:"productID"` //产品id
	Content   string `json:"content"`   //配置内容
}

type ProductRemoteConfigIndexReq struct {
	ProductID string    `json:"productID"`     //产品id
	Page      *PageInfo `json:"page,optional"` //分页信息
}

type ProductRemoteConfigIndexResp struct {
	List  []*ProductRemoteConfig `json:"list"`  //产品信息
	Total int64                  `json:"total"` //拥有的总数
}

type ProductRemoteConfigPushAllReq struct {
	ProductID string `json:"productID"` //产品id
}

type ProductRemoteConfigLastestReadReq struct {
	ProductID string `json:"productID"` //产品id
}

type ProductRemoteConfigLastestReadResp struct {
	ProductRemoteConfig
}

type GroupInfo struct {
	GroupID     int64  `json:"groupID,string"`     //分组ID
	GroupName   string `json:"groupName"`          //分组名称
	ParentID    int64  `json:"parentID,string"`    //父组ID
	CreatedTime int64  `json:"createdTime,string"` //创建时间
	Desc        string `json:"desc,optional"`      //分组描述
	Tags        []*Tag `json:"tags,optional"`      //分组tag
}

type GroupInfoCreateReq struct {
	GroupName string `json:"groupName"`       //分组名称
	ParentID  int64  `json:"parentID,string"` //父组ID
	Desc      string `json:"desc,optional"`   //分组描述
}

type GroupInfoIndexReq struct {
	Page      PageInfo `json:"page,optional"`      //分页信息 只获取一个则不填
	ParentID  int64    `json:"parentID,string"`    //父组ID
	GroupName string   `json:"groupName,optional"` //分组名称
	Tags      []*Tag   `json:"tags,optional"`      //分组tag
}

type GroupInfoIndexResp struct {
	List    []*GroupInfo `json:"list"`    //分组信息
	Total   int64        `json:"total"`   //总数(只有分页的时候会返回)
	ListAll []*GroupInfo `json:"listAll"` //完整分分组信息
}

type GroupInfoReadReq struct {
	GroupID int64 `json:"groupID,string"` //分组ID
}

type GroupInfoDeleteReq struct {
	GroupID int64 `json:"groupID,string"` //分组ID
}

type GroupInfoUpdateReq struct {
	GroupID   int64   `json:"groupID,string"`     //分组ID
	GroupName *string `json:"groupName,optional"` //分组名称
	Desc      *string `json:"desc,optional"`      //分组描述
	Tags      []*Tag  `json:"tags,optional"`      //分组tag
}

type GroupDeviceIndexReq struct {
	Page       PageInfo `json:"page,optional"`       //分页信息 只获取一个则不填
	GroupID    int64    `json:"groupID,string"`      //分组ID
	ProductID  string   `json:"productID,optional"`  //产品ID
	DeviceName string   `json:"deviceName,optional"` //设备名称
}

type GroupDeviceIndexResp struct {
	List  []*DeviceInfo `json:"list"`  //分组信息
	Total int64         `json:"total"` //总数(只有分页的时候会返回)
}

type GroupDeviceMultiCreateReq struct {
	GroupID int64         `json:"groupID,string"` //分组ID
	List    []*DeviceCore `json:"list,optional"`  //分组tag
}

type GroupDeviceMultiDeleteReq struct {
	GroupID int64         `json:"groupID,string"` //分组ID
	List    []*DeviceCore `json:"list,optional"`  //分组tag
}
