package iot

// MediaItem 
type MediaItem struct {

    // json格式附加信息
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 专辑id
    
    AlbumId   string `json:"album_id,omitempty" xml:"album_id,omitempty"`
    

    // 歌曲源
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

    // 歌曲长度
    
    Length   int64 `json:"length,omitempty" xml:"length,omitempty"`
    

    // 是否为专辑
    
    IsAlbum   bool `json:"is_album,omitempty" xml:"is_album,omitempty"`
    

    // 歌曲id
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 扩展字段
    
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
    

}
