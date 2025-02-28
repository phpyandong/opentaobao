package taotv

// Videos 
type Videos struct {

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 视频标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 视频来源
    
    From   string `json:"from,omitempty" xml:"from,omitempty"`
    

    // 视频ID
    
    VideoId   string `json:"video_id,omitempty" xml:"video_id,omitempty"`
    

    // 视频图片
    
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    

}
