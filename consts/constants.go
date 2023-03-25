package consts

const (
	NO_RECORD       = 0  //(未发送)
	SENT            = 1  //,"已发送"),
	TRANSIT         = 2  //,"转运中"),
	DELIVERY        = 3  //,"妥投"),
	TIMEOUT         = 4  //,"超时"),
	CUT_OFF         = 5  //,"扣关"),
	WRONG_ADDRESS   = 6  //,"地址错误"),
	EXPRESS_LOST    = 7  //,"快件丢失"),
	SEND_BACK       = 8  //,"退件"),
	OTHER_EXCEPTION = 9  //,"其他异常"),
	DESTORY         = 10 //,"销毁");
)

const (
	// inner tracking password
	PASSWORD = "TB1285644CD24CFB7BD7"
)

const (
	ZERO           = 0      //没有结果或零处理"),
	SUCCESS        = 1      //没有结果或零处理"),
	NEGATIVE_1     = -1     //唯一性字段值重复，操作失败"),
	NEGATIVE_2     = -2     //记录不存在，操作失败"),
	NEGATIVE_3     = -3     //未提供必须的请求参数，操作失败"),
	NEGATIVE_4     = -4     //请求不支持，版本错误或请求未实现"),
	NEGATIVE_7     = -7     //安全校验失败，不是配置的IP或数字签名错误"),
	NEGATIVE_8     = -8     //授权失败"),
	NEGATIVE_9     = -9     //API接口程序错误"),
	NEGATIVE_102   = -102   //运单不存在"),
	NEGATIVE_710   = -710   //icID 错误，未提供或小于1"),
	NEGATIVE_711   = -711   //icID 错误，客户不存在"),
	NEGATIVE_720   = -720   //TimeStamp 错误，超出了同步阈值"),
	NEGATIVE_730   = -730   //MD5 错误，长度不是32字符"),
	NEGATIVE_731   = -731   //MD5 错误，不匹配"),
	NEGATIVE_10001 = -10001 //cKey值不正确"),
	NEGATIVE_10002 = -10002 //钮门接口不支持批量查询");
)
