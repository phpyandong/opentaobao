package taotv

// PlayListNavRbo 
type PlayListNavRbo struct {

    // 当前播单id
    
    CurPlayListId   string `json:"cur_play_list_id,omitempty" xml:"cur_play_list_id,omitempty"`
    

    // 当前播单视频列表
    
    Videos   []Videos `json:"videos,omitempty" xml:"videos,omitempty"`
    

    // 播单列表
    
    PlayList   []Playlist `json:"play_list,omitempty" xml:"play_list,omitempty"`
    

}
