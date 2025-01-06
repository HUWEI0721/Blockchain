package controller

import (
	"backend/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 保存了野生水产品上链与查询的函数

// 野生水产品上链函数，新增Image字段的处理
func Uplink(c *gin.Context) {
	// 与userID不一样，取ID从第二位作为溯源码，即18位，雪花ID的结构如下（转化为10进制共19位）：
	// +--------------------------------------------------------------------------+
	// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	// +--------------------------------------------------------------------------+
	fisherman_traceability_code := pkg.GenerateID()[1:]
	args := buildArgs(c, fisherman_traceability_code)
	if args == nil {
		return
	}
	res, err := pkg.ChaincodeInvoke("Uplink", args)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "uplink failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":              200,
		"message":           "uplink success",
		"txid":              res,
		"traceability_code": args[1],
	})
}

// 获取野生水产品的上链信息
func GetSeafoodInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetSeafoodInfo", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败：" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取用户的野生水产品ID列表
func GetSeafoodList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(200, gin.H{
			"message": "未找到用户ID",
		})
		return
	}
	res, err := pkg.ChaincodeQuery("GetSeafoodList", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取所有的水产品信息
func GetAllSeafoodInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetAllSeafoodInfo", "")
	fmt.Print("res", res)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取野生水产品上链历史
func GetSeafoodHistory(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetSeafoodHistory", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// buildArgs 构建传递给链码的参数，新增Image字段的处理
func buildArgs(c *gin.Context, fisherman_traceability_code string) []string {
	var args []string
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(200, gin.H{
			"message": "未找到用户ID",
		})
		return nil
	}

	// 获取用户类型
	userType, err := pkg.ChaincodeQuery("GetUserType", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "获取用户类型失败：" + err.Error(),
		})
		return nil
	}

	args = append(args, userID.(string))
	fmt.Print(userID)

	// 渔民不需要输入溯源码，其他用户需要，通过雪花算法生成ID
	if userType == "渔民" {
		args = append(args, fisherman_traceability_code)
	} else {
		// 检查溯源码是否正确
		traceability_code := c.PostForm("traceability_code")
		res, err := pkg.ChaincodeQuery("GetSeafoodInfo", traceability_code)
		if res == "" || err != nil || len(traceability_code) != 18 {
			c.JSON(200, gin.H{
				"message": "请检查溯源码是否正确!!",
			})
			return nil
		} else {
			args = append(args, traceability_code)
		}
	}

	// 依次添加arg1到arg5
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	args = append(args, c.PostForm("arg5"))

	// 添加Image字段作为arg6
	image := c.PostForm("image")
	if image == "" {
		c.JSON(200, gin.H{
			"message": "Image字段不能为空",
		})
		return nil
	}
	args = append(args, image)

	return args
}
