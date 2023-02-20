package api

const (
	URL             = "http://api.feieyun.cn/Api/Open/" // 飞鹅API请求接口
	AddList         = "Open_printerAddlist"             // 添加打印机
	PrintMSG        = "Open_printMsg"                   // 小票打印
	PrintLabelMSG   = "Open_printLabelMsg"              // 标签打印
	DelList         = "Open_printerDelList"             // 删除打印机
	EdtList         = "Open_printerEdit"                //  修改打印机
	DelSqs          = "Open_delPrinterSqs"              // 清空待打印信息队列
	OrderStatus     = "Open_queryOrderState"            // 根据订单Id查询 是否打印成功
	OrderInfoByDate = "Open_queryOrderInfoByDate"       // 查询打印机某天的订单详情，返回已打印订单数和等待打印订单数
	PrinterStatus   = "Open_queryPrinterStatus"         // 查询打印机状态
)
