package iot

// BusinessRecipeOpenParam 
type BusinessRecipeOpenParam struct {

    // 菜谱描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 品类Id
    
    DevTypeId   int64 `json:"dev_type_id,omitempty" xml:"dev_type_id,omitempty"`
    

    // 食谱id
    
    RecipeId   int64 `json:"recipe_id,omitempty" xml:"recipe_id,omitempty"`
    

    // 开放账号id
    
    OpenAccountId   string `json:"open_account_id,omitempty" xml:"open_account_id,omitempty"`
    

    // 食谱图片
    
    RecipeImage   *ImageUrlParam `json:"recipe_image,omitempty" xml:"recipe_image,omitempty"`
    

    // 重量及单位
    
    RecipeIngredientList   []RecipeIngredientParam `json:"recipe_ingredient_list,omitempty" xml:"recipe_ingredient_list,omitempty"`
    

    // 食谱中文名
    
    RecipeNameCn   string `json:"recipe_name_cn,omitempty" xml:"recipe_name_cn,omitempty"`
    

    // 食谱类型 0视频菜谱 1图文菜谱
    
    RecipeType   int64 `json:"recipe_type,omitempty" xml:"recipe_type,omitempty"`
    

    // 食谱视频
    
    RecipeVideo   *VideoUrlParam `json:"recipe_video,omitempty" xml:"recipe_video,omitempty"`
    

    // 标签列表
    
    TagIdList   []int64 `json:"tag_id_list,omitempty" xml:"tag_id_list>int64,omitempty"`
    

    // 菜谱提示信息
    
    Tips   string `json:"tips,omitempty" xml:"tips,omitempty"`
    

}
