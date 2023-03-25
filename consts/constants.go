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
