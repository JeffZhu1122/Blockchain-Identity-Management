package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
	"strconv"
)

type RealEstateRequestBody struct {
	AccountId   string  `json:"accountId"`   //操作人ID
	Proprietor  string  `json:"proprietor"`  //所有者(业主)(业主AccountId)
	TotalArea   float64 `json:"totalArea"`   //总面积
	LivingSpace float64 `json:"livingSpace"` //生活空间

	Firstname string  `json:"firstname"`
	Lastname string  `json:"lastname"`
	DOB string  `json:"dob"`
	Gender string  `json:"gender"`
	Status string  `json:"status"`
	Photo string  `json:"photo"`
}

type RealEstateQueryRequestBody struct {
	Proprietor string `json:"proprietor"` //所有者(业主)(业主AccountId)
}

type DeletePeopleRequestBody struct {
	Key string  `json:"key"`
	Id1 string  `json:"id1"`
	Id2 string  `json:"id2"`
}
// @Summary 新建房地产(管理员)
// @Param realEstate body RealEstateRequestBody true "realEstate"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createRealEstate [post]
func CreateRealEstate(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateRequestBody)

	fmt.Println(body)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.TotalArea <= 0 || body.LivingSpace <= 0 || body.LivingSpace > body.TotalArea {
		appG.Response(http.StatusBadRequest, "失败", "TotalArea总面积和LivingSpace生活空间必须大于0，且生活空间小于等于总面积")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.TotalArea, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.LivingSpace, 'E', -1, 64)))

	bodyBytes = append(bodyBytes, []byte(body.Firstname))
	bodyBytes = append(bodyBytes, []byte(body.Lastname))
	bodyBytes = append(bodyBytes, []byte(body.DOB))
	bodyBytes = append(bodyBytes, []byte(body.Gender))
	bodyBytes = append(bodyBytes, []byte(body.Status))
	bodyBytes = append(bodyBytes, []byte(body.Photo))

	//调用智能合约
	resp, err := bc.ChannelExecute("createRealEstate", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	fmt.Println(data)
	appG.Response(http.StatusOK, "成功", data)
}

func DeletePeople(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DeletePeopleRequestBody)

	fmt.Println(body)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Key))
	bodyBytes = append(bodyBytes, []byte(body.Id1))
	bodyBytes = append(bodyBytes, []byte(body.Id2))
	
	//调用智能合约
	resp, err := bc.ChannelExecute("deleteData", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	fmt.Println(data)
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 获取房地产信息(空json{}可以查询所有，指定proprietor可以查询指定业主名下房产)
// @Param realEstateQuery body RealEstateQueryRequestBody true "realEstateQuery"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryRealEstateList [post]
func QueryRealEstateList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Proprietor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRealEstateList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
