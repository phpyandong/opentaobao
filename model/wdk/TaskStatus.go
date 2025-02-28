package wdk

// TaskStatus 
type TaskStatus struct {

    // 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
    
    NodeType   string `json:"node_type,omitempty" xml:"node_type,omitempty"`
    

    // warehouseCode, 出库仓，由基础店仓维护，盒马全域统一
    
    NodeCode   string `json:"node_code,omitempty" xml:"node_code,omitempty"`
    

    // 批次号
    
    BatchId   string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
    

    // 作业状态变更类型： START_PICK(“开始拣货”)， PICK_FINISH(“拣货完成”)， START_PACKAGE(“开始打包”), PACKAGE _FINISH(“打包完成”);
    
    StatusChangeType   string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
    

    // 作业状态变更时间
    
    StatusChangeTime   string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
    

    // 操作员
    
    OperatorCode   string `json:"operator_code,omitempty" xml:"operator_code,omitempty"`
    

    // 是否最终状态（打包完成\整批次所有sku全部缺货：true，其他：false）
    
    IsFinalStatus   bool `json:"is_final_status,omitempty" xml:"is_final_status,omitempty"`
    

    // 容器编号列表
    
    ContainerInfoList   []string `json:"container_info_list,omitempty" xml:"container_info_list>string,omitempty"`
    

    // 履约单列表
    
    FulfillOrderList   []FulfillOrder `json:"fulfill_order_list,omitempty" xml:"fulfill_order_list,omitempty"`
    

    // 容器数量
    
    ContainerCount   int64 `json:"container_count,omitempty" xml:"container_count,omitempty"`
    

    // 异形件数量
    
    AbnormalPackCount   string `json:"abnormal_pack_count,omitempty" xml:"abnormal_pack_count,omitempty"`
    

    // 扩展属性
    
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

}
