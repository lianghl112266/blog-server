package models

import (
	"blogServer/models/ctype"
)

// ArticleModel 文章表
type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                           // 文章标题
	Abstract      string         `json:"abstract"`                                       // 文章简介
	Content       string         `json:"content"`                                        // 文章内容
	LookCount     int            `json:"look_count"`                                     // 浏览量
	CommentCount  int            `json:"comment_count"`                                  // 评论量
	DiggCount     int            `json:"digg_count"`                                     // 点赞量
	CollectsCount int            `json:"collects_count"`                                 // 收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  // 文章的评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`                     // 文章作者
	UserID        uint           `json:"user_id"`                                        // 用户id
	Category      string         `gorm:"size:20" json:"category"`                        // 文章分类
	Source        string         `json:"source"`                                         // 文章来源
	Link          string         `json:"link"`                                           // 原文链接
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`                   // 文章封面
	BannerID      uint           `json:"banner_id"`                                      // 文章封面id
	NickName      string         `gorm:"size:42" json:"nick_name"`                       // 发布文章的用户昵称
	BannerPath    string         `json:"banner_path"`                                    // 文章的封面
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`                // 文章标签
}

//       },
//       "digg_count": {
//         "type": "integer"
//       },
//       "collects_count": {
//         "type": "integer"
//       },
//       "user_id": {
//         "type": "integer"
//       },
//       "user_nick_name": {
//         "type": "keyword"
//       },
//       "user_avatar": {
//         "type": "keyword"
//       },
//       "category": {
//         "type": "keyword"
//       },
//       "source": {
//         "type": "keyword"
//       },
//       "link": {
//         "type": "keyword"
//       },
//       "banner_id": {
//         "type": "integer"
//       },
//       "banner_url": {
//         "type": "keyword"
//       },
//       "tags": {
//         "type": "keyword"
//       },
//       "created_at":{
//         "type": "date",
//         "null_value": "null",
//         "format": "[yyyy-MM-dd HH:mm:ss]"
//       },
//       "updated_at":{
//         "type": "date",
//         "null_value": "null",
//         "format": "[yyyy-MM-dd HH:mm:ss]"
//       }
//     }
//   }
// }
// `
// }

// // IndexExists 索引是否存在
// func (a ArticleModel) IndexExists() bool {
// 	exists, err := global.ESClient.
// 		IndexExists(a.Index()).
// 		Do(context.Background())
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		return exists
// 	}
// 	return exists
// }

// // CreateIndex 创建索引
// func (a ArticleModel) CreateIndex() error {
// 	if a.IndexExists() {
// 		// 有索引
// 		a.RemoveIndex()
// 	}
// 	// 没有索引
// 	// 创建索引
// 	createIndex, err := global.ESClient.
// 		CreateIndex(a.Index()).
// 		BodyString(a.Mapping()).
// 		Do(context.Background())
// 	if err != nil {
// 		logrus.Error("创建索引失败")
// 		logrus.Error(err.Error())
// 		return err
// 	}
// 	if !createIndex.Acknowledged {
// 		logrus.Error("创建失败")
// 		return err
// 	}
// 	logrus.Infof("索引 %s 创建成功", a.Index())
// 	return nil
// }

// // RemoveIndex 删除索引
// func (a ArticleModel) RemoveIndex() error {
// 	logrus.Info("索引存在，删除索引")
// 	// 删除索引
// 	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
// 	if err != nil {
// 		logrus.Error("删除索引失败")
// 		logrus.Error(err.Error())
// 		return err
// 	}
// 	if !indexDelete.Acknowledged {
// 		logrus.Error("删除索引失败")
// 		return err
// 	}
// 	logrus.Info("索引删除成功")
// 	return nil
// }

// // Create 添加的方法
// func (a *ArticleModel) Create() (err error) {
// 	indexResponse, err := global.ESClient.Index().
// 		Index(a.Index()).
// 		BodyJson(a).Do(context.Background())
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		return err
// 	}
// 	a.ID = indexResponse.Id
// 	return nil
// }

// // ISExistData 是否存在该文章
// func (a ArticleModel) ISExistData() bool {
// 	res, err := global.ESClient.
// 		Search(a.Index()).
// 		Query(elastic.NewTermQuery("keyword", a.Title)).
// 		Size(1).
// 		Do(context.Background())
// 	if err != nil {
// 		logrus.Error(err.Error())
// 		return false
// 	}
// 	if res.Hits.TotalHits.Value > 0 {
// 		return true
// 	}
// 	return false
// }
// func (a *ArticleModel) GetDataByID(id string) error {
// 	res, err := global.ESClient.
// 		Get().
// 		Index(a.Index()).
// 		Id(id).
// 		Do(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(res.Source, a)
// 	return err
// }
